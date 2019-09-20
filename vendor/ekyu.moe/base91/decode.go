package base91

import (
	"io"
)

func decode(src []byte) []byte {
	// b and n must be at least uint32
	var b, n uint32 = 0, 0
	v := -1
	decoded := []byte{}

	for i := 0; i < len(src); i++ {
		p := dectab[src[i]]
		if p > 90 {
			// skip invalid character silently
			continue
		}
		if v < 0 {
			v = int(p)
			continue
		}
		v += int(p) * 91
		b |= uint32(v) << n
		if v&0x1fff > 88 {
			n += 13
		} else {
			n += 14
		}
		for {
			decoded = append(decoded, uint8(b))
			b >>= 8
			n -= 8
			if n <= 7 {
				break
			}
		}
		v = -1
	}

	if v > -1 {
		decoded = append(decoded, uint8(b|uint32(v)<<n))
	}

	return decoded
}

// Decode decodes src into dst. It returns the number of bytes written to dst.
// Whereas base91 uses up to nearly all printable characters, Decode skips
// invalid characters silently.
func Decode(dst, src []byte) int {
	return copy(dst, decode(src))
}

// DecodeString returns the bytes represented by the base91 string s, probably
// what you want.
func DecodeString(s string) []byte {
	return decode([]byte(s))
}

type decoder struct {
	// input
	reader io.Reader
	err    error

	outBuf []byte
	inBuf  *[1024]byte
	nInBuf int

	b, n uint32
	v    int
}

// NewDecoder constructs a new base91 stream decoder. Data read from the
// returned reader is base91 decoded from r.
func NewDecoder(r io.Reader) io.Reader {
	return &decoder{
		reader: r,
		err:    nil,

		outBuf: []byte{},
		inBuf:  new([1024]byte),
		nInBuf: 0,

		b: 0,
		n: 0,
		v: -1,
	}
}

func (d *decoder) Read(c []byte) (int, error) {
	if len(c) == 0 {
		return 0, nil
	}
	if d.err != nil {
		return 0, d.err
	}

	n := 0
	for n < len(c) && d.err == nil {
		var upn int
		upn, d.err = d.reader.Read(d.inBuf[d.nInBuf:])

		next := d.nInBuf + upn
		for ; d.nInBuf < next; d.nInBuf++ {
			p := dectab[d.inBuf[d.nInBuf]]
			if p > 90 {
				continue
			}
			if d.v < 0 {
				d.v = int(p)
				continue
			}
			d.v += int(p) * 91
			d.b |= uint32(d.v) << d.n
			if d.v&0x1fff > 88 {
				d.n += 13
			} else {
				d.n += 14
			}
			for {
				d.outBuf = append(d.outBuf, uint8(d.b))
				d.b >>= 8
				d.n -= 8
				if d.n <= 7 {
					break
				}
			}
			d.v = -1
		}

		if d.nInBuf >= 1024 {
			d.nInBuf = 0
		}

		if d.err != nil && d.v > -1 {
			// EOF is met, flush
			d.outBuf = append(d.outBuf, uint8(d.b|uint32(d.v)<<d.n))
		}

		m := copy(c[n:], d.outBuf)
		d.outBuf = d.outBuf[m:]
		n += m
		// if m < len(c), then outBuf did not suffuse c
	}

	return n, d.err
}
