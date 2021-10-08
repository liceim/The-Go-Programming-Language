package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

const input = `SplitFunc is the signature of the split function used to tokenize the input. The arguments are an initial substring of the remaining unprocessed data and a flag, atEOF, that reports whether the Reader has no more data to give. The return values are the number of bytes to advance the input and the next token to return to the user, if any, plus an error, if any.
Scanning stops if the function returns an error, in which case some of the input may be discarded.
Otherwise, the Scanner advances the input. If the token is not nil, the Scanner returns it to the user. If the token is nil, the Scanner reads more data and continues scanning; if there is no more data--if atEOF was true--the Scanner returns. If the data does not yet hold a complete token, for instance if it has no newline while scanning lines, a SplitFunc can return (0, nil, nil) to signal the Scanner to read more data into the slice and try again with a longer slice starting at the same point in the input.
The function is never called with an empty data slice unless atEOF is true. If atEOF is true, however, data may be non-empty and, as always, holds unprocessed text.`

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
	}

	return len(p), scanner.Err()
}

func (c *WordCounter) String() string {
	return fmt.Sprintf("%d words", *c)
}

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))

	for scanner.Scan() {
		*l++
	}

	return len(p), scanner.Err()
}

func (l *LineCounter) String() string {
	return fmt.Sprintf("%d lines", *l)
}

func main() {
	var w WordCounter
	var l LineCounter

	w.Write([]byte(input))
	fmt.Println(w.String())

	l.Write([]byte(input))
	fmt.Println(l.String())
}
