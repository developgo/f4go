package fortran

import (
	"bytes"
	"fmt"
	goast "go/ast"
	goparser "go/parser"
	"go/token"
	"os"
	"strconv"
	"strings"
)

type varInitialization struct {
	name string
	typ  goType
}

type varInits []varInitialization

func (v varInits) get(n string) (varInitialization, bool) {
	n = strings.ToUpper(n)
	for _, val := range []varInitialization(v) {
		if val.name == n {
			return val, true
		}
	}
	return varInitialization{}, false
}

func (v *varInits) del(n string) {
	vs := []varInitialization(*v)
	n = strings.ToUpper(n)
	for i, val := range vs {
		if val.name == n {
			vs = append(vs[:i], vs[i+1:]...)
			*v = varInits(vs)
			return
		}
	}
}

func (v *varInits) add(name string, typ goType) {
	vs := []varInitialization(*v)
	vs = append(vs, varInitialization{name: strings.ToUpper(name), typ: typ})
	*v = varInits(vs)
}

func (p parser) getSize(name string, col int) (size int, ok bool) {
	v, ok := p.initVars.get(name)
	if !ok {
		panic("Cannot find variable : " + name)
	}
	if v.typ.baseType == "string" {
		col++
	}
	if len(v.typ.arrayNode[col]) == 1 && v.typ.arrayNode[col][0].tok == token.INT {
		val, _ := strconv.Atoi(string(v.typ.arrayNode[col][0].b))
		return val, true
	}
	if vv, ok := p.initVars.get(nodesToString(v.typ.arrayNode[col])); ok {
		if n, ok := p.constants[vv.name]; ok {
			size, _ = strconv.Atoi(nodesToString(n))
			return size, true
		}
	}
	for i, n := range v.typ.arrayNode[col] {
		// Example:
		// -1 : 1
		// Nodes:
		// [[-, `-`] [INT, `1`] [:, `:`] [INT, `1`]]
		// 99 : 101
		// Nodes:
		// [[INT, `99`] [:, `:`] [INT, `101`]]
		if n.tok == token.COLON {
			begin, err := strconv.Atoi(strings.Replace(nodesToString(v.typ.arrayNode[col][:i]), " ", "", -1))
			if err != nil {
				p.addError("Cannot parse begin value : " + nodesToString(v.typ.arrayNode[col][:i]))
				break
			}
			end, err := strconv.Atoi(strings.Replace(nodesToString(v.typ.arrayNode[col][i+1:]), " ", "", -1))
			if err != nil {
				p.addError("Cannot parse end value : " + nodesToString(v.typ.arrayNode[col][i+1:]))
				break
			}
			return end - begin + 1, true
		}
	}
	return -1, false
}

func (p parser) getArrayBegin(name string, col int) int {
	v, ok := p.initVars.get(name)
	if !ok {
		panic("Cannot find variable : " + name)
	}
	if v.typ.baseType == "string" {
		col++
	}
	if col >= len(v.typ.arrayNode) {
		return 1
	}
	for i, n := range v.typ.arrayNode[col] {
		// Example:
		// -1 : 1
		// Nodes:
		// [[-, `-`] [INT, `1`] [:, `:`] [INT, `1`]]
		// 99 : 101
		// Nodes:
		// [[INT, `99`] [:, `:`] [INT, `101`]]
		if n.tok == token.COLON {
			strBegin := strings.Replace(nodesToString(v.typ.arrayNode[col][:i]), " ", "", -1)
			b, err := strconv.Atoi(strBegin)
			if err != nil {
				p.addError("Cannot parse begin value: " + strBegin)
			}
			return b
		}
	}
	return 1
}

func (p parser) getArrayLen(name string) int {
	v, ok := p.initVars.get(name)
	if !ok {
		panic("Cannot find variable : " + name)
	}
	lenArray := len(v.typ.arrayNode)
	if v.typ.baseType == "string" {
		lenArray--
	}
	return lenArray
}

type parser struct {
	ast   goast.File
	ident int
	ns    []node

	functionExternalName []string

	initVars varInits // map of name to type

	comments []string

	pkgs        map[string]bool // import packages
	endLabelDo  map[string]int  // label of DO
	allLabels   map[string]bool // list of all labels
	foundLabels map[string]bool // list labels found in source

	parameters map[string]string // constants

	formats map[string][]node // source line with command FORMAT

	constants map[string][]node

	errs []error
}

func (p *parser) addImport(pkg string) {
	p.pkgs[pkg] = true
}

func (p *parser) init() {
	p.functionExternalName = make([]string, 0)
	p.endLabelDo = map[string]int{}
	p.allLabels = map[string]bool{}
	p.foundLabels = map[string]bool{}
	p.initVars = varInits{}
	p.parameters = map[string]string{}
	p.formats = map[string][]node{}

	p.constants = map[string][]node{}
}

// list view - only for debugging
func lv(ns []node) (output string) {
	for _, n := range ns {
		b := string(n.b)
		if n.tok != ftNewLine {
			output += fmt.Sprintf("%10s\t%10s\t|`%s`\n",
				view(n.tok),
				fmt.Sprintf("%v", n.pos),
				b)
		} else {
			output += fmt.Sprintf("%20s\n",
				view(n.tok))
		}
	}
	return
}

// Parse is convert fortran source to go ast tree
func Parse(b []byte, packageName string) (goast.File, []error) {

	if packageName == "" {
		packageName = "main"
	}

	var p parser

	if p.pkgs == nil {
		p.pkgs = map[string]bool{}
	}

	p.ns = scan(b)

	p.ast.Name = goast.NewIdent(packageName)

	var decls []goast.Decl
	p.ident = 0
	decls = p.parseNodes()
	if len(p.errs) > 0 {
		return p.ast, p.errs
	}

	// add packages
	for pkg := range p.pkgs {
		p.ast.Decls = append(p.ast.Decls, &goast.GenDecl{
			Tok: token.IMPORT,
			Specs: []goast.Spec{
				&goast.ImportSpec{
					Path: &goast.BasicLit{
						Kind:  token.STRING,
						Value: "\"" + pkg + "\"",
					},
				},
			},
		})
	}

	// TODO : add INTRINSIC fortran functions

	p.ast.Decls = append(p.ast.Decls, decls...)

	strC := strChanger{}
	goast.Walk(strC, &p.ast)

	return p.ast, p.errs
}

// go/ast Visitor for comment label
type commentLabel struct {
	labels map[string]bool
}

func (c commentLabel) Visit(node goast.Node) (w goast.Visitor) {
	if ident, ok := node.(*goast.Ident); ok && ident != nil {
		if _, ok := c.labels[ident.Name]; ok {
			ident.Name = "//" + ident.Name
		}
	}
	return c
}

// go/ast Visitor for change "strings" to "[]byte"
type strChanger struct {
}

func (s strChanger) Visit(node goast.Node) (w goast.Visitor) {
	if call, ok := node.(*goast.CallExpr); ok {
		if sel, ok := call.Fun.(*goast.SelectorExpr); ok {
			if id, ok := sel.X.(*goast.Ident); ok {
				if id.Name == "fmt" || id.Name == "math" {
					return nil
				}
			}
		}
	}
	if _, ok := node.(*goast.ImportSpec); ok {
		return nil
	}
	if st, ok := node.(*goast.BasicLit); ok && st.Kind == token.STRING {
		if len(st.Value) == 3 {
			st.Kind = token.CHAR
			st.Value = fmt.Sprintf("'%c'", st.Value[1])
		} else {
			st.Value = fmt.Sprintf("*func()*[]byte{y:=[]byte(%s);return &y}()",
				st.Value)
		}
	}
	return s
}

