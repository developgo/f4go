package fortran

import (
	"bytes"
	"fmt"
	goast "go/ast"
	goparser "go/parser"
	"go/token"
	"strconv"
)

// Example:
//  WRITE ( * , FMT = 9999 ) SRNAME ( 1 : LEN_TRIM ( SRNAME ) ) , INFO
//  9999 FORMAT ( ' ** On entry to ' , A , ' parameter number ' , I2 , ' had ' , 'an illegal value' )
//
// write (*, '(I1,A2,I1)') i,'YY',i
func (p *parser) parseWrite() (stmts []goast.Stmt) {

	p.expect(ftWrite)
	p.ident++
	p.expect(token.LPAREN)
	p.ident++
	if !(p.ns[p.ident].tok == token.MUL ||
		(p.ns[p.ident].tok == token.INT && string(p.ns[p.ident].b) == "6")) {
		// allowable letters: `*` or `6`
		// value with const 6, so ok
		var isOk bool
		if vv, ok := p.initVars.get(string(p.ns[p.ident].b)); ok {
			if nodesToString(vv.constant) == "6" {
				isOk = true
				stmts = append(stmts, &goast.AssignStmt{
					Lhs: []goast.Expr{goast.NewIdent("_")},
					Tok: token.ASSIGN,
					Rhs: []goast.Expr{goast.NewIdent(vv.name)},
				})
			}
		}
		if !isOk {
			goto externalFunc
		}
	}

	p.ident++
	p.expect(token.COMMA)
	p.ident++

	// From:
	//  WRITE(6,10) A
	// To:
	//  WRITE(6,FMT=10) A
	if p.ns[p.ident].tok == token.INT {
		p.ns = append(p.ns[:p.ident], append([]node{
			{tok: token.IDENT, b: []byte("FMT")},
			{tok: token.ASSIGN, b: []byte("=")},
		}, p.ns[p.ident:]...)...)
	}

	if p.ns[p.ident].tok == token.IDENT &&
		bytes.Equal(bytes.ToUpper(p.ns[p.ident].b), []byte("FMT")) {

		p.ident++
		p.expect(token.ASSIGN)
		p.ident++
		p.expect(token.INT)
		line := p.getLineByLabel(p.ns[p.ident].b)
		fs := p.parseFormat(line[2:])
		p.addImport("fmt")
		p.ident++
		p.expect(token.RPAREN)
		p.ident++
		// separate to expression by comma
		exprs := p.scanWriteExprs()
		p.expect(ftNewLine)
		var args []goast.Expr
		args = append(args, goast.NewIdent(fs))
		args = append(args, exprs...)
		stmts = append(stmts, &goast.ExprStmt{
			X: &goast.CallExpr{
				Fun: &goast.SelectorExpr{
					X:   goast.NewIdent("fmt"),
					Sel: goast.NewIdent("Printf"),
				},
				Lparen: 1,
				Args:   args,
			},
		})
	} else if p.ns[p.ident].tok == token.MUL {
		p.expect(token.MUL)
		p.ident++
		p.expect(token.RPAREN)
		p.ident++
		exprs := p.scanWriteExprs()
		p.expect(ftNewLine)
		var format string
		format = "\""
		for i := 0; i < len(exprs); i++ {
			format += " %s"
		}
		format += "\\n\""
		stmts = append(stmts, &goast.ExprStmt{
			X: &goast.CallExpr{
				Fun: &goast.SelectorExpr{
					X:   goast.NewIdent("fmt"),
					Sel: goast.NewIdent("Printf"),
				},
				Lparen: 1,
				Args:   append([]goast.Expr{goast.NewIdent(format)}, exprs...),
			},
		})
	} else if p.ns[p.ident].tok == token.STRING {
		// write (*, '(I1,A2,I1)') i,'YY',i
		format := p.ns[p.ident].b
		toks := scan(format[2 : len(format)-2])
		fs := p.parseFormat(toks)
		p.ident++
		p.addImport("fmt")
		p.expect(token.RPAREN)
		p.ident++
		// separate to expression by comma
		exprs := p.scanWriteExprs()
		p.expect(ftNewLine)
		var args []goast.Expr
		args = append(args, goast.NewIdent(fs))
		args = append(args, exprs...)
		stmts = append(stmts, &goast.ExprStmt{
			X: &goast.CallExpr{
				Fun: &goast.SelectorExpr{
					X:   goast.NewIdent("fmt"),
					Sel: goast.NewIdent("Printf"),
				},
				Lparen: 1,
				Args:   args,
			},
		})
	}
	return

externalFunc:

	p.ident--
	p.expect(token.LPAREN)
	// WRITE ( 1, *) R
	//       ========= this out
	args, end := separateArgsParen(p.ns[p.ident:])
	p.ident += end

	// Pattern:
	//  WRITE( UNIT = ..., FMT = ...)
	// Other parameters are ignored

	// Part : UNIT
	unit := string(args[0][0].b)
	if len(args[0]) == 3 {
		unit = string(args[0][2].b)
	}

	// Part: FMT
	fmts := args[1][0]
	if len(args[1]) == 3 {
		fmts = args[1][2]
	}

	var fs string
	if fmts.tok == token.INT {
		line := p.getLineByLabel(fmts.b)
		fs = p.parseFormat(line[2:])
	} else {
		// Example :
		// '(A80)'
		ns := scan(fmts.b[2 : len(fmts.b)-2])
		fs = p.parseFormat(ns)
	}

	// TRANSA , TRANSB , SAME , ERR
	for end = p.ident; p.ns[end].tok != ftNewLine; end++ {
	}

	s := fmt.Sprintf("intrinsic.WRITE(%s,%s,%s)", unit, fs,
		nodesToString(p.ns[p.ident:end]))

	p.addImport("github.com/Konstantin8105/f4go/intrinsic")

	ast, err := goparser.ParseExpr(s)
	if err != nil {
		panic(err)
	}
	stmts = append(stmts, &goast.ExprStmt{
		X: ast,
	})

	p.ident = end

	return
}

