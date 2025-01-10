package code_test

import (
	"testing"

	"gitlab.com/linkinlog/compiler/code"
)

func TestMake(t *testing.T) {
	testcases := map[string]struct {
		op       code.Opcode
		operands []int
		expected []byte
	}{
		"penultimate constant": {
			code.OpConstant,
			[]int{65534},
			[]byte{byte(code.OpConstant), 255, 254},
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			instructions := code.Make(tc.op, tc.operands...)

			if len(instructions) != len(tc.expected) {
				t.Fatalf("expected len(%d), got len(%d)", len(tc.expected), len(instructions))
			}

			for idx, instruction := range instructions {
				if instruction != tc.expected[idx] {
					t.Fatalf("wrong byte at pos %d, got: %v, wanted: %v", idx, instruction, tc.expected[idx])
				}
			}
		})
	}
}