// parseNodes
func (p *parser) parseNodes() (decls []goast.Decl) {

	if p.ident < 0 || p.ident >= len(p.ns) {
		p.errs = append(p.errs,
			fmt.Errorf("Ident is outside nodes: %d/%d", p.ident, len(p.ns)))
		return
	}

	// find all names of FUNCTION, SUBROUTINE, PROGRAM
	var internalFunction []string
	for ; p.ident < len(p.ns); p.ident++ {
		switch p.ns[p.ident].tok {
		case ftSubroutine:
			p.expect(ftSubroutine)
			p.ident++
			p.expect(token.IDENT)
			internalFunction = append(internalFunction, string(p.ns[p.ident].b))
			continue
		case ftProgram:
			p.expect(ftProgram)
			p.ident++
			p.expect(token.IDENT)
			internalFunction = append(internalFunction, string(p.ns[p.ident].b))
			continue
		}

		// Example:
		//   RECURSIVE SUBROUTINE CGELQT3( M, N, A, LDA, T, LDT, INFO )
		if strings.ToUpper(string(p.ns[p.ident].b)) == "RECURSIVE" {
			p.ns[p.ident].tok, p.ns[p.ident].b = ftNewLine, []byte("\n")
			continue
		}

		// FUNCTION
		for i := p.ident; i < len(p.ns) && p.ns[i].tok != ftNewLine; i++ {
			if p.ns[p.ident].tok == ftFunction {
				p.expect(ftFunction)
				p.ident++
				p.expect(token.IDENT)
				internalFunction = append(internalFunction, string(p.ns[p.ident].b))
			}
		}
	}
	p.ident = 0

	for ; p.ident < len(p.ns); p.ident++ {
		p.init()
		p.functionExternalName = append(p.functionExternalName,
			internalFunction...)

		var next bool
		switch p.ns[p.ident].tok {
		case ftDefine:
			p.addError("Cannot parse #DEFINE: " + p.getLine())
			p.gotoEndLine()
			continue

		case ftNewLine:
			next = true // TODO
		case token.COMMENT:
			p.comments = append(p.comments,
				"//"+string(p.ns[p.ident].b))
			next = true // TODO
		case ftSubroutine: // SUBROUTINE
			var decl goast.Decl
			decl = p.parseSubroutine()
			decls = append(decls, decl)
			next = true
		case ftProgram: // PROGRAM
			var decl goast.Decl
			decl = p.parseProgram()
			decls = append(decls, decl)
			next = true
		default:
			// Example :
			//  COMPLEX FUNCTION CDOTU ( N , CX , INCX , CY , INCY )
			for i := p.ident; i < len(p.ns) && p.ns[i].tok != ftNewLine; i++ {
				if p.ns[i].tok == ftFunction {
					decl := p.parseFunction()
					decls = append(decls, decl)
					next = true
				}
			}
		}
		if next {
			continue
		}

		if p.ident >= len(p.ns) {
			break
		}

		switch p.ns[p.ident].tok {
		case ftNewLine, token.EOF:
			continue
		}

		// if at the begin we haven't SUBROUTINE , FUNCTION,...
		// then add fake Program
		var comb []node
		comb = append(comb, p.ns[:p.ident]...)
		comb = append(comb, []node{
			{tok: ftNewLine, b: []byte("\n")},
			{tok: ftProgram, b: []byte("PROGRAM")},
			{tok: token.IDENT, b: []byte("MAIN")},
			{tok: ftNewLine, b: []byte("\n")},
		}...)
		comb = append(comb, p.ns[p.ident:]...)
		p.ns = comb
		p.ident--

		fmt.Fprintf(os.Stdout, "Add fake PROGRAM MAIN in pos : %v", p.ns[p.ident].pos)
	}

	return
}

func (p *parser) gotoEndLine() {
	for ; p.ident < len(p.ns) && p.ns[p.ident].tok != ftNewLine; p.ident++ {
	}
}

func (p *parser) getLine() (line string) {
	if p.ident < 0 {
		p.ident = 0
	}
	if !(p.ident < len(p.ns)) {
		p.ident = len(p.ns) - 1
	}

	last := p.ident
	defer func() {
		p.ident = last
	}()
	for ; p.ident >= 0 && p.ns[p.ident].tok != ftNewLine; p.ident-- {
	}
	p.ident++
	for ; p.ident < len(p.ns) && p.ns[p.ident].tok != ftNewLine; p.ident++ {
		line += " " + string(p.ns[p.ident].b)
	}
	return
}

// go/ast Visitor for parse FUNCTION
type vis struct {
	// map [from] to
	c map[string]string
}

func initVis() *vis {
	var v vis
	v.c = map[string]string{}
	return &v
}

func (v vis) Visit(node goast.Node) (w goast.Visitor) {
	if ident, ok := node.(*goast.Ident); ok {
		if to, ok := v.c[strings.ToUpper(ident.Name)]; ok {
			ident.Name = "(" + to + ")"
		}
	}
	return v
}

// delete external function type definition
func (p *parser) removeExternalFunction() {
	for _, f := range p.functionExternalName {
		p.initVars.del(f)
	}
}

// add correct type of subroutine arguments
func (p *parser) argumentCorrection(fd goast.FuncDecl) (removedVars []string) {
checkArguments:
	for i := range fd.Type.Params.List {
		fieldName := fd.Type.Params.List[i].Names[0].Name
		if v, ok := p.initVars.get(fieldName); ok {
			fd.Type.Params.List[i].Type = goast.NewIdent(v.typ.String())

			// Remove to arg
			removedVars = append(removedVars, fieldName)
			p.initVars.del(fieldName)
			goto checkArguments
		}
	}
	return
}