func (p *parser) scanWriteExprs() (exprs []goast.Expr) {
	st := p.ident
	for ; p.ns[p.ident].tok != ftNewLine; p.ident++ {
		for ; ; p.ident++ {
			if p.ns[p.ident].tok == token.COMMA || p.ns[p.ident].tok == ftNewLine {
				break
			}
			if p.ns[p.ident].tok != token.LPAREN {
				continue
			}
			counter := 0
			for ; ; p.ident++ {
				if p.ns[p.ident].tok == token.RPAREN {
					counter--
				}
				if p.ns[p.ident].tok == token.LPAREN {
					counter++
				}
				if counter != 0 {
					continue
				}
				break
			}
		}
		// parse expr
		exprs = append(exprs, p.parseExpr(st, p.ident))
		st = p.ident + 1
		if p.ns[p.ident].tok == ftNewLine {
			p.ident--
		}
	}
	return
}

func (p *parser) getLineByLabel(label []byte) (fs []node) {

	// memorization of FORMAT lines
	if v, ok := p.formats[string(label)]; ok {
		return v
	}

	var found bool
	var st int
	for st = p.ident; st < len(p.ns); st++ {
		if p.ns[st-1].tok == ftNewLine && bytes.Equal(p.ns[st].b, label) {
			found = true
			break
		}
	}
	if !found {
		p.addError("Cannot found label :" + string(label))
		return
	}

	for i := st; i < len(p.ns) && p.ns[i].tok != ftNewLine; i++ {
		fs = append(fs, p.ns[i])
		// remove line
		p.ns[i].tok, p.ns[i].b = ftNewLine, []byte("\n")
	}

	p.formats[string(label)] = fs

	return
}

func (p *parser) parseFormat(in []node) (s string) {
	var fs []node
	fs = append(fs, in...)
	if len(fs) == 0 {
		s = "\"\\n\""
		return
	}
	// From:
	//  ... / ...
	// To:
	//  ... , "\\n" , ...
	for i := 0; i < len(fs); i++ {
		if fs[i].tok != token.QUO { // not /
			continue
		}
		fs = append(fs[:i], append([]node{
			{tok: token.COMMA, b: []byte(",")},
			{tok: token.STRING, b: []byte("\\n")},
			{tok: token.COMMA, b: []byte(",")},
		}, fs[i+1:]...)...)
	}

	for i := 0; i < len(fs); i++ {
		f := fs[i]
		switch f.tok {
		case token.IDENT:
			switch f.b[0] {
			case 'I':
				s += "%" + string(f.b[1:]) + "d"
			case 'F':
				s += "%" + string(f.b[1:])
				if i+1 < len(fs) && fs[i+1].tok == token.PERIOD {
					i += 1
					s += "."
					if i+1 < len(fs) && fs[i+1].tok == token.INT {
						s += string(fs[i+1].b)
						i += 1
					}
				}
				s += "f"
			case 'E', 'D':
				s += "%" + string(f.b[1:])
				if i+1 < len(fs) && fs[i+1].tok == token.PERIOD {
					i += 1
					s += "."
					if i+1 < len(fs) && fs[i+1].tok == token.INT {
						s += string(fs[i+1].b)
						i += 1
					}
				}
				s += "E"
			case 'A':
				if len(f.b) > 1 {
					s += "%" + string(f.b[1:]) + "s"
				} else {
					s += "%s"
				}
			default:
				p.addError("Not support format : " + string(f.b))
			}
		case token.INT:
			// 1X
			// 5X
			v, _ := strconv.Atoi(string(f.b))
			if fs[i+1].b[0] == 'X' {
				for i := 0; i < v-1; i++ {
					s += " "
				}
				i++
			}

		case token.STRING:
			str := string(f.b)
			if str[0] == '"' || str[0] == '\'' {
				s += str[1 : len(str)-1]
			} else {
				s += str
			}
		case token.COMMA, token.LPAREN, token.RPAREN:
			// ignore
		default:
			s += "%v"
		}
	}
	return "\"" + s + "\\n\""
}

