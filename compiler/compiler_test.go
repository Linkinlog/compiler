package compiler_test

import (
	"testing"

	"gitlab.com/linkinlog/compiler/ast"
	"gitlab.com/linkinlog/compiler/code"
	"gitlab.com/linkinlog/compiler/compiler"
	"gitlab.com/linkinlog/compiler/lexer"
	"gitlab.com/linkinlog/compiler/parser"
)

type compilerTestCase struct {
	input                string
	expectedInstructions []code.Instructions
	expectedConstants    []interface{}
}

func TestIntegerArithmetic(t *testing.T) {
	tests := []compilerTestCase{
		{
			input: "1 + 2",
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
			},
		},
	}

	runCompilerTests(t, tests)
}

func runCompilerTests(t *testing.T, testCases []compilerTestCase) {
	t.Helper()

	for _, tc := range testCases {
		program := parse(tc.input)

		compiler := compiler.New()

		err := compiler.Compile(program)
		if err != nil {
			t.Fatalf("compiler error: %s", err)
		}

		bytecode := compiler.Bytecode()

		err = testInstructions(tc.expectedInstructions, bytecode.Instructions)
		if err != nil {
			t.Fatalf("testInstructions failed: %s", err)
		}

		err = testConstants(tc.expectedConstants, bytecode.Constants)
		if err != nil {
			t.Fatalf("testConstants failed: %s", err)
		}
	}
}

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)

	return p.ParseProgram()
}
