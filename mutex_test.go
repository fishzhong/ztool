package ztool

import "testing"

func TestGet(t *testing.T) {
	db1 := GetDb()
	db1.Get()
}

//func BenchmarkGet(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		db1 := GetDb()
//		db1.Get()
//	}
//}
