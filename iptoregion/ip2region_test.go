package iptoregion

import "testing"

func TestIp(t *testing.T) {
	ip("1.2.3.4")
}

func BenchmarkIp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ip("8.217.159.0")
	}
}
