package reader

import (
	"bufio"
	"bytes"
	"io"
)

// ScanLinesByDotNewLine Modified and copied from bufio.ScanLines
// Split by '.' and '\n' symbols
func ScanLinesByDotNewLine(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	newLine := bytes.IndexByte(data, '\n')
	dot := bytes.IndexByte(data, '.')
	var end int

	if dot <= newLine {
		end = dot
	} else {
		end = newLine
	}

	if end >= 0 {
		return end + 1, dropCR(data[0:end]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

// dropCR copied from bufio.dropCR
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

func CreateBufferedScanner(reader io.Reader, bufferMaxLen int, splitFunc bufio.SplitFunc) *bufio.Scanner {
	scanner := bufio.NewScanner(reader)
	scanner.Split(splitFunc)
	buffer := make([]byte, 0, bufferMaxLen)
	scanner.Buffer(buffer, bufferMaxLen)
	return scanner
}
