package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func SortFileString(filename string, reverse_flag bool,
	unique_flag bool, nsort_flag bool, column int) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("no such file or catalog")
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var array []string

	for scanner.Scan() {
		array = append(array, scanner.Text())
	}

	if nsort_flag {
		sort.Slice(array, func(i, j int) bool {
			vi, _ := strconv.Atoi(data[i])
			vj, _ := strconv.Atoi(data[j])
			return vi < vj
		})
	} else {
		sort.Strings(array)
	}

	if reverse_flag {
		for i := len(array) - 1; i >= 0; i-- {
			fmt.Println(array[i])
		}
	} else {
		for i := 0; i < len(array); i++ {
			fmt.Println(array[i])
		}
	}
}

func main() {
	reverse_flag := flag.Bool("r", false, "")
	unique_flag := flag.Bool("u", false, "")
	nsort_flag := flag.Bool("n", false, "")
	column_flag := flag.Int("k", 0, "")

	flag.Parse()

	SortFileString(flag.Arg(0), *reverse_flag, *unique_flag, *nsort_flag, *column_flag)
}
