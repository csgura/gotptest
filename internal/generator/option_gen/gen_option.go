package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"

	"gotptest/internal/maxproduct"
)

func typeArgs(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "A%d", j)
	}
	return f.String()
}

func main() {
	f := &bytes.Buffer{}

	fmt.Fprintf(f, "package option\n\n")

	fmt.Fprintln(f, `
import (
	"gotptest/fp"
	"gotptest/curried"
)`)

	for i := 2; i < maxproduct.MaxProduct; i++ {
		fmt.Fprintf(f, "type ApplicativeFunctor%d [HT, ", i)

		for j := 1; j <= i; j++ {
			fmt.Fprintf(f, "A%d,", j)
		}
		fmt.Fprintf(f, "R any]")

		fmt.Fprintf(f, " struct {\n")
		fmt.Fprintf(f, "  h fp.Option[HT]\n")
		fmt.Fprintf(f, "  fn fp.Option[")
		endBracket := "]"
		for j := 1; j <= i; j++ {
			fmt.Fprintf(f, "fp.Func1[A%d, ", j)
			endBracket = endBracket + "]"
		}
		fmt.Fprintf(f, "R%s\n", endBracket)

		fmt.Fprintf(f, "}\n")

		typeparams := "[HT,"

		nexttp := "["
		for j := 1; j <= i; j++ {
			typeparams = fmt.Sprintf("%s A%d,", typeparams, j)
			nexttp = fmt.Sprintf("%s A%d,", nexttp, j)

		}
		typeparams = typeparams + " R]"
		nexttp = nexttp + " R]"

		receiver := fmt.Sprintf("func (r ApplicativeFunctor%d%s)", i, typeparams)

		fmt.Fprintf(f, "%s Shift() ApplicativeFunctor%d[HT,%s,A1,R] {\n", receiver, i, typeArgs(2, i))
		fmt.Fprintf(f, `
	nf := fp.Compose(curried.Revert%d[%s, R], fp.Compose(fp.Func%d[%s, R].Shift, fp.Func%d[%s, A1, R].Curried))
	return ApplicativeFunctor%d[HT,%s, A1, R]{
		r.h,
		Map(r.fn, nf),
	}

}
`, i, typeArgs(1, i), i, typeArgs(1, i), i, typeArgs(2, i), i, typeArgs(2, i))

		fmt.Fprintf(f, "%s ApOption( a fp.Option[A1]) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
		fmt.Fprintf(f, `
	

	return ApplicativeFunctor%d%s{a, Ap(r.fn, a)}
}
`, i-1, nexttp)

		fmt.Fprintf(f, "%s Ap( a A1) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
		fmt.Fprintln(f, `
	return r.ApOption(Some(a))

}`)

		// 		fmt.Fprintf(f, "%s FlatMap( a func(HT) fp.Option[A1]) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
		// 		fmt.Fprintln(f, `
		// 	av := FlatMap(r.h, func(v HT) fp.Option[A1] {
		// 		return a(v)
		// 	})
		// 	return r.ApOption(av)
		// }`)

		// 		fmt.Fprintf(f, "%s Map( a func(HT) A1) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
		// 		fmt.Fprintln(f, `
		// 	return r.FlatMap(func(h HT) fp.Option[A1] {
		// 		return Some(a(h))
		// 	})
		// }`)

		tpr := ""
		for j := 1; j <= i; j++ {
			if j != 1 {
				tpr = tpr + ","
			}
			tpr = tpr + fmt.Sprintf("A%d", j)
		}
		tpr = tpr + ",R"

		fmt.Fprintf(f, "func Applicative%d[%s any](fn fp.Func%d[%s]) ApplicativeFunctor%d[fp.Unit, %s] {\n", i, tpr, i, tpr, i, tpr)
		fmt.Fprintf(f, "    return ApplicativeFunctor%d[fp.Unit, %s]{ None[fp.Unit](),Some(fn.Curried())}\n", i, tpr)
		fmt.Fprintf(f, "}\n")

		// for j := 1; j <= i; j++ {
		// 	if j != 1 {
		// 		fmt.Fprintf(f, ",")
		// 	}
		// 	fmt.Fprintf(f, "a%d A%d", j, j)
		// }
		// fmt.Fprintf(f, ") R\n\n")
	}

	formatted, err := format.Source(f.Bytes())
	if err != nil {
		log.Print(f.String())
		log.Fatal("format error ", err)

		return
	}

	err = ioutil.WriteFile("applicative_gen.go", formatted, 0666)
	if err != nil {
		return
	}
}
