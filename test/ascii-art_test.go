package main
import (
	"ascii-art-output/internal" // Import the internal package
	"io"
	"os"
	"strings"
	"testing"
)
func TestPrintAllStringASCII(t *testing.T) {
	testCases := []struct {
		inputFile          string
		expectedOutputFile string
	}{
		{"testdata/example1.txt", "testdata/expected_example1.txt"},
		{"testdata/example2.txt", "testdata/expected_example2.txt"},
		{"testdata/example3.txt", "testdata/expected_example3.txt"},
	}
	for _, tc := range testCases {
		// Load input and expected output files
		input, _ := os.ReadFile(tc.inputFile)
		expectedOutput, _ := os.ReadFile(tc.expectedOutputFile)
		// Use the internal package functions
		fileLines := internal.ReadBannerFile("/Users/mnozimjo/ascii-art/banners/standard.txt") // Specify the correct file path
		asciiTemplates := internal.ParseBanner(fileLines)
		// Call the function
		output := capturePrint(func() {
			internal.RenderASCIIArt(string(input), asciiTemplates) // Use internal.RenderASCIIArt
		})
		// Compare the actual and expected output
		if strings.TrimSpace(output) != strings.TrimSpace(string(expectedOutput)) {
			t.Errorf("Expected output does not match actual output for input %s\nExpected:\n%s\nActual:\n%s",
				tc.inputFile, strings.TrimSpace(string(expectedOutput)), strings.TrimSpace(output))
		}
	}
}
// Utility function to capture the printed output
func capturePrint(fn func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	fn()
	w.Close()
	os.Stdout = old
	out, _ := io.ReadAll(r)
	return string(out)
}