// Example:
//  READ ( NIN , FMT = * ) TSTERR
//  READ ( NIN , FMT = * ) THRESH
//  READ ( NIN , FMT = * ) ( IDIM ( I ) , I = 1 , NIDIM )
func (p *parser) parseRead() (stmts []goast.Stmt) {
	p.expect(ftRead)
	p.ident++
	p.expect(token.LPAREN)

	args, end := separateArgsParen(p.ns[p.ident:])
	p.ident += end

	// Pattern:
	//  READ ( NIN , FMT = * ) TS
	// Other parameters are ignored

	// Part : UNIT
	unit := string(args[0][0].b)
	if len(args[0]) == 3 {
		unit = string(args[0][2].b)
	}

	// Part: FMT
	fmts := args[1][0]
	if len(args[1]) == 3 {
		fmts = args[1][2]
	}

	var fs string
	if fmts.tok == token.INT {
		line := p.getLineByLabel(fmts.b)
		fs = p.parseFormat(line[2:])
	} else if fmts.tok == token.MUL {
		// Example: *
		fs = "\"%v\""
	} else {
		// Example :
		// '(A80)'
		ns := scan(fmts.b[2 : len(fmts.b)-2])
		fs = p.parseFormat(ns)
	}

	// TRANSA , TRANSB , SAME , ERR
	for end = p.ident; p.ns[end].tok != ftNewLine; end++ {
	}

	s := fmt.Sprintf("intrinsic.READ(%s,%s,%s)", unit, fs,
		nodesToString(p.ns[p.ident:end]))

	p.addImport("github.com/Konstantin8105/f4go/intrinsic")

	ast, err := goparser.ParseExpr(s)
	if err != nil {
		panic(fmt.Errorf("%s:%v", s, err))
	}
	stmts = append(stmts, &goast.ExprStmt{
		X: ast,
	})

	p.ident = end

	return
}

// Example:
//  OPEN ( NTRA , FILE = SNAPS )
//  OPEN ( NOUT , FILE = SUMMRY , STATUS = 'UNKNOWN' )
//  OPEN ( UNIT = 2 , FILE = "./testdata/main.f" )
func (p *parser) parseOpen() (stmts []goast.Stmt) {
	p.expect(ftOpen)
	p.ident++
	p.expect(token.LPAREN)
	args, end := separateArgsParen(p.ns[p.ident:])
	p.ident += end

	// Pattern:
	//  OPEN( UNIT = ..., FILE = ...)
	// Other parameters are ignored

	// Part : UNIT
	unit := string(args[0][0].b)
	if len(args[0]) == 3 {
		unit = string(args[0][2].b)
	}

	// Part : FILE
	file := string(args[1][0].b)
	if len(args[1]) == 3 {
		file = string(args[1][2].b)
	}

	s := fmt.Sprintf("intrinsic.OPEN(%s,%s)", unit, file)
	p.addImport("github.com/Konstantin8105/f4go/intrinsic")
	ast, err := goparser.ParseExpr(s)
	if err != nil {
		panic(err)
	}
	stmts = append(stmts, &goast.ExprStmt{
		X: ast,
	})

	return
}

// Example:
//  CLOSE ( 2 )
//  CLOSE ( NIN )
func (p *parser) parseClose() (stmts []goast.Stmt) {
	p.expect(ftClose)
	p.ident++
	p.expect(token.LPAREN)
	args, end := separateArgsParen(p.ns[p.ident:])
	p.ident += end

	s := fmt.Sprintf("intrinsic.CLOSE(%s)", string(args[0][0].b))
	p.addImport("github.com/Konstantin8105/f4go/intrinsic")
	ast, err := goparser.ParseExpr(s)
	if err != nil {
		panic(err)
	}
	stmts = append(stmts, &goast.ExprStmt{
		X: ast,
	})

	return
}