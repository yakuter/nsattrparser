// Package attrparser provides functionality to parse binary encoded NSAttributedString data.
package nsattrparser

import (
	"bytes"
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
	left, right := 0, len(stream)-1
	for left < right {
		if !bytes.Equal(stream[left:left+2], startPattern) {
			left++
		} else if !bytes.Equal(stream[right-1:right+1], endPattern) {
			right--
		} else {
			stream = stream[left+2 : right-1] // exclude patterns
			if utf8.Valid(stream) {
				return string(stream[1:]), nil
			} else {
				return string(stream[3:]), nil
			}
		}
	}
	return "", errors.New("start or end pattern not found")
}
