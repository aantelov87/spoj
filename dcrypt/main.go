package main

import (
	"bufio"
	"fmt"
	"os"
)

const max = 100000

var outWriter = bufio.NewWriterSize(os.Stdout, max)

func decrypt(m []byte, k int8) {
	l := len(m)
	var d int8
	var rc int8 = 0
	var c int8
	if k >= 26 && k <= 51 {
		rc = 1
	}
	for i := 0; i < l; i++ {
		if m[i] == '.' {
			m[i] = 0x20
			continue
		}
		d = 0x41
		c = 0x20
		if m[i] > 0x60 {
			d = 0x61
			c = -0x20
		}
		v := (int8(m[i])-d+k)%26 + d
		if rc == 1 {
			v += c
		}
		m[i] = byte(v)
	}

	outWriter.Write(m)
	outWriter.WriteByte('\n')
	outWriter.Flush()
}

func main() {
	var t int
	var k int8
	var s string
	fmt.Scanf("%d", &t)
	for t > 0 {
		fmt.Scanf("%d", &k)
		fmt.Scanf("%s", &s)
		b := []byte(s)
		decrypt(b, k)
		t--
	}
}
