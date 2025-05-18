package lab2

import (
	"bytes"
	"strings"
	"testing"
)

func TestComputeHandler_Success(t *testing.T) {
	input := strings.NewReader("+ 3 5")
	output := &bytes.Buffer{}

	handler := &ComputeHandler{Input: input, Output: output}
	err := handler.Compute()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := "8"
	if output.String() != expected {
		t.Errorf("Expected %s, got %s", expected, output.String())
	}
}

func TestComputeHandler_Error(t *testing.T) {
	input := strings.NewReader("invalid expression")
	output := &bytes.Buffer{}

	handler := &ComputeHandler{Input: input, Output: output}
	err := handler.Compute()

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}
