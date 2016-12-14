package string_concat

import (
	"bytes"
	"fmt"
	"strings"
)

func joinAdd(a, b string) string {
	return a + b
}

func joinBytes(a, b string) string {
	bf := bytes.NewBuffer([]byte{})
	bf.WriteString(a)
	bf.WriteString(b)
	return bf.String()
}

func joinFmt(a, b string) string {
	return fmt.Sprintf("%s%s", a, b)
}

func joinStringsJoin(a, b string) string {
	return strings.Join([]string{a, b}, "")
}

func selfConcatAdd(s string, n int) string {
	ret := ""
	for i := 0; i < n; i++ {
		ret += s
	}
	return ret
}

func selfConcatBytes(s string, n int) string {
	bf := bytes.NewBuffer([]byte{})
	for i := 0; i < n; i++ {
		bf.WriteString(s)
	}
	return bf.String()
}

func selfConcatJoin(s string, n int) string {
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		strs[i] = s
	}
	return strings.Join(strs, "")
}
