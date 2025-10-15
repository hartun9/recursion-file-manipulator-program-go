package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

var binaryPath string

// TestMain builds the binary before running tests
func TestMain(m *testing.M) {
	// Build the binary
	binaryPath = filepath.Join(os.TempDir(), "file-manipulator-test")
	build := exec.Command("go", "build", "-o", binaryPath, ".")
	if err := build.Run(); err != nil {
		os.Exit(1)
	}

	// Run tests
	code := m.Run()

	// Cleanup
	os.Remove(binaryPath)
	os.Exit(code)
}

// TestReverseCommand tests the reverse command
func TestReverseCommand(t *testing.T) {
	// Create temporary input file
	inputFile, err := os.CreateTemp("", "input-*.txt")
	if err != nil {
		t.Fatalf("Failed to create input file: %v", err)
	}
	defer os.Remove(inputFile.Name())

	inputContent := "line1\nline2\nline3\n"
	if _, err := inputFile.WriteString(inputContent); err != nil {
		t.Fatalf("Failed to write to input file: %v", err)
	}
	inputFile.Close()

	// Create temporary output file
	outputFile, err := os.CreateTemp("", "output-*.txt")
	if err != nil {
		t.Fatalf("Failed to create output file: %v", err)
	}
	outputPath := outputFile.Name()
	outputFile.Close()
	defer os.Remove(outputPath)

	// Run the command
	cmd := exec.Command(binaryPath, "reverse", inputFile.Name(), outputPath)
	if err := cmd.Run(); err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	// Verify output
	result, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	expected := "line3\nline2\nline1\n"
	if string(result) != expected {
		t.Errorf("Expected %q, got %q", expected, string(result))
	}
}

// TestCopyCommand tests the copy command
func TestCopyCommand(t *testing.T) {
	// Create temporary input file
	inputFile, err := os.CreateTemp("", "input-*.txt")
	if err != nil {
		t.Fatalf("Failed to create input file: %v", err)
	}
	defer os.Remove(inputFile.Name())

	inputContent := "Hello, World!\n"
	if _, err := inputFile.WriteString(inputContent); err != nil {
		t.Fatalf("Failed to write to input file: %v", err)
	}
	inputFile.Close()

	// Create temporary output file
	outputFile, err := os.CreateTemp("", "output-*.txt")
	if err != nil {
		t.Fatalf("Failed to create output file: %v", err)
	}
	outputPath := outputFile.Name()
	outputFile.Close()
	defer os.Remove(outputPath)

	// Run the command
	cmd := exec.Command(binaryPath, "copy", inputFile.Name(), outputPath)
	if err := cmd.Run(); err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	// Verify output
	result, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	if string(result) != inputContent {
		t.Errorf("Expected %q, got %q", inputContent, string(result))
	}
}

// TestDuplicateContentsCommand tests the duplicate-contents command
func TestDuplicateContentsCommand(t *testing.T) {
	// Create temporary input file
	inputFile, err := os.CreateTemp("", "input-*.txt")
	if err != nil {
		t.Fatalf("Failed to create input file: %v", err)
	}
	defer os.Remove(inputFile.Name())

	inputContent := "Test\n"
	if _, err := inputFile.WriteString(inputContent); err != nil {
		t.Fatalf("Failed to write to input file: %v", err)
	}
	inputFile.Close()

	// Run the command
	cmd := exec.Command(binaryPath, "duplicate-contents", inputFile.Name(), "2")
	if err := cmd.Run(); err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	// Verify output
	result, err := os.ReadFile(inputFile.Name())
	if err != nil {
		t.Fatalf("Failed to read input file: %v", err)
	}

	expected := "Test\nTest\nTest\n"
	if string(result) != expected {
		t.Errorf("Expected %q, got %q", expected, string(result))
	}
}

// TestReplaceStringCommand tests the replace-string command
func TestReplaceStringCommand(t *testing.T) {
	// Create temporary input file
	inputFile, err := os.CreateTemp("", "input-*.txt")
	if err != nil {
		t.Fatalf("Failed to create input file: %v", err)
	}
	defer os.Remove(inputFile.Name())

	inputContent := "Hello World"
	if _, err := inputFile.WriteString(inputContent); err != nil {
		t.Fatalf("Failed to write to input file: %v", err)
	}
	inputFile.Close()

	// Run the command
	cmd := exec.Command(binaryPath, "replace-string", inputFile.Name(), "World", "Go")
	if err := cmd.Run(); err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	// Verify output
	result, err := os.ReadFile(inputFile.Name())
	if err != nil {
		t.Fatalf("Failed to read input file: %v", err)
	}

	expected := "Hello Go"
	if string(result) != expected {
		t.Errorf("Expected %q, got %q", expected, string(result))
	}
}

// TestInvalidCommand tests that an invalid command returns an error
func TestInvalidCommand(t *testing.T) {
	cmd := exec.Command(binaryPath, "invalid-command")
	err := cmd.Run()
	if err == nil {
		t.Error("Expected command to fail with invalid command, but it succeeded")
	}
}

// TestMissingArguments tests that missing arguments return an error
func TestMissingArguments(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "No arguments",
			args: []string{},
		},
		{
			name: "Copy with missing output",
			args: []string{"copy", "input.txt"},
		},
		{
			name: "Reverse with missing output",
			args: []string{"reverse", "input.txt"},
		},
		{
			name: "Duplicate-contents with missing n",
			args: []string{"duplicate-contents", "input.txt"},
		},
		{
			name: "Replace-string with missing newstring",
			args: []string{"replace-string", "input.txt", "needle"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(binaryPath, tt.args...)
			err := cmd.Run()
			if err == nil {
				t.Errorf("Expected command to fail with missing arguments, but it succeeded")
			}
		})
	}
}
