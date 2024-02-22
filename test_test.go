package ztool

import "testing"

func BenchmarkForeach(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Foreach(accounts)
	}
}

func BenchmarkForeach2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Foreach2(account2)
	}
}
