/*
Exercise 7.5: The LimitReader function in the io package accepts an io.Reader r and a number of bytesn, 
and returns another Reader that reads from r but reports an end-of-file condition after n bytes. Implement it.
*/

type LimitedReader struct {
	r io.Reader
	n int64
}

func NewLimitedReader(r io.Reader, n int64) *LimitedReader {
	return &LimitedReader{r, n}
}

func (r *LimitedReader) Read(p []byte) (n int, err error) {
	if r.n <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > r.n {
		p = p[:r.n]
	}

	n, err = r.r.Read(p)
	r.n -= int64(n)
	return
}
