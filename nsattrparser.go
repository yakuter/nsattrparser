// Package attrparser provides functionality to parse binary encoded NSAttributedString data.
package nsattrparser

import (
	"errors"
	"unicode/utf8"
)

var (
	startPattern = []byte{0x01, 0x2B} // Corresponds to `[<Start of Heading> (SOH), +]`
	endPattern   = []byte{0x86, 0x84} // Corresponds to `[<Start of Selected Area> (SSA), <Index> (IND)]`
)

// Parse attempts to extract and return the text contained within the binary data,
// delimited by specified start and end patterns.
func Parse(stream []byte) (string, error) {
	startIdx := findPatternIndex(stream, startPattern)
	if startIdx == -1 {
		return "", errors.New("start pattern not found")
	}
	stream = stream[startIdx+len(startPattern):]

	endIdx := findPatternIndex(stream, endPattern)
	if endIdx == -1 {
		return "", errors.New("end pattern not found")
	}
	stream = stream[:endIdx]

	if utf8.Valid(stream) {
		return string(stream[1:]), nil
	} else {
		return string(stream[3:]), nil
	}
}

func findPatternIndex(stream []byte, pattern []byte) int {
	for i := 0; i <= len(stream)-len(pattern); i++ {
		if match := stream[i : i+len(pattern)]; string(match) == string(pattern) {
			return i
		}
	}
	return -1
}
