package string_concat

import "testing"

func BenchmarkJoinAdd(b *testing.B) {
	str1 := "Hello "
	str2 := "World"
	str3 := ""
	_ = str3
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str3 = joinAdd(str1, str2)
	}
}

func BenchmarkJoinBytes(b *testing.B) {
	str1 := "Hello "
	str2 := "World"
	str3 := ""
	_ = str3
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str3 = joinBytes(str1, str2)
	}
}

func BenchmarkJoinFmt(b *testing.B) {
	str1 := "Hello "
	str2 := "World"
	str3 := ""
	_ = str3
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str3 = joinFmt(str1, str2)
	}
}

func BenchmarkJoinStringJoin(b *testing.B) {
	str1 := "Hello "
	str2 := "World"
	str3 := ""
	_ = str3
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str3 = joinStringsJoin(str1, str2)
	}
}
