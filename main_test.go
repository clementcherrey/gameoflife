package bitmap

import (
	"testing"
)

func TestSetBit(t *testing.T) {
	mn := new(MinNodes)
	mn.SetBit(2)
	if *mn != 4 {
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
	testCases := []struct {
		mn   *MinNodes // initial MinNode
		move uint      // number of move
	}{
		{
			NewMinNodes(0x4E40), // top-left corner
			0,
		}, {
			NewMinNodes(0x2720), // top-right corner
			1,
		}, {
			NewMinNodes(0x04E4), // bottom-left corner
			2,
		}, {
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

func TestMinNodes_ApplyGoFRuleToTopLeft(t *testing.T) {
	testCases := []struct {
		mn       *MinNodes // initial MinNode
		expected bool
	}{
		{
			NewMinNodes(0x4E40),
			false,
		}, {
			NewMinNodes(0x4E4F),
			false,
		}, {
			NewMinNodes(0x4040),
			false,
		}, {
			NewMinNodes(0x0000),
			false,
		}, {
			NewMinNodes(0x4240),
			true,
		}, {
			NewMinNodes(0x0E40),
			true,
		}, {
			NewMinNodes(0x0A40),
			true,
		}, {
			NewMinNodes(0x0340),
			false,
		},
	}
	for _, tc := range testCases {
		result := tc.mn.ApplyGoFRuleToTopLeft()
		if result != tc.expected {
			t.Errorf("ApplyGoFRuleToTopLeft to 0x%X fail. "+
				"Got %v. Expected %v", *tc.mn, result, tc.expected)
		}
	}
}

// TODO: finish
func TestMinNodes_ApplyGoFRule(t *testing.T) {
	testCases := []struct {
		mn       *MinNodes
		expected byte
	}{
		{
			NewMinNodes(0),
			0,
		}, {
			NewMinNodes(0x0E40),
			0x08,
		},{
			NewMinNodes(0x0E20),
			0x0C,
		},
	}

	for _, tc := range testCases {
		result := tc.mn.ApplyGoFRule()
		if result != tc.expected {
			t.Errorf("ApplyGoFRule to 0x%X fail. "+
				"Got 0x%X. Expected 0x%X", *tc.mn, result, tc.expected)
		}
	}
}
