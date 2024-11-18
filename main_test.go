package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMainOutputSuccess(t *testing.T) {
	// Create a pipe to capture the output
	r, w, _ := os.Pipe()
	out := os.Stdout
	os.Stdout = w

	// Run the main function
	main()

	// Close the writer and restore the original stdout
	w.Close()
	os.Stdout = out

	// Read the captured output
	var buf bytes.Buffer
	buf.ReadFrom(r)

	expected := "Hello, World!\n"
	if buf.String() == expected {
		t.Logf("The results are as expected: %q", expected)
	}
	if buf.String() != expected {
		t.Errorf("expected %q but got %q", expected, buf.String())
	}
}

func TestMainOutputFailure(t *testing.T) {
	// Create a pipe to capture the output
	r, w, _ := os.Pipe()
	out := os.Stdout
	os.Stdout = w

	// Run the main function
	main()

	// Close the writer and restore the original stdout
	w.Close()
	os.Stdout = out

	// Read the captured output
	var buf bytes.Buffer
	buf.ReadFrom(r)

	// Intentionally incorrect expected output
	expected := "Goodbye, World!\n"
	if buf.String() == expected {
		t.Logf("The results are as expected: %q", expected)
	}
	if buf.String() != expected {
		t.Errorf("expected %q but got %q", expected, buf.String())
	}
}
