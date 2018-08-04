// +build integration

package main

import (
	"bytes"
	"go/format"
	"go/token"
	"io/ioutil"
	"os/exec"
	"testing"

	"github.com/Konstantin8105/f4go/fortran"
)

func TestIntegration(t *testing.T) {
	// gfortran ./testdata/main.f -o ./testdata/a.out
	out, err := exec.Command(
		"gfortran",
		"./testdata/main.f",
		"-o", "./testdata/a.out",
	).CombinedOutput()
	if err != nil {
		t.Fatalf("Cannot compile by gfortran: %v\n%s", err, out)
	}

	// ./testdata/a.out
	fortranOutput, err := exec.Command(
		"./testdata/a.out",
	).CombinedOutput()
	if err != nil {
		t.Fatalf("Cannot fortran executable file : %v\n%s", err, fortranOutput)
	}

	t.Logf("Fortran output:\n%s\n", fortranOutput)
	t.Logf("fortran source is ok")

	// parsing to Go code
	dat, err := ioutil.ReadFile("./testdata/main.f")
	if err != nil {
		t.Fatalf("Cannot fortran source: %v", err)
	}
	ast, errs := fortran.Parse(dat)
	if len(errs) > 0 {
		for _, err := range errs {
			t.Logf("Error: %v", err)
		}
		t.Fatal("Errors is more zero")
	}

	var buf bytes.Buffer
	if err = format.Node(&buf, token.NewFileSet(), &ast); err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile("./testdata/g.go", buf.Bytes(), 0644)
	if err != nil {
		t.Fatalf("Cannot write Go source: %v", err)
	}

	// run Go code
	goOutput, err := exec.Command(
		"go", "run", "./testdata/g.go",
	).CombinedOutput()
	if err != nil {
		t.Fatalf("Cannot go executable file : %v\n%s", err, goOutput)
	}

	if !bytes.Equal(fortranOutput, goOutput) {
		t.Errorf("Results is not same: `%v` != `%v`",
			string(fortranOutput),
			string(goOutput))
	}
}