// init vars
func (p *parser) initializeVars() (vars []goast.Stmt) {
	for i := range []varInitialization(p.initVars) {
		name := ([]varInitialization(p.initVars)[i]).name
		goT := ([]varInitialization(p.initVars)[i]).typ
		switch p.getArrayLen(name) {
		case 0:
			decl := goast.GenDecl{
				Tok: token.VAR,
				Specs: []goast.Spec{
					&goast.ValueSpec{
						Names: []*goast.Ident{
							goast.NewIdent(name),
						},
						Type: goast.NewIdent(
							goT.String()),
					},
				},
			}
			if val, ok := p.constants[name]; ok {
				decl.Specs[0].(*goast.ValueSpec).Values = []goast.Expr{
					p.parseExprNodes(val),
				}
			} else if v, ok := p.initVars.get(name); ok && v.typ.baseType == "string" {
				decl.Specs[0].(*goast.ValueSpec).Values = []goast.Expr{
					&goast.CallExpr{
						Fun: goast.NewIdent("make"),
						Args: []goast.Expr{
							&goast.ArrayType{Elt: goast.NewIdent("byte")},
							goast.NewIdent(nodesToString(v.typ.arrayNode[0])),
						},
					},
				}

			}
			vars = append(vars, &goast.DeclStmt{Decl: &decl})

		case 1: // vector
			arrayType := goT.getBaseType()
			for i := 0; i < p.getArrayLen(name); i++ {
				arrayType = "[]" + arrayType
			}
			size, ok := p.getSize(name, 0)
			if !ok {
				vars = append(vars, &goast.DeclStmt{
					Decl: &goast.GenDecl{
						Tok: token.VAR,
						Specs: []goast.Spec{
							&goast.ValueSpec{
								Names: []*goast.Ident{
									goast.NewIdent(name),
								},
								Type: goast.NewIdent(
									goT.String()),
							},
						},
					},
				})
				continue
			}
			vars = append(vars, &goast.AssignStmt{
				Lhs: []goast.Expr{goast.NewIdent(name)},
				Tok: token.DEFINE,
				Rhs: []goast.Expr{
					&goast.CallExpr{
						Fun:    goast.NewIdent("make"),
						Lparen: 1,
						Args: []goast.Expr{
							goast.NewIdent(arrayType),
							goast.NewIdent(strconv.Itoa(size)),
						},
					}},
			})

		case 2: // matrix
			fset := token.NewFileSet() // positions are relative to fset
			src := `package main
func main() {
	%s := make([][]%s, %d)
	for u := 0; u < %d; u++ {
		%s[u] = make([]%s, %d)
	}
}
`
			size0, _ := p.getSize(name, 0)
			size1, _ := p.getSize(name, 1)
			s := fmt.Sprintf(src,
				name,
				goT.getBaseType(),
				size0,
				size0,
				name,
				goT.getBaseType(),
				size1,
			)
			f, err := goparser.ParseFile(fset, "", s, 0)
			if err != nil {
				panic(fmt.Errorf("Error: %v\nSource:\n%s\npos=%s",
					err, s, goT.arrayNode))
			}
			vars = append(vars, f.Decls[0].(*goast.FuncDecl).Body.List...)

		case 3: // ()()()
			fset := token.NewFileSet() // positions are relative to fset
			src := `package main
func main() {
	%s := make([][][]%s, %d)
	for u := 0; u < %d; u++ {
		%s[u] = make([][]%s, %d)
		for w := 0; w < %d; w++ {
			%s[u][w] = make([]%s, %d)
		}
	}
}
`
			size0, _ := p.getSize(name, 0)
			size1, _ := p.getSize(name, 1)
			size2, _ := p.getSize(name, 2)
			s := fmt.Sprintf(src,
				// line 1
				name,
				goT.getBaseType(),
				size0,
				// line 2
				size0,
				// line 3
				name,
				goT.getBaseType(),
				size1,
				// line 4
				size1,
				// line 5
				name,
				goT.getBaseType(),
				size2,
			)
			f, err := goparser.ParseFile(fset, "", s, 0)
			if err != nil {
				panic(fmt.Errorf("Error: %v\nSource:\n%s\npos=%s",
					err, s, goT.arrayNode))
			}
			vars = append(vars, f.Decls[0].(*goast.FuncDecl).Body.List...)
		default:
			panic(fmt.Errorf(
				"not correct amount of array : %v", goT))
		}
	}

	return
}

// go/ast Visitor for comment label
type callArg struct {
	p *parser
}

// Example
//  From :
// ab_min(3, 14)
//  To:
// ab_min(func() *int { y := 3; return &y }(), func() *int { y := 14; return &y }())
func (c callArg) Visit(node goast.Node) (w goast.Visitor) {
	if call, ok := node.(*goast.CallExpr); ok && call != nil {

		if sel, ok := call.Fun.(*goast.SelectorExpr); ok {
			if name, ok := sel.X.(*goast.Ident); ok {
				if name.Name == "math" || name.Name == "fmt" {
					goto end
				}
			}
		}
		if call, ok := node.(*goast.CallExpr); ok {
			if id, ok := call.Fun.(*goast.Ident); ok {
				if id.Name == "append" {
					return nil
				}
			}
		}

		for i := range call.Args {
			switch a := call.Args[i].(type) {
			case *goast.BasicLit:
				switch a.Kind {
				case token.STRING:
					call.Args[i] = goast.NewIdent(
						fmt.Sprintf("func()*[]byte{y:=[]byte(%s);return &y}()", a.Value))
					if len(a.Value) == 3 {
						a.Value = strings.Replace(a.Value, "\"", "'", -1)
						call.Args[i] = goast.NewIdent(
							fmt.Sprintf("func()*byte{y:=byte(%s);return &y}()", a.Value))
					}
				case token.INT:
					call.Args[i] = goast.NewIdent(
						fmt.Sprintf("func()*int{y:=%s;return &y}()", a.Value))
				case token.FLOAT:
					call.Args[i] = goast.NewIdent(
						fmt.Sprintf("func()*float64{y:=%s;return &y}()", a.Value))
				case token.CHAR:
					call.Args[i] = goast.NewIdent(
						fmt.Sprintf("func()*byte{y:=%s;return &y}()", a.Value))
				default:
					panic(fmt.Errorf(
						"Not support basiclit token: %T ", a.Kind))
				}

			case *goast.Ident: // TODO : not correct for array
				id := call.Args[i].(*goast.Ident)
				id.Name = "&(" + id.Name + ")"

			case *goast.IndexExpr:
				call.Args[i] = &goast.UnaryExpr{
					Op: token.AND,
					X: &goast.ParenExpr{
						Lparen: 1,
						X:      call.Args[i],
					},
				}

				// TODO:
				// default:
				// 	goast.Print(token.NewFileSet(), a)
				// 	panic(fmt.Errorf(
				// 		"Not support arg call token: %T ", a))
			}
		}
	}
end:
	return c
}

// Example :
//  COMPLEX FUNCTION CDOTU ( N , CX , INCX , CY , INCY )
//  DOUBLE PRECISION FUNCTION DNRM2 ( N , X , INCX )
//  COMPLEX * 16 FUNCTION ZDOTC ( N , ZX , INCX , ZY , INCY )
func (p *parser) parseFunction() (decl goast.Decl) {
	for i := p.ident; i < len(p.ns) && p.ns[i].tok != ftNewLine; i++ {
		if p.ns[i].tok == ftFunction {
			p.ns[i].tok = ftSubroutine
		}
	}
	return p.parseSubroutine()
}

// Example:
//   PROGRAM MAIN
func (p *parser) parseProgram() (decl goast.Decl) {
	p.expect(ftProgram)
	p.ns[p.ident].tok = ftSubroutine
	decl = p.parseSubroutine()
	if fd, ok := decl.(*goast.FuncDecl); ok {
		if strings.ToUpper(fd.Name.Name) == "MAIN" {
			fd.Name.Name = "main"
		}
	}
	return
}

