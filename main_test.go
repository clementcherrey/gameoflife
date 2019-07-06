package bitmap

import (
	"fmt"
	"testing"
)

func TestSetBit(t *testing.T) {
	mn := new(MinNodes)
	fmt.Println(*mn)

	mn.SetBit(2)
	if  *mn != 4 {
		t.Error()
	}
}

func BenchmarkMinNodes_SetBit(b *testing.B) {
	mn := NewMinNodes(253)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mn.SetBit(2)
		mn.SetBit(3)
		mn.UnSetBit(3)
		mn.UnSetBit(4)
		mn.SetBit(2)
		mn.SetBit(3)
		mn.UnSetBit(3)
	}
}

func TestDummy(t *testing.T) {
	m := 18 << 1
	fmt.Println(m)
}