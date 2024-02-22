package worker

import (
	"testing"
)

//func TestName(t *testing.T) {
//	dispatch := NewDispatcher(MaxWorker)
//	dispatch.Run()
//	for i := 0; i < 10000; i++ {
//		payloadHandler(i)
//	}
//	fmt.Println(111111111111111111)
//	time.Sleep(5 * time.Second)
//}

func BenchmarkPayloadHandler(b *testing.B) {
	dispatch := NewDispatcher(MaxWorker)
	dispatch.Run()
	for i := 0; i < b.N; i++ {
		payloadHandler(i)
	}
}