// parseSubroutine  is parsed SUBROUTINE, FUNCTION, PROGRAM
// Example :
//  SUBROUTINE CHBMV ( UPLO , N , K , ALPHA , A , LDA , X , INCX , BETA , Y , INCY )
//  PROGRAM MAIN
//  COMPLEX FUNCTION CDOTU ( N , CX , INCX , CY , INCY )
func (p *parser) parseSubroutine() (decl goast.Decl) {
	var fd goast.FuncDecl
	fd.Type = &goast.FuncType{
		Params: &goast.FieldList{},
	}

	defer func() {
		fd.Doc = &goast.CommentGroup{}
		for _, c := range p.comments {
			fd.Doc.List = append(fd.Doc.List, &goast.Comment{
				Text: c,
			})
		}
		p.comments = []string{}
	}()

	// check return type
	var returnType []node
	for ; p.ns[p.ident].tok != ftSubroutine && p.ns[p.ident].tok != ftNewLine; p.ident++ {
		returnType = append(returnType, p.ns[p.ident])
	}

	p.expect(ftSubroutine)

	p.ident++
	p.expect(token.IDENT)
	name := strings.ToUpper(string(p.ns[p.ident].b))
	fd.Name = goast.NewIdent(name)

	// Add return type is exist
	returnName := name + "_RES"
	if len(returnType) > 0 {
		fd.Type.Results = &goast.FieldList{
			List: []*goast.Field{
				{
					Names: []*goast.Ident{goast.NewIdent(returnName)},
					Type:  goast.NewIdent(parseType(returnType).String()),
				},
			},
		}
	}
	defer func() {
		// change function name variable to returnName
		if len(returnType) > 0 {
			v := initVis()
			v.c[name] = returnName
			goast.Walk(v, fd.Body)
		}
	}()

	// Parameters
	p.ident++
	fd.Type.Params.List = p.parseParamDecl()

	p.ident++
	fd.Body = &goast.BlockStmt{
		Lbrace: 1,
		List:   p.parseListStmt(),
	}

	// delete external function type definition
	p.removeExternalFunction()

	// remove from arguments arg with type string
	arrayArguments := map[string]bool{}
	for i := range fd.Type.Params.List {
		fieldName := fd.Type.Params.List[i].Names[0].Name
		if v, ok := p.initVars.get(fieldName); ok {
			if v.typ.isArray() {
				arrayArguments[fieldName] = true
			}
		}
	}

	// add correct type of subroutine arguments
	arguments := p.argumentCorrection(fd)

	// change arguments
	// From:
	//  a
	// To:
	//  *a
	v := initVis()
	for _, arg := range arguments {
		v.c[arg] = "*" + arg
	}
	goast.Walk(v, fd.Body)

	// changes arguments in func
	for i := range fd.Type.Params.List {
		switch fd.Type.Params.List[i].Type.(type) {
		case *goast.Ident:
			id := fd.Type.Params.List[i].Type.(*goast.Ident)
			id.Name = "*" + id.Name
		default:
			panic(fmt.Errorf("Cannot parse type in fields: %T",
				fd.Type.Params.List[i].Type))
		}
	}

	// replace call argument constants
	c := callArg{p: p}
	goast.Walk(c, fd.Body)

	// init vars
	fd.Body.List = append(p.initializeVars(), fd.Body.List...)

	// remove unused labels
	removedLabels := map[string]bool{}
	for k := range p.allLabels {
		if _, ok := p.foundLabels[k]; !ok {
			removedLabels[k] = true
		}
	}
	cl := commentLabel{labels: removedLabels}
	goast.Walk(cl, fd.Body)

	in := intrinsic{p: p}
	goast.Walk(in, fd.Body)

	var cas callArgumentSimplification
	goast.Walk(cas, fd.Body)

	decl = &fd
	return
}

func (p *parser) addError(msg string) {
	last := p.ident
	defer func() {
		p.ident = last
	}()

	p.errs = append(p.errs, fmt.Errorf("%s", msg))
}

func (p *parser) expect(t token.Token) {
	if t != p.ns[p.ident].tok {
		// Show all errors
		for _, err := range p.errs {
			fmt.Println("Error : ", err.Error())
		}
		// Panic
		panic(fmt.Errorf("Expect %s, but we have {{%s,%s}}. Pos = %v",
			view(t), view(p.ns[p.ident].tok), string(p.ns[p.ident].b),
			p.ns[p.ident].pos))
	}
}

func (p *parser) parseListStmt() (stmts []goast.Stmt) {
	for p.ident < len(p.ns) {
		// Only for debugging
		// fmt.Println("---------------")
		// for i := 0; i < len(p.ns); i++ {
		// 	if p.ns[i].tok != ftNewLine {
		// 		fmt.Printf("%8v %4d %3d %v %v\n", p.ident == i,
		// 			i, p.ns[i].pos.line, p.ns[i].tok, string(p.ns[i].b))
		// 		continue
		// 	}
		// 	fmt.Printf("%8v %4d %3d %v\n", p.ident == i,
		// 		i, p.ns[i].pos.line, p.ns[i].tok)
		// }

		if p.ns[p.ident].tok == token.COMMENT {
			stmts = append(stmts, &goast.ExprStmt{
				X: goast.NewIdent("//" + string(p.ns[p.ident].b)),
			})
			p.ident++
			continue
		}
		if p.ns[p.ident].tok == ftNewLine {
			p.ident++
			continue
		}

		if p.ns[p.ident].tok == ftEnd {
			p.ident++
			p.gotoEndLine()
			// TODO need gotoEndLine() ??
			break
		}
		if p.ns[p.ident].tok == token.ELSE {
			// gotoEndLine() is no need for case:
			// ELSE IF (...)...
			break
		}

		stmt := p.parseStmt()
		if stmt == nil {
			// p.addError("stmt is nil in line ")
			// break
			continue
		}
		stmts = append(stmts, stmt...)
	}
	return
}

// Examples:
//  INTEGER INCX , INCY , N
//  COMPLEX CX ( * ) , CY ( * )
//  COMPLEX*16 A(LDA,*),X(*)
//  REAL A(LDA,*),B(LDB,*)
//  DOUBLE PRECISION DX(*)
//  LOGICAL CONJA,CONJB,NOTA,NOTB
//  CHARACTER*32 SRNAME
func (p *parser) parseInit() (stmts []goast.Stmt) {

	// parse base type
	var baseType []node
	for ; p.ns[p.ident].tok != token.IDENT; p.ident++ {
		baseType = append(baseType, p.ns[p.ident])
	}
	p.expect(token.IDENT)

	var name string
	var additionType []node
	for ; p.ns[p.ident].tok != ftNewLine &&
		p.ns[p.ident].tok != token.EOF; p.ident++ {
		// parse name
		p.expect(token.IDENT)
		name = string(p.ns[p.ident].b)

		// parse addition type
		additionType = []node{}
		p.ident++
		for ; p.ns[p.ident].tok != ftNewLine &&
			p.ns[p.ident].tok != token.EOF &&
			p.ns[p.ident].tok != token.COMMA; p.ident++ {
			if p.ns[p.ident].tok == token.LPAREN {
				counter := 0
				for ; ; p.ident++ {
					switch p.ns[p.ident].tok {
					case token.LPAREN:
						counter++
					case token.RPAREN:
						counter--
					case ftNewLine:
						p.addError("Cannot parse type : not expected NEW_LINE")
						return
					}
					if counter == 0 {
						break
					}
					additionType = append(additionType, p.ns[p.ident])
				}
			}
			additionType = append(additionType, p.ns[p.ident])
		}

		// parse type = base type + addition type
		p.initVars.add(name, parseType(append(baseType, additionType...)))
		if p.ns[p.ident].tok != token.COMMA {
			p.ident--
		}
	}

	return
}

func (p *parser) parseDoWhile() (sDo goast.ForStmt) {
	p.expect(ftDo)
	p.ident++
	p.expect(ftWhile)
	p.ident++
	start := p.ident
	for ; p.ident < len(p.ns); p.ident++ {
		if p.ns[p.ident].tok == ftNewLine {
			break
		}
	}
	sDo.Cond = p.parseExpr(start, p.ident)

	p.expect(ftNewLine)
	p.ident++

	sDo.Body = &goast.BlockStmt{
		Lbrace: 1,
		List:   p.parseListStmt(),
	}

	return
}

