package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

var (
	linenum_flag    bool // +
	fixed_flag      bool // +
	invert_flag     bool // +
	ignorecase_flag bool
	count_flag      bool // +
	context_flag    int
	after_flag      int
	before_flag     int
)

func GrepUse(pattern string, filename string) {
	if context_flag > 0 {
		after_flag = context_flag
		before_flag = context_flag
	}

	var fileobj *os.File

	if filename != "" {
		var err error
		fileobj, err = os.Open(filename)
		if err != nil {
			fmt.Println("no such file or catalog")
			return
		}
		defer fileobj.Close()
	} else {
		fileobj = os.Stdin
	}

	scanner := bufio.NewScanner(fileobj)
	i := 0
	count := 0

	var history []string
	set := make(map[int]bool)

	for scanner.Scan() {
		history = append(history, scanner.Text())

		var matched bool
		if fixed_flag {
			matched = strings.Contains(scanner.Text(), pattern)
		} else {
			matched, _ = regexp.MatchString(pattern, scanner.Text())
		}
		i++

		if matched && !invert_flag || !matched && invert_flag {
			if count_flag {
				count++
				continue
			}

			for k := i - before_flag; k <= i+after_flag; k++ {
				set[k] = true
			}

		}
	}

	str_number := make([]int, len(set))

	for key := range set {
		str_number = append(str_number, key)
	}

	sort.Ints(str_number)

	for _, num := range str_number {
		if num-1 < 0 {
			continue
		}
		if linenum_flag {
			fmt.Printf("%d", num)
		}
		fmt.Println(history[num-1])
	}
}

func main() {
	flag.BoolVar(&linenum_flag, "n", false, "")
	flag.BoolVar(&fixed_flag, "F", false, "")
	flag.BoolVar(&invert_flag, "v", false, "")
	flag.BoolVar(&ignorecase_flag, "i", false, "")
	flag.BoolVar(&count_flag, "c", false, "")
	flag.IntVar(&context_flag, "C", 0, "")
	flag.IntVar(&after_flag, "a", 0, "")
	flag.IntVar(&before_flag, "b", 0, "")

	flag.Parse()

	GrepUse(flag.Arg(0), flag.Arg(1))
}
