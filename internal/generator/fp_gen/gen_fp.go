package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"log"

	"gotptest/internal/maxproduct"
)

func curriedType(start, until int) string {
	f := &bytes.Buffer{}
	endBracket := ""
	for j := start; j <= until; j++ {
		fmt.Fprintf(f, "Func1[A%d, ", j)
		endBracket = endBracket + "]"
	}
	fmt.Fprintf(f, "R%s", endBracket)

	return f.String()
}

func funcTypeArgs(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "A%d", j)
	}
	return f.String()
}

func tupleTypeArgs(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "T%d", j)
	}
	return f.String()
}

func formatStr(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ",")
		}
		fmt.Fprintf(f, "%s", "%v")
	}
	return f.String()
}

func tupleArgs(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "r.I%d", j)
	}
	return f.String()
}

func consType(start, until int) string {
	ret := "hlist.Nil"
	for j := until; j >= start; j-- {
		ret = fmt.Sprintf("hlist.Cons[T%d, %s]", j, ret)
	}
	return ret
}

func funcDeclArgs(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "a%d A%d", j, j)
	}
	return f.String()
}

func generate(packname string, filename string, writeFunc func(w io.Writer)) {
	f := &bytes.Buffer{}

	fmt.Fprintf(f, "package %s\n\n", packname)
	writeFunc(f)

	formatted, err := format.Source(f.Bytes())
	if err != nil {
		log.Print(f.String())
		log.Fatal("format error ", err)

		return
	}

	err = ioutil.WriteFile(filename, formatted, 0644)
	if err != nil {
		return
	}
}

func funcCallArgs(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "a%d", j)
	}
	return f.String()
}

func main() {
	generate("fp", "func_gen.go", func(f io.Writer) {
		for i := 2; i < maxproduct.MaxProduct; i++ {
			fmt.Fprintf(f, "type Func%d", i)
			fmt.Fprintf(f, "[")

			for j := 1; j <= i; j++ {
				if j != 1 {
					fmt.Fprintf(f, ",")
				}
				fmt.Fprintf(f, "A%d", j)
			}
			fmt.Fprintf(f, ",R any]")

			fmt.Fprintf(f, "func( ")

			for j := 1; j <= i; j++ {
				if j != 1 {
					fmt.Fprintf(f, ",")
				}
				fmt.Fprintf(f, "a%d A%d", j, j)
			}
			fmt.Fprintf(f, ") R\n\n")

		}
		for i := 2; i < maxproduct.MaxProduct; i++ {

			fmt.Fprintf(f, `
func(r Func%d[%s,R]) Curried() %s {
	return func(a1 A1) %s {
		return Func%d[%s,R](func(%s) R {
			return r(%s)
		}).Curried()
	}	
}
`, i, funcTypeArgs(1, i), curriedType(1, i), curriedType(2, i), i-1, funcTypeArgs(2, i), funcDeclArgs(2, i), funcCallArgs(1, i))

			fmt.Fprintf(f, `
func(r Func%d[%s,R]) Shift() Func%d[%s,A1,R] {
	return func(%s , a1 A1) R {
		return r(%s)
	}	
}
`, i, funcTypeArgs(1, i), i, funcTypeArgs(2, i), funcDeclArgs(2, i), funcCallArgs(1, i))

		}
	})

}
