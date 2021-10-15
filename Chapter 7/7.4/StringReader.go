/*
Exercise 7.4: The strings.NewReader function returns a value that satisfies the io.Reader interface (and others) by reading from its argument, a string. 
Implement a simple version of NewReader yourself, and use it to make the HTML parser (§5.2) take input from a string.
*/

type StringReader struct {
	s string
	i int64
}

func (r *StringReader) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}

	n = copy(p, r.s[r.i:])
	if r.i += int64(n); r.i >= int64(len(r.s)) {
		err = io.EOF
	}
	return
}

func NewReader(s string) *StringReader {
	return &StringReader{s, 0}
}
