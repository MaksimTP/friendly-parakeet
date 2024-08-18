package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var months []string = []string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}

type Options struct {
	k int // []int
	n bool
	r bool
	u bool
	M bool
	b bool
	c bool
	h bool
}

func parseArgs() (Options, []string) {
	kPtr := flag.Int("k", 0, "column to sort")
	nPtr := flag.Bool("n", false, "sort by number")
	rPtr := flag.Bool("r", false, "reverse sort")
	uPtr := flag.Bool("u", false, "unique rows")
	MPtr := flag.Bool("M", false, "sort by month")
	bPtr := flag.Bool("b", false, "ignore tail spaces")
	hPtr := flag.Bool("c", false, "")

	flag.Parse()

	files := flag.Args()

	if len(files) == 0 {
		log.Fatalln("error: no files provided")
	}

	opts := Options{k: *kPtr, n: *nPtr, r: *rPtr, u: *uPtr, M: *MPtr, b: *bPtr, h: *hPtr}

	return opts, files
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

func isSorted(data []string, reversed bool) bool {

}

func isMonth(s string) bool {
	if len(s) < 3 {
		return false
	}
	for _, month := range months {
		if s[:3] == month {
			return true
		}
	}
	return false
}

func isNumeric(s string) bool {
	ss := strings.Split(s, " ")
	_, err := strconv.Atoi(ss[0])
	return err == nil
}

func isHumanNumeric(s string) bool {

}

func isDuplicate(arr []string, s string) bool {
	for _, line := range arr {
		if line == s {
			return true
		}
	}
	return false
}

func sort(data []string, opts Options) []string {

}

func main() {
	opts, files := parseArgs()
	fmt.Println(opts, files)
	data := make([]string, 0)
	if len(files) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			data = append(data, scanner.Text())
		}
	} else {
		for _, file := range files {
			data, err := readFile(file)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
