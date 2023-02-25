package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dvvedz/uniques"
)

func main() {

	dashes := flag.Bool("dashes", false, "include words with dashes, otherwise it removes them from output")
	file := flag.String("file", "", "takes a file with a list of subdomains")

	flag.Parse()

	var all []string
	i := inputStdin()

	if len(*file) < 1 && i == false {
		flag.Usage()
		os.Exit(1)
	}

	data := readInput(i, *file)

	for _, d := range data {
		sep := strings.Split(d, ".")
		sep = sep[:len(sep)-2]

		for _, e := range sep {
			if *dashes {
				all = append(all, e)
			}
			sep2 := strings.Split(e, "-")
			for _, ee := range sep2 {
				all = append(all, ee)
			}
		}
	}

	us := uniques.StringSlice(all)
	for _, e := range us {
		fmt.Println(e)
	}
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
		// if len(file) {
		// fmt.Println("you need to provide a file as argument or trough stdin")
		// fmt.Println("")
		// flag.Usage()
		// os.Exit(1)
		// }
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

func inputStdin() bool {
	file := os.Stdin
	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat()", err)
	}
	size := fi.Size()
	if size > 0 {
		return true
	} else {
		return false
	}
}

// func printHelp() {
// fmt.Println("Usage: subkeyw file.txt")
// fmt.Println("Flags:")
// fmt.Println("-h\t Help menu")
// fmt.Println("-dashes\t inlcude words with dashes")
// fmt.Println("e.g.")
// fmt.Println("\tcat file.txt | subkeyw")
// fmt.Println("\tsubkeyw file.txt")
// os.Exit(1)
// }
