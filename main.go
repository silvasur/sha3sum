package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"golang.org/x/crypto/sha3"
	"hash"
	"io"
	"os"
)

var (
	algo = flag.String("algo", "256", "Which algo should be used (224,256,384,512)")
)

var algos = map[string]func() hash.Hash{
	"256": sha3.New256,
	"224": sha3.New224,
	"384": sha3.New384,
	"512": sha3.New512,
}

func main() {
	flag.Parse()

	files := append([]string{}, flag.Args()...)

	if len(files) == 0 {
		files = append(files, "-")
	}

	for _, fn := range files {
		if !doFile(fn) {
			os.Exit(1)
		}
	}
}

func doFile(fn string) bool {
	var r io.Reader

	if fn == "-" {
		r = os.Stdin
	} else {
		f, err := os.Open(fn)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return false
		}
		defer f.Close()

		r = f
	}

	gethash, ok := algos[*algo]
	if !ok {
		fmt.Fprintf(os.Stderr, "Unknown algo: %s\n", *algo)
		return false
	}

	hash := gethash()

	if _, err := io.Copy(hash, r); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return false
	}

	fmt.Printf("%s  %s\n", hex.EncodeToString(hash.Sum([]byte{})), fn)
	return true
}
