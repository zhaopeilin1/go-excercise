package main

import (
	"fmt"
	"io"
)

type errWriter struct {
	io.Writer
	err error
}

func (e *errWriter) Write(buf []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}
	var n int
	n, e.err = e.Writer.Write(buf)
	return n, nil
}
func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
	ew := &errWriter{Writer: w}
	fmt.Fprintf(ew, "HTTP/1.1 %d %srn", st.Code, st.Reason)
	for _, h := range headers {
		fmt.Fprintf(ew, "%s: %srn", h.Key, h.Value)
	}
	fmt.Fprint(ew, "rn")
	io.Copy(ew, body)
	return ew.err
}

func main() {

}
