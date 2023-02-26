package subkeyw

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

	dashes := flag.Bool("dashes", false, "include words with dashes, otherwise these are removed")
	file := flag.String("file", "", "takes a file with a list of subdomains")

	flag.Parse()

	var all []string
	i := inputIsStdin()

	if *file == "" && i == false {
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
	file := os.Stdin
	fi, err := file.Stat()
	if err != nil {
		log.Fatal("file.Stat()", err)
	}
	size := fi.Size()
	if size > 0 {
		// from stdin
		return true
	} else {
		// not from stdin
		return false
	}

}
