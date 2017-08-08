package main

import (
	"bufio"
	"io"
	"math/rand"
	"time"
)

// Shuffle takes a slice of strings and writes them
// out in random order using fisher-yates shuffle
func Shuffle(in []string, out io.Writer) {
	w := bufio.NewWriter(out)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	m := len(in)
	for m > 0 {
		i := r.Intn(m)
		w.WriteString(in[i] + "\n")
		in[i] = in[m-1]
		m--
	}
	w.Flush()
}
