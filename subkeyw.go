package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	dashes := flag.Bool("dashes", false, "include words with dashes, otherwise these are removed")
	file := flag.String("file", "", "takes a file with a list of subdomains")
	frequency := flag.Bool("frequency", false, "print how many times a word occurred")

	flag.Parse()

	var all []string
	i := inputIsStdin()

	if *file == "" && !i {
		flag.Usage()
		os.Exit(1)
	}

	data := readInput(i, *file)

	for _, d := range data {
		sep := strings.Split(d, ".")
		sep = sep[:len(sep)-2]

		for _, e := range sep {
			e = strings.TrimSpace(e)
			if *dashes {
				all = append(all, e)
			}
			sep2 := strings.Split(e, "-")
			all = append(all, sep2...)
		}
	}

	counts := occurrenceOfElement(all)

	keys := make([]string, 0, len(counts))
	for key := range counts {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool { return counts[keys[i]] < counts[keys[j]] })

	for _, key := range keys {
		if key != "" {
			if *frequency {
				fmt.Println(counts[key], key)
			} else {
				fmt.Println(key)
			}
		}
	}

}

func occurrenceOfElement(arr []string) map[string]int {
	dict := make(map[string]int)
	for _, ele := range arr {
		dict[ele] = dict[ele] + 1
	}
	return dict
}

func readInput(stdin bool, file string) []string {
	var sl []string
	if stdin {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			sl = append(sl, scanner.Text())
		}

		if scanner.Err() != nil {
			log.Fatal("scanner error")
		}
	} else {
		f, err := os.Open(file)
		if err != nil {
			log.Fatal("os.Open()", err)
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			sl = append(sl, scanner.Text())
		}

		if scanner.Err() != nil {
			log.Fatal("scanner error")
		}
	}

	return sl
}

func inputIsStdin() bool {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return true
	} else {
		return false
	}
	// file := os.Stdin
	// fi, err := file.Stat()
	// if err != nil {
	// log.Fatal("file.Stat()", err)
	// }
	// size := fi.Size()
	// if size > 0 {
	// // from stdin
	// return true
	// } else {
	// // not from stdin
	// return false
	// }
}
