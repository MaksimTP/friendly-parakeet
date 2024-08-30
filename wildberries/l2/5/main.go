package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/akamensky/argparse"
)

// TODO: Implement A,B,C
type Options struct {
	A        int      // -A - "after" печатать +N строк после совпадения
	B        int      // -B - "before" печатать +N строк до совпадения
	C        int      // -C - "context" (A+B) печатать ±N строк вокруг совпадения
	c        bool     // -c - "count" (количество строк)
	i        bool     // -i - "ignore-case" (игнорировать регистр)
	v        bool     // -v - "invert" (вместо совпадения, исключать)
	F        bool     // -F - "fixed", точное совпадение со строкой, не паттерн
	n        bool     // -n - "line num", напечатать номер строки
	patterns []string //
}

func parseArgs() (Options, []string) {
	parser := argparse.NewParser("cut", "this is a cut program")

	APtr := parser.Int("A", "after", nil)
	BPtr := parser.Int("B", "before", nil)
	CPtr := parser.Int("C", "context", nil)
	cPtr := parser.Flag("c", "count", nil)
	iPtr := parser.Flag("i", "ignore-case", nil)
	vPtr := parser.Flag("v", "invert", nil)
	FPtr := parser.Flag("F", "fixed", nil)
	nPtr := parser.Flag("n", "line-num", nil)
	files := make([]*string, 50)

	for i := 0; i < 50; i++ {
		s := parser.StringPositional(nil)
		files = append(files, s)
	}

	err := parser.Parse(os.Args)
	returnFiles := make([]string, 0)

	for _, file := range files {
		if file != nil {
			if *file != "" {
				returnFiles = append(returnFiles, *file)
			}
		}
	}

	if err != nil {
		log.Fatalln(err)
	}

	opts := Options{A: *APtr, B: *BPtr, C: *CPtr, c: *cPtr, i: *iPtr, v: *vPtr, F: *FPtr, n: *nPtr}

	return opts, returnFiles
}

func readFile(path string) ([]string, error) {
	fp, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}
	defer fp.Close()
	res := make([]string, 0)
	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}
	return res, nil
}

func grep(data []string, opts Options) {
	matchedCount := 0
	matched := false
	result := make([]string, 0)
	for lineCount, line := range data {
		var lineToAppend string
		matched = false
		changed_line := strings.Clone(line)
		if opts.n {
			lineToAppend += fmt.Sprintf("%d:", lineCount+1)
			// fmt.Printf("%d:", lineCount+1)
		}
		if opts.i {
			changed_line = strings.ToLower(line)
		}
		for _, pattern := range opts.patterns {

			if opts.F {
				matched = pattern == changed_line
			} else {
				matched, _ = regexp.MatchString(pattern, changed_line)
			}
			if opts.v {
				matched = !matched
			}

			if matched && !opts.c && !opts.F {
				lineToAppend += fmt.Sprintf("%s\n", changed_line)
				result = append(result, lineToAppend)
				// fmt.Printf("%s\n", changed_line)
				break
			}
		}
		if matched {
			matchedCount++
		}
	}

	if opts.c {
		fmt.Println(matchedCount)
	} else {
		for _, line := range result {
			fmt.Print(line)
		}
	}
}

func main() {
	opts, files := parseArgs()
	data := make([]string, 0)
	if len(files) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			data = append(data, scanner.Text())
		}
		grep(data, opts)
	} else {
		for _, file := range files {
			data, err := readFile(file)
			if err != nil {
				log.Fatalln(err)
			}
			grep(data, opts)
		}
	}

}