func (p *parser) parseDo() (sDo goast.ForStmt) {
	p.expect(ftDo)
	p.ident++
	if p.ns[p.ident].tok == ftWhile {
		p.ident--
		return p.parseDoWhile()
	}
	// possible label
	if p.ns[p.ident].tok == token.INT {
		p.endLabelDo[string(p.ns[p.ident].b)]++
		p.ident++
	}
	// for case with comma "DO 40, J = 1, N"
	if p.ns[p.ident].tok == token.COMMA {
		p.ident++
	}

	p.expect(token.IDENT)
	name := string(p.ns[p.ident].b)

	p.ident++
	p.expect(token.ASSIGN)

	p.ident++
	// Init is expression
	start := p.ident
	counter := 0
	for ; p.ident < len(p.ns); p.ident++ {
		if p.ns[p.ident].tok == token.LPAREN {
			counter++
			continue
		}
		if p.ns[p.ident].tok == token.RPAREN {
			counter--
			continue
		}
		if p.ns[p.ident].tok == token.COMMA && counter == 0 {
			break
		}
	}
	sDo.Init = &goast.AssignStmt{
		Lhs: []goast.Expr{
			goast.NewIdent(name),
		},
		Tok: token.ASSIGN,
		Rhs: []goast.Expr{
			p.parseExpr(start, p.ident),
		},
	}

	p.expect(token.COMMA)

	// Cond is expression
	p.ident++
	start = p.ident
	counter = 0
	for ; p.ident < len(p.ns); p.ident++ {
		if p.ns[p.ident].tok == token.LPAREN {
			counter++
			continue
		}
		if p.ns[p.ident].tok == token.RPAREN {
			counter--
			continue
		}
		if (p.ns[p.ident].tok == token.COMMA || p.ns[p.ident].tok == ftNewLine) &&
			counter == 0 {
			break
		}
	}
	sDo.Cond = &goast.BinaryExpr{
		X:  goast.NewIdent(name),
		Op: token.LEQ,
		Y:  p.parseExpr(start, p.ident),
	}

	if p.ns[p.ident].tok == ftNewLine {
		sDo.Post = &goast.IncDecStmt{
			X:   goast.NewIdent(name),
			Tok: token.INC,
		}
	} else {
		p.expect(token.COMMA)
		p.ident++

		// Post is expression
		start = p.ident
		for ; p.ident < len(p.ns); p.ident++ {
			if p.ns[p.ident].tok == ftNewLine {
				break
			}
		}
		sDo.Post = &goast.AssignStmt{
			Lhs: []goast.Expr{goast.NewIdent(name)},
			Tok: token.ADD_ASSIGN,
			Rhs: []goast.Expr{p.parseExpr(start, p.ident)},
		}
	}

	p.expect(ftNewLine)

	sDo.Body = &goast.BlockStmt{
		Lbrace: 1,
		List:   p.parseListStmt(),
	}

	return
}

func (p *parser) parseIf() (sIf goast.IfStmt) {
	p.ident++
	p.expect(token.LPAREN)

	p.ident++
	start := p.ident
	for counter := 1; p.ns[p.ident].tok != token.EOF; p.ident++ {
		var exit bool
		switch p.ns[p.ident].tok {
		case token.LPAREN:
			counter++
		case token.RPAREN:
			counter--
			if counter == 0 {
				exit = true
			}
		}
		if exit {
			break
		}
	}

	sIf.Cond = p.parseExpr(start, p.ident)

	p.expect(token.RPAREN)
	p.ident++

	if p.ns[p.ident].tok == ftThen {
		p.gotoEndLine()
		p.ident++
		sIf.Body = &goast.BlockStmt{
			Lbrace: 1,
			List:   p.parseListStmt(),
		}
	} else {
		sIf.Body = &goast.BlockStmt{
			Lbrace: 1,
			List:   p.parseStmt(),
		}
		return
	}

	if p.ident >= len(p.ns) {
		return
	}

	if p.ns[p.ident].tok == token.ELSE {
		p.ident++
		if p.ns[p.ident].tok == token.IF {
			ifr := p.parseIf()
			sIf.Else = &ifr
		} else {
			sIf.Else = &goast.BlockStmt{
				Lbrace: 1,
				List:   p.parseListStmt(),
			}
		}
	}

	return
}

func (p *parser) parseExternal() {
	p.expect(ftExternal)

	p.ident++
	for ; p.ns[p.ident].tok != token.EOF; p.ident++ {
		if p.ns[p.ident].tok == ftNewLine {
			p.ident++
			break
		}
		switch p.ns[p.ident].tok {
		case token.IDENT, ftInteger, ftReal, ftComplex:
			name := string(p.ns[p.ident].b)
			p.functionExternalName = append(p.functionExternalName, name)
			// fmt.Println("Function external: ", name)
		case token.COMMA:
			// ingore
		default:
			p.addError("Cannot parse External " + string(p.ns[p.ident].b))
		}
	}
}

func (p *parser) parseStmt() (stmts []goast.Stmt) {

	pos := p.ns[p.ident].pos

	defer func() {
		if r := recover(); r != nil {
			p.addError(fmt.Sprintf("Recover parseStmt pos{%v}: %v", pos, r))
			p.gotoEndLine()
		}
	}()

	switch p.ns[p.ident].tok {
	case ftInteger, ftCharacter, ftComplex, ftLogical, ftReal, ftDouble:
		stmts = append(stmts, p.parseInit()...)

	case ftEquivalence:
		p.addError(p.getLine())
		p.gotoEndLine()

	case ftRewind:
		s := p.parseRewind()
		stmts = append(stmts, s...)

	case ftFormat:
		stmts = append(stmts, &goast.ExprStmt{
			X: goast.NewIdent("// Unused by f4go : " + p.getLine()),
		})
		p.gotoEndLine()

	case ftCommon:
		// TODO: Add support COMMON, use memory pool
		p.gotoEndLine()

	case token.RETURN:
		stmts = append(stmts, &goast.ReturnStmt{})
		p.gotoEndLine()
		p.expect(ftNewLine)

	case ftParameter:
		//  PARAMETER ( ONE = ( 1.0E+0 , 0.0E+0 )  , ZERO = 0.0E+0 )
		s := p.parseParameter()
		stmts = append(stmts, s...)

	case ftOpen:
		s := p.parseOpen()
		stmts = append(stmts, s...)

	case ftRead:
		s := p.parseRead()
		stmts = append(stmts, s...)

	case ftClose:
		s := p.parseClose()
		stmts = append(stmts, s...)

	case ftAssign:
		s := p.parseAssign()
		stmts = append(stmts, s...)

	case ftDefine:
		p.addError("#DEFINE is not support :" + p.getLine())
		p.gotoEndLine()

	case ftSave:
		p.expect(ftSave)
		// ignore command SAVE
		// that command only for optimization
		p.gotoEndLine()

	case ftExternal:
		p.parseExternal()

	case ftNewLine:
		// ignore
		p.ident++

	case token.IF:
		sIf := p.parseIf()
		stmts = append(stmts, &sIf)

	case ftDo:
		sDo := p.parseDo()
		stmts = append(stmts, &sDo)

	case ftCall:
		// Example:
		// CALL XERBLA ( 'CGEMM ' , INFO )
		p.expect(ftCall)
		p.ident++
		start := p.ident
		for ; p.ns[p.ident].tok != ftNewLine; p.ident++ {
		}
		f := p.parseExpr(start, p.ident)
		stmts = append(stmts, &goast.ExprStmt{
			X: f,
		})
		p.expect(ftNewLine)

	case ftIntrinsic:
		// Example:
		//  INTRINSIC CONJG , MAX
		p.expect(ftIntrinsic)
		p.ns[p.ident].tok = ftExternal
		p.parseExternal()

	case ftData:
		// Example:
		// DATA GAM , GAMSQ , RGAMSQ / 4096.D0 , 16777216.D0 , 5.9604645D-8 /
		sData := p.parseData()
		stmts = append(stmts, sData...)

	case ftWrite:
		sWrite := p.parseWrite()
		stmts = append(stmts, sWrite...)

	case ftStop:
		p.expect(ftStop)
		p.ident++
		p.expect(ftNewLine)
		stmts = append(stmts, &goast.ReturnStmt{})

	case token.GOTO:
		// Examples:
		//  GO TO 30
		//  GO TO ( 40, 80 )IEXC
		// TODO: go to next,(30, 50, 70, 90, 110)
		sGoto := p.parseGoto()
		stmts = append(stmts, sGoto...)
		p.expect(ftNewLine)

	case ftImplicit:
		// TODO: add support IMPLICIT
		var nodes []node
		for ; p.ident < len(p.ns); p.ident++ {
			if p.ns[p.ident].tok == ftNewLine || p.ns[p.ident].tok == token.EOF {
				break
			}
			nodes = append(nodes, p.ns[p.ident])
		}
		// p.addError("IMPLICIT is not support.\n" + nodesToString(nodes))
		// ignore
		_ = nodes

	case token.INT:
		labelName := string(p.ns[p.ident].b)
		if v, ok := p.endLabelDo[labelName]; ok && v > 0 {
			// if after END DO, then remove
			for i := p.ident; p.ns[i].tok != ftNewLine; i++ {
				p.ns[i].tok, p.ns[i].b = ftNewLine, []byte("\n")
			}

			// add END DO before that label
			var add []node
			for j := 0; j < v; j++ {
				add = append(add, []node{
					{tok: ftNewLine, b: []byte("\n")},
					{tok: ftEnd, b: []byte("END")},
					{tok: ftNewLine, b: []byte("\n")},
				}...)
			}
			var comb []node
			comb = append(comb, p.ns[:p.ident-1]...)
			comb = append(comb, []node{
				{tok: ftNewLine, b: []byte("\n")},
				{tok: ftNewLine, b: []byte("\n")},
			}...)
			comb = append(comb, add...)
			comb = append(comb, []node{
				{tok: ftNewLine, b: []byte("\n")},
			}...)
			comb = append(comb, p.ns[p.ident-1:]...)
			p.ns = comb
			// remove do labels from map
			p.endLabelDo[labelName] = 0
			return
		}

		if p.ns[p.ident+1].tok == token.CONTINUE {
			stmts = append(stmts, p.addLabel(p.ns[p.ident].b))
			// replace CONTINUE to NEW_LINE
			p.ident++
			p.ns[p.ident].tok, p.ns[p.ident].b = ftNewLine, []byte("\n")
			return
		}

		stmts = append(stmts, p.addLabel(p.ns[p.ident].b))
		p.ident++
		return

	default:
		start := p.ident
		for ; p.ident < len(p.ns); p.ident++ {
			if p.ns[p.ident].tok == ftNewLine {
				break
			}
		}
		var isAssignStmt bool
		pos := start
		if p.ns[start].tok == token.IDENT {
			pos++
			if p.ns[pos].tok == token.LPAREN {
				counter := 0
				for ; pos < len(p.ns); pos++ {
					switch p.ns[pos].tok {
					case token.LPAREN:
						counter++
					case token.RPAREN:
						counter--
					}
					if counter == 0 {
						break
					}
				}
				pos++
			}
			if p.ns[pos].tok == token.ASSIGN {
				isAssignStmt = true
			}
		}

		if isAssignStmt {
			assign := goast.AssignStmt{
				Lhs: []goast.Expr{p.parseExpr(start, pos)},
				Tok: token.ASSIGN,
				Rhs: []goast.Expr{p.parseExpr(pos+1, p.ident)},
			}
			stmts = append(stmts, &assign)
		} else {
			stmts = append(stmts, &goast.ExprStmt{
				X: p.parseExpr(start, p.ident),
			})
		}

		p.ident++
	}

	return
}

