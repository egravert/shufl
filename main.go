package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	flag "github.com/spf13/pflag"
)

type config struct {
	echo       bool
	inputRange interval
	head       int
	file       string
}

func main() {
	var cfg config

	flag.BoolVarP(&cfg.echo, "echo", "e", false, "treat each ARG as an input line")
	flag.VarP(&cfg.inputRange, "input-range", "i", "treat each number LO through HI as an input line")
	flag.IntVarP(&cfg.head, "head-count", "n", 0, "Output at most COUNT lines")
	flag.StringVarP(&cfg.file, "output", "o", "", "write result to FILE instead of standard output")

	flag.Parse()

	in := input(cfg, flag.Args())
	out := output(cfg)

	Shuffle(in, out)
}

func output(c config) io.Writer {
	var out io.Writer = os.Stdout
	var err error
	if c.file != "" {
		if out, err = os.Create(c.file); err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
	}
	return out
}

func input(c config, args []string) []string {
	if c.echo {
		return args
	}

	if c.inputRange.isSet {
		return c.inputRange.Sequence()
	}

	if len(args) == 0 || args[0] == "-" {
		return readAll(os.Stdin)
	}

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	return readAll(file)
}

func readAll(r io.Reader) []string {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
