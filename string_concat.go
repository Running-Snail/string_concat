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
