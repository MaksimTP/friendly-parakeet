package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/akamensky/argparse"
)

// !DONE
// Реализовать утилиту аналог консольной команды cut (man cut).
// Утилита должна принимать строки через STDIN, разбивать по
// разделителю (TAB) на колонки и выводить запрошенные.

type Options struct {
	d string // -d - "delimiter" - использовать другой разделитель
	f []int  // -f - "fields" - выбрать поля (колонки)
	s bool   // -s - "separated" - только строки с разделителем
}

func parseArgs() Options {
	parser := argparse.NewParser("cut", "this is a cut program")

	d := parser.String("d", "delimiter", &argparse.Options{Default: "\t"})
	f := parser.IntList("f", "fields", nil)
	s := parser.Flag("s", "separated", nil)
	parser.Parse(os.Args)
	return Options{d: *d, f: *f, s: *s}
}

func cut(s string, opts Options) (res string) {
	arrS := strings.Split(s, opts.d)
	if opts.s && len(arrS) == 1 {
		return ""
	}

	for i, field := range opts.f {
		indx := field - 1
		if indx > len(arrS)-1 {
			continue
		}
		res += arrS[indx]
		if i == len(opts.f)-1 {
			res += "\n"
		} else {
			res += opts.d
		}
	}

	return
}

func main() {
	opts := parseArgs()
	fmt.Println(opts)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := scanner.Text()
		res := cut(data, opts)
		fmt.Print(res)
	}

}
