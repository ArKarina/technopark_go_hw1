package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ArKarina/technopark_go_hw1/unique"
)

func parseOptions(ostream []string) (options unique.Options, inputFile, outputFile string) {
	flag.BoolVar(&options.Count, "c", false, "count the number of repetitions of a string")
	flag.BoolVar(&options.Duplicate, "d", false, "output duplicate lines")
	flag.BoolVar(&options.Unique, "u", false, "output unique lines")
	flag.IntVar(&options.SkipFields, "f", 0, "skip n fields")
	flag.IntVar(&options.SkipChars, "s", 0, "skip n characters")
	flag.BoolVar(&options.Ignore, "i", false, "ignore case")

	flag.Parse()

	inputFile = flag.Arg(0)
	outputFile = flag.Arg(1)

	return options, inputFile, outputFile
}

func readInfo(inputFile string) (buf []string, err error) {
	var r io.Reader

	if inputFile != "" {
		f, err := os.Open(inputFile)
		if err != nil {
			return buf, errors.New("error open file")
		}
		defer f.Close()
		r = f
	} else {
		r = os.Stdin
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		buf = append(buf, scanner.Text())
	}

	return buf, nil
}

func writeInfo(buf []string, outputFile string) error {
	var w io.Writer

	if outputFile != "" {
		f, err := os.Create(outputFile)
		if err != nil {
			return errors.New("error create file")
		}
		defer f.Close()
		w = f
	} else {
		w = os.Stdout
	}

	for _, elem := range buf {
		if _, err := w.Write([]byte(elem + "\n")); err != nil {
			return errors.New("error write to file")
		}
	}

	return nil
}

func main() {
	options, inputFile, outputFile := parseOptions(os.Args)

	if !options.Correct() {
		fmt.Println(unique.Info)
		return
	}

	buf, err := readInfo(inputFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	resBuf := unique.Unique(buf, options)

	if err = writeInfo(resBuf, outputFile); err != nil {
		log.Fatal(err)
		return
	}
}