func (p *parser) addLabel(label []byte) (stmt goast.Stmt) {
	labelName := "Label" + string(label)
	p.allLabels[labelName] = true
	return &goast.LabeledStmt{
		Label: goast.NewIdent(labelName),
		Colon: 1,
		Stmt:  &goast.EmptyStmt{},
	}
}

func (p *parser) parseParamDecl() (fields []*goast.Field) {
	if p.ns[p.ident].tok != token.LPAREN {
		// Function or SUBROUTINE without arguments
		// Example:
		//  SubRoutine CLS
		return
	}
	p.expect(token.LPAREN)

	// Parameters
	p.ident++
	for ; p.ns[p.ident].tok != token.EOF; p.ident++ {
		var exit bool
		switch p.ns[p.ident].tok {
		case token.COMMA:
			// ignore
		case token.IDENT:
			id := strings.ToUpper(string(p.ns[p.ident].b))
			field := &goast.Field{
				Names: []*goast.Ident{goast.NewIdent(id)},
				Type:  goast.NewIdent("int"),
			}
			fields = append(fields, field)
		case token.RPAREN:
			p.ident--
			exit = true
		default:
			p.addError("Cannot parse parameter decl " + string(p.ns[p.ident].b))
			return
		}
		if exit {
			break
		}
	}

	p.ident++
	p.expect(token.RPAREN)

	p.ident++
	p.expect(ftNewLine)

	return
}

// Example:
// DATA GAM , GAMSQ , RGAMSQ / 4096.D0 , 16777216.D0 , 5.9604645D-8 /
//
// LOGICAL            ZSWAP( 4 )
// DATA               ZSWAP / .FALSE., .FALSE., .TRUE., .TRUE. /
//
// INTEGER            IPIVOT( 4, 4 )
// DATA               IPIVOT / 1, 2, 3, 4, 2, 1, 4, 3, 3, 4, 1, 2, 4, 3, 2, 1 /
//
// INTEGER            LOCL12( 4 ), LOCU21( 4 ),
// DATA               LOCU12 / 3, 4, 1, 2 / , LOCL21 / 2, 1, 4, 3 /
//
// TODO:
//
// INTEGER            LV, IPW2
// PARAMETER          ( LV = 128 )
// INTEGER            J
// INTEGER            MM( LV, 4 )
// DATA               ( MM( 1, J ), J = 1, 4 ) / 494, 322, 2508, 2549 /

