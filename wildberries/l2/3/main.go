package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/akamensky/argparse"
)

var months []string = []string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}

type Options struct {
	k []int
	n bool
	r bool
	u bool
	M bool
	b bool
	c bool
	h bool
}

// -k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
//
//	-n — сортировать по числовому значению
//	-r — сортировать в обратном порядке
//	-u — не выводить повторяющиеся строки
//	Дополнительно
//	Реализовать поддержку утилитой следующих ключей:
//	-M — сортировать по названию месяца
//	-b — игнорировать хвостовые пробелы
//	-c — проверять отсортированы ли данные
//	-h — сортировать по числовому значению с учетом суффиксов
func parseArgs() (Options, []string) {
	parser := argparse.NewParser("cut", "this is a cut program")

	kPtr := parser.IntList("k", "key", nil)
	nPtr := parser.Flag("n", "numeric-sort", nil)
	rPtr := parser.Flag("r", "reverse", nil)
	uPtr := parser.Flag("u", "unique", nil)
	MPtr := parser.Flag("M", "month-sort", nil)
	bPtr := parser.Flag("b", "ignore-leading-blanks", nil)
	cPtr := parser.Flag("c", "check", nil)
	hPtr := parser.Flag("h", "human-numeric-sort", nil)

	parser.Parse(os.Args)

	files := make([]*string, 50)

	for i := 0; i < 50; i++ {
		s := parser.StringPositional(nil)
		files = append(files, s)
	}

	err := parser.Parse(os.Args)

	if err != nil {
		log.Fatalln(err)
	}

	returnFiles := make([]string, 0)

	for _, file := range files {
		if file != nil {
			if *file != "" {
				returnFiles = append(returnFiles, *file)
			}
		}
	}

	if len(files) == 0 {
		log.Fatalln("error: no files provided")
	}

	opts := Options{k: *kPtr, n: *nPtr, r: *rPtr, u: *uPtr, M: *MPtr, b: *bPtr, c: *cPtr, h: *hPtr}

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

// func isSorted(data []string, reversed bool) bool {

// }

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

// func isHumanNumeric(s string) bool {

// }

func isDuplicate(arr []string, s string) bool {
	for _, line := range arr {
		if line == s {
			return true
		}
	}
	return false
}

func sort(data []string, opts Options) []string {
	res := []string{}
	for _, line := range data {

	}
	return res
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
		res := sort(data, opts)
		fmt.Println(res)
	} else {
		for _, file := range files {
			data, err := readFile(file)
			if err != nil {
				log.Fatalln(err)
			}
			res := sort(data, opts)
			fmt.Println(res)
		}
	}
}
