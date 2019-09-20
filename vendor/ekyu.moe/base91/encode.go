package base91

import (
	"io"
)

func encode(src []byte) []byte {
	var b, n uint32 = 0, 0
	encoded := []byte{}

	for i := 0; i < len(src); i++ {
		b |= uint32(src[i]) << n
		n += 8

		if n <= 13 {
			continue
		}

		v := b & 0x1fff
		if v > 88 {
			b >>= 13
			n -= 13
		} else {
			v = b & 0x3fff
			b >>= 14
			n -= 14
		}

		encoded = append(encoded, enctab[v%91], enctab[v/91])
	}

	if n != 0 {
		encoded = append(encoded, enctab[b%91])
		if n > 7 || b > 90 {
			encoded = append(encoded, enctab[b/91])
		}
	}

	return encoded
}

// Encode encodes src into dst. It returns the number of bytes written to dst.
func Encode(dst, src []byte) int {
	return copy(dst, encode(src))
}

// EncodeToString returns the encoded base91 string of src, probably what you
// want.
func EncodeToString(src []byte) string {
	return string(encode(src))
}

type encoder struct {
	// output
	writer io.Writer
	buf    []byte

	b, n uint32
}

// NewEncoder returns a new base91 stream encoder. Data written to the returned
// writer will be encoded using base91 and then written to w. When finished
// writing, the caller must Close the returned encoder to flush any partially
// written blocks.
func NewEncoder(w io.Writer) io.WriteCloser {
	return &encoder{
		writer: w,
		b:      0,
		n:      0,
		buf:    []byte{},
	}
}

func (e *encoder) Write(c []byte) (int, error) {
	var err error
	var i int
	e.buf = e.buf[:0]

	for i = 0; i < len(c); i++ {
		e.b |= uint32(c[i]) << e.n
		e.n += 8

		if e.n <= 13 {
			continue
		}

		v := e.b & 0x1fff
		if v > 88 {
			e.b >>= 13
			e.n -= 13
		} else {
			v = e.b & 0x3fff
			e.b >>= 14
			e.n -= 14
		}

		e.buf = append(e.buf, enctab[v%91], enctab[v/91])
	}

	_, err = e.writer.Write(e.buf)
	return i, err
}

func (e *encoder) Close() error {
	var err error
	e.buf = e.buf[:0]

	if e.n != 0 {
		e.buf = append(e.buf, enctab[e.b%91])
		if e.n > 7 || e.b > 90 {
			e.buf = append(e.buf, enctab[e.b/91])
		}
	}

	_, err = e.writer.Write(e.buf)
	return err
}
