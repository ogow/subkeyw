package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	i := inputStdin()
	fmt.Println(i)
	data := readInput(i)

	for _, d := range data {
		sep := strings.Split(d, ".")
		sep = sep[:len(sep)-2]

		for _, e := range sep {
			fmt.Println(e)
			sep2 := strings.Split(e, "-")

			for _, ee := range sep2 {
				fmt.Println(ee)
			}
		}
	}
}

func readInput(stdin bool) []string {
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
		if len(os.Args) < 2 {
			fmt.Println("you need to provide a file as argument or trough stdin")
			fmt.Println("e.g.")
			fmt.Println("\tcat file.txt | subkeyw")
			fmt.Println("\tsubkeyw file.txt")
			os.Exit(1)
		}
		arg := os.Args[1]
		f, err := os.Open(arg)
		if err != nil {
			log.Fatal("os.Open()", err)
		}
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