func (p *parser) parseData() (stmts []goast.Stmt) {
	p.expect(ftData)
	p.ident++

	// parse names and values
	var names [][]node
	names = append(names, []node{})
	var values [][]node
	counter := 0
	isNames := true
	for ; p.ident < len(p.ns); p.ident++ {
		if p.ns[p.ident].tok == ftNewLine {
			break
		}
		if p.ns[p.ident].tok == token.QUO {
			if isNames {
				values = append(values, []node{})
			}
			isNames = !isNames
			continue
		}
		if p.ns[p.ident].tok == token.LPAREN {
			counter++
		}
		if p.ns[p.ident].tok == token.RPAREN {
			counter--
		}
		if p.ns[p.ident].tok == token.COMMA && counter == 0 {
			if isNames {
				names = append(names, []node{})
			} else {
				values = append(values, []node{})
			}
			continue
		}
		if isNames {
			names[len(names)-1] = append(names[len(names)-1], p.ns[p.ident])
		} else {
			values[len(values)-1] = append(values[len(values)-1], p.ns[p.ident])
		}
	}

	// Example of names:
	// LL                       - value
	// LL                       - vector fully
	// LL                       - matrix fully
	// LL (1)                   - one value of vector
	// LL (1,1)                 - one value of matrix
	// (LL( J ), J = 1, 4 )     - one row of vector
	// (LL( 1, J ), J = 1, 4 )  - one row of matrix
	type tExpr struct {
		expr   goast.Expr
		isByte bool
	}
	var nameExpr []tExpr
	for _, name := range names {
		if len(name) == 1 {
			// LL                       - value
			// LL                       - vector fully
			// LL                       - matrix fully
			v, ok := p.initVars.get(nodesToString(name))
			if !ok {
				p.initVars.add(nodesToString(name), goType{
					baseType: "float64",
				})
				v, _ = p.initVars.get(nodesToString(name))
			}
			lenArray := p.getArrayLen(v.name)
			isByte := v.typ.getBaseType() == "byte"
			switch lenArray {
			case 0:
				nameExpr = append(nameExpr, tExpr{
					expr:   p.parseExprNodes(name),
					isByte: isByte,
				})
			case 1: // vector
				size, ok := p.getSize(v.name, 0)
				if !ok {
					panic("Not ok : " + v.name)
				}
				for i := 0; i < size; i++ {
					nameExpr = append(nameExpr, tExpr{
						expr: &goast.IndexExpr{
							X:      goast.NewIdent(nodesToString(name)),
							Lbrack: 1,
							Index: &goast.BasicLit{
								Kind:  token.INT,
								Value: strconv.Itoa(i),
							},
						},
						isByte: isByte})
				}
			case 2: // matrix
				size0, _ := p.getSize(v.name, 0)
				size1, _ := p.getSize(v.name, 1)
				for i := 0; i < size0; i++ {
					for j := 0; j < size1; j++ {
						nameExpr = append(nameExpr, tExpr{
							expr: &goast.IndexExpr{
								X: &goast.IndexExpr{
									X:      goast.NewIdent(nodesToString(name)),
									Lbrack: 1,
									Index: &goast.BasicLit{
										Kind:  token.INT,
										Value: strconv.Itoa(j),
									},
								},
								Lbrack: 1,
								Index: &goast.BasicLit{
									Kind:  token.INT,
									Value: strconv.Itoa(i),
								},
							},
							isByte: isByte})
					}
				}
			case 3: //matrix ()()()
				size0, _ := p.getSize(v.name, 0)
				size1, _ := p.getSize(v.name, 1)
				size2, _ := p.getSize(v.name, 2)
				for k := 0; k < size2; k++ {
					for j := 0; j < size1; j++ {
						for i := 0; i < size0; i++ {
							nameExpr = append(nameExpr, tExpr{
								expr: &goast.IndexExpr{
									X: &goast.IndexExpr{
										X: &goast.IndexExpr{
											X:      goast.NewIdent(nodesToString(name)),
											Lbrack: 1,
											Index: &goast.BasicLit{
												Kind:  token.INT,
												Value: strconv.Itoa(i),
											},
										},
										Lbrack: 1,
										Index: &goast.BasicLit{
											Kind:  token.INT,
											Value: strconv.Itoa(j),
										},
									},
									Lbrack: 1,
									Index: &goast.BasicLit{
										Kind:  token.INT,
										Value: strconv.Itoa(k),
									},
								},
								isByte: isByte})
						}
					}
				}
			default:
				panic("Not acceptable type : " + nodesToString(name))
			}

			continue
		}
		if v, ok := p.initVars.get(string(name[0].b)); ok {
			lenArray := p.getArrayLen(v.name)
			isByte := v.typ.getBaseType() == "byte"
			switch lenArray {
			case 1: // vector
				// LL (1)                   - one value of vector
				nameExpr = append(nameExpr, tExpr{
					expr: &goast.IndexExpr{
						X:      goast.NewIdent(string(name[0].b)),
						Lbrack: 1,
						Index: goast.NewIdent(nodesToString(
							append(name[2:len(name)-1], []node{
								{tok: token.SUB, b: []byte("-")},
								{tok: token.INT, b: []byte("1")},
							}...))),
					},
					isByte: isByte})

			case 2: // matrix
				// LL (1,1)                 - one value of matrix
				var mid int
				for mid = 2; mid < len(name); mid++ {
					if name[mid].tok == token.COMMA {
						break
					}
				}
				var i []node
				i = append(i, name[2:mid]...)
				i = append(i, []node{
					{tok: token.SUB, b: []byte("-")},
					{tok: token.INT, b: []byte("1")},
				}...)
				nameExpr = append(nameExpr, tExpr{
					expr: &goast.IndexExpr{
						X: &goast.IndexExpr{
							X:      goast.NewIdent(string(name[0].b)),
							Lbrack: 1,
							Index:  goast.NewIdent(nodesToString(i)),
						},
						Lbrack: 1,
						Index: goast.NewIdent(nodesToString(
							append(name[mid+1:len(name)-1], []node{
								{tok: token.SUB, b: []byte("-")},
								{tok: token.INT, b: []byte("1")},
							}...))),
					},
					isByte: isByte})

			default:
				panic("Not acceptable type : " + nodesToString(name))
			}
			continue
		}

		if v, ok := p.initVars.get(string(name[1].b)); ok {
			isByte := v.typ.getBaseType() == "byte"
			switch p.getArrayLen(v.name) {
			case 1: // vector
				// (LL( J ), J = 1, 4 )     - one row of vector
				start, _ := strconv.Atoi(string(name[8].b))
				end, _ := strconv.Atoi(string(name[10].b))
				for i := start - 1; i < end; i++ {
					nameExpr = append(nameExpr, tExpr{
						expr: &goast.IndexExpr{
							X:      goast.NewIdent(string(name[1].b)),
							Lbrack: 1,
							Index: &goast.BasicLit{
								Kind:  token.INT,
								Value: strconv.Itoa(i),
							},
						},
						isByte: isByte})
				}
				stmts = append(stmts, &goast.AssignStmt{
					Lhs: []goast.Expr{goast.NewIdent("_")},
					Tok: token.ASSIGN,
					Rhs: []goast.Expr{goast.NewIdent(string(name[3].b))},
				})

			case 2: // matrix
				// (LL( 1, J ), J = 1, 4 )  - one row of matrix
				if bytes.Equal(name[3].b, name[8].b) {
					// (LL( J, 1 ), J = 1, 4 )  - one row of matrix
					panic("TODO: Not support")
				}
				c, _ := strconv.Atoi(string(name[3].b))
				start, _ := strconv.Atoi(string(name[10].b))
				end, _ := strconv.Atoi(string(name[12].b))
				for j := start - 1; j < end; j++ {
					nameExpr = append(nameExpr, tExpr{
						expr: &goast.IndexExpr{
							X: &goast.IndexExpr{
								X:      goast.NewIdent(string(name[1].b)),
								Lbrack: 1,
								Index: &goast.BasicLit{
									Kind:  token.INT,
									Value: strconv.Itoa(c - 1),
								},
							},
							Lbrack: 1,
							Index: &goast.BasicLit{
								Kind:  token.INT,
								Value: strconv.Itoa(j),
							},
						},
						isByte: isByte})
				}
				stmts = append(stmts, &goast.AssignStmt{
					Lhs: []goast.Expr{goast.NewIdent("_")},
					Tok: token.ASSIGN,
					Rhs: []goast.Expr{goast.NewIdent(string(name[5].b))},
				})

			default:
				panic("Not acceptable type : " + nodesToString(name))
			}
			continue
		}
		if v, ok := p.initVars.get(string(name[2].b)); ok {
			isByte := v.typ.getBaseType() == "byte"
			switch p.getArrayLen(v.name) {
			case 3: // ()()()
				// ((CV(I,J,1),I=1,2),J=1,2)
				startI, _ := strconv.Atoi(string(name[13].b))
				endI, _ := strconv.Atoi(string(name[15].b))
				startJ, _ := strconv.Atoi(string(name[20].b))
				endJ, _ := strconv.Atoi(string(name[22].b))
				valueK, _ := strconv.Atoi(string(name[8].b))

				for j := startJ - 1; j < endJ; j++ {
					for i := startI - 1; i < endI; i++ {
						nameExpr = append(nameExpr, tExpr{
							expr: &goast.IndexExpr{
								X: &goast.IndexExpr{
									X: &goast.IndexExpr{
										X:      goast.NewIdent(string(name[2].b)),
										Lbrack: 1,
										Index: &goast.BasicLit{
											Kind:  token.INT,
											Value: strconv.Itoa(i),
										},
									},
									Lbrack: 1,
									Index: &goast.BasicLit{
										Kind:  token.INT,
										Value: strconv.Itoa(j),
									},
								},
								Lbrack: 1,
								Index: &goast.BasicLit{
									Kind:  token.INT,
									Value: strconv.Itoa(valueK - 1),
								},
							},
							isByte: isByte})
					}
				}

				stmts = append(stmts, &goast.AssignStmt{
					Lhs: []goast.Expr{goast.NewIdent("_")},
					Tok: token.ASSIGN,
					Rhs: []goast.Expr{goast.NewIdent(string(name[4].b))},
				})
				stmts = append(stmts, &goast.AssignStmt{
					Lhs: []goast.Expr{goast.NewIdent("_")},
					Tok: token.ASSIGN,
					Rhs: []goast.Expr{goast.NewIdent(string(name[6].b))},
				})
			}
			continue
		}

		panic("Not acceptable type : " + nodesToString(name))
	}

mul:
	for k := range values {
		// Example :
		// DATA R / 5*6 /
		// Equal:
		// DATA R / 6,6,6,6,6 /
		var haveStar bool
		var starPos int
		for i, vi := range values[k] {
			if vi.tok == token.MUL {
				haveStar = true
				starPos = i
				break
			}
		}

		if !haveStar {
			continue
		}

		amount, _ := strconv.Atoi(string(values[k][starPos-1].b))
		var inject [][]node
		for i := 0; i < amount; i++ {
			inject = append(inject, values[k][starPos+1:])
		}
		values = append(values[:k], append(inject, values[k+1:]...)...)
		goto mul
	}

	if len(nameExpr) != len(values) {
		var str string
		for i := range names {
			str += fmt.Sprintln(">>", nodesToString(names[i]))
			v, ok := p.initVars.get(nodesToString(names[i]))
			if ok {
				fmt.Println("1) ", v.name)
				fmt.Println("2) ", v.typ)
				fmt.Println("3) ", v.typ.baseType)
				fmt.Println("4) ", v.typ.getBaseType())
				fmt.Println("5) ", v.typ.arrayNode)
			}
		}
		for i := range values {
			str += fmt.Sprintln("<<", nodesToString(values[i]))
		}
		panic(fmt.Errorf("Size is not same %d!=%d\n%v",
			len(nameExpr), len(values), str))
	}

	for i := range nameExpr {
		if nameExpr[i].isByte {
			e := p.parseExprNodes(values[i])
			e.(*goast.BasicLit).Kind = token.CHAR
			e.(*goast.BasicLit).Value = fmt.Sprintf("'%c'", e.(*goast.BasicLit).Value[1])
			stmts = append(stmts, &goast.AssignStmt{
				Lhs: []goast.Expr{nameExpr[i].expr},
				Tok: token.ASSIGN,
				Rhs: []goast.Expr{e},
			})
			continue
		}
		stmts = append(stmts, &goast.AssignStmt{
			Lhs: []goast.Expr{nameExpr[i].expr},
			Tok: token.ASSIGN,
			Rhs: []goast.Expr{p.parseExprNodes(values[i])},
		})
	}

	return
}

