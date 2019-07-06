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

// TODO: what are we measuring here?
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

func TestMinNodes_MoveTopLeft(t *testing.T) {
	testCases := []struct{
		mn 		*MinNodes // initial MinNode
		move 	uint // number of move
	}{
		{
			NewMinNodes(0x4E40), // top-left corner
			0,
		},{
			NewMinNodes(0x2720), // top-right corner
			1,
		},{
			NewMinNodes(0x04E4), // bottom-left corner
			2,
		},{
			NewMinNodes(0x0272), // bottom-right corner
			3,
		},
	}
	crossTopLeft := MinNodes(0x4E40)

	for _, tc := range testCases {
		result := tc.mn.MoveTopLeft(tc.move)
		if result != crossTopLeft {
			t.Errorf("fait to do %d move. 0x%X != 0x%X", tc.move, result, crossTopLeft)
		}
	}
}