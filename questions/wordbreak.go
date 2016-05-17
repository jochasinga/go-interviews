/* http://thenoisychannel.com/2011/08/08/retiring-a-great-interview-problem
 * Given an input string and a dictionary of words,
 * segment the input string into a space-separated
 * sequence of dictionary words if possible. For
 * example, if the input string is "applepie" and
 * dictionary contains a standard set of English words,
 * then we would return the string "apple pie" as output.
 */

package main

import (
	"fmt"
	"strings"
)

func SegmentString(src string, ref []string) string {
	// BASE CASE: ref must not be empty
	prefix := ""
	suffix := ""
	if len(ref) > 0 {
		srcLen := len(src)
		for i, _ := range src {
			for _, word := range ref {
				prefix = src[0:i+1]
				if strings.Contains(word, prefix) {
					suffix = src[i+1:srcLen]
				}
				if strings.Contains(word, suffix) {
					return prefix + " " + suffix
				}
			}
		}
	}
	return ""
}

func main() {
	fmt.Println(SegmentString("applepie", []string{"apple", "pie", "foo"}))
}