// Examples:
//  GO TO 30
//  GO TO ( 40, 80 )IEXC
//  GO TO next,(30, 50, 70, 90, 110)
func (p *parser) parseGoto() (stmts []goast.Stmt) {
	p.expect(token.GOTO)

	p.ident++
	if p.ns[p.ident].tok != token.LPAREN {
		//  GO TO 30
		p.foundLabels["Label"+string(p.ns[p.ident].b)] = true
		stmts = append(stmts, &goast.BranchStmt{
			Tok:   token.GOTO,
			Label: goast.NewIdent("Label" + string(p.ns[p.ident].b)),
		})
		p.ident++
		return
	}
	//  GO TO ( 40, 80, 100 )IEXC
	//  GO TO ( 40 )IEXC

	// parse labels
	p.expect(token.LPAREN)
	var labelNames []string
	for ; p.ident < len(p.ns); p.ident++ {
		var out bool
		switch p.ns[p.ident].tok {
		case token.LPAREN:
			// do nothing
		case token.RPAREN:
			out = true
		case token.COMMA:
			// do nothing
		default:
			labelNames = append(labelNames, string(p.ns[p.ident].b))
			p.foundLabels["Label"+string(p.ns[p.ident].b)] = true
		}
		if out {
			break
		}
	}
	p.expect(token.RPAREN)
	p.ident++

	// ignore COMMA
	if p.ns[p.ident].tok == token.COMMA {
		p.ident++
	}

	if len(labelNames) == 0 {
		panic("Not acceptable amount of labels in GOTO")
	}

	// get expr
	st := p.ident
	for ; p.ident < len(p.ns) && p.ns[p.ident].tok != ftNewLine; p.ident++ {
	}
	// generate Go code
	var sw goast.SwitchStmt
	sw.Tag = p.parseExpr(st, p.ident)
	sw.Body = &goast.BlockStmt{}
	for i := 0; i < len(labelNames); i++ {
		sw.Body.List = append(sw.Body.List, &goast.CaseClause{
			List: []goast.Expr{goast.NewIdent(strconv.Itoa(i + 1))},
			Body: []goast.Stmt{&goast.BranchStmt{
				Tok:   token.GOTO,
				Label: goast.NewIdent("Label" + labelNames[i]),
			}},
		})
	}

	stmts = append(stmts, &sw)

	return
}

//  PARAMETER ( ONE = ( 1.0E+0 , 0.0E+0 )  , ZERO = 0.0E+0 )
//  PARAMETER ( LV = 2 )
func (p *parser) parseParameter() (stmts []goast.Stmt) {
	p.expect(ftParameter)
	p.ident++
	p.expect(token.LPAREN)
	// parse values
	var names [][]node
	counter := 1
	p.ident++
	for ; p.ident < len(p.ns); p.ident++ {
		if p.ns[p.ident].tok == token.LPAREN {
			counter++
		}
		if p.ns[p.ident].tok == token.RPAREN {
			counter--
		}
		if p.ns[p.ident].tok == token.RPAREN && counter == 0 {
			break
		}
		if p.ns[p.ident].tok == token.COMMA && counter == 1 {
			names = append(names, []node{})
			continue
		}
		if len(names) == 0 {
			names = append(names, []node{})
		}
		names[len(names)-1] = append(names[len(names)-1], p.ns[p.ident])
	}

	// split to name and value
	for _, val := range names {
		for i := 0; i < len(val); i++ {
			if val[i].tok == token.ASSIGN {
				// add parameters in parser
				p.constants[nodesToString(val[:i])] = val[i+1:]
			}
		}
	}

	p.expect(token.RPAREN)
	p.ident++
	return
}

func (p *parser) parseAssign() (stmts []goast.Stmt) {
	p.expect(ftAssign)
	p.ident++

	statement := string(p.ns[p.ident].b)
	p.ident++

	// ignore TO
	p.ident++

	intVar := string(p.ns[p.ident].b)
	p.ident++
	stmts = append(stmts, &goast.ExprStmt{
		X: goast.NewIdent("// ASSIGN " + statement + " TO " + intVar),
	}, &goast.AssignStmt{
		Lhs: []goast.Expr{goast.NewIdent(intVar)},
		Tok: token.ASSIGN,
		Rhs: []goast.Expr{goast.NewIdent(statement)},
	}, &goast.AssignStmt{
		Lhs: []goast.Expr{goast.NewIdent("_")},
		Tok: token.ASSIGN,
		Rhs: []goast.Expr{goast.NewIdent(intVar)},
	})

	return
}
