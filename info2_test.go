package info2parser

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		testname  string
		inputfile string
		expected  Info2
		err       error
	}{
		{
			testname:  "Valid Input 1",
			inputfile: "$IHIR5CD.msi",
			expected: Info2{
				Header:         2,
				FileSize:       69922816,
				DeletionTime:   133884183010530000,
				FileNameLength: uint32(48),
				OriginalPath:   "C:\\Users\\skibi\\Desktop\\go1.24.1.windows-386.msi",
			},
			err: nil,
		},
		{
			testname:  "Valid Input 2",
			inputfile: "$IJDBMHL.lnk",
			expected: Info2{
				Header:         2,
				FileSize:       1930,
				DeletionTime:   133884183068260000,
				FileNameLength: uint32(34),
				OriginalPath:   "C:\\Users\\skibi\\Desktop\\x64dbg.lnk",
			},
			err: nil,
		},
		{
			testname:  "Valid Input 3",
			inputfile: "$IR4K54E.txt",
			expected: Info2{
				Header:         2,
				FileSize:       75,
				DeletionTime:   133884183416750000,
				FileNameLength: uint32(51),
				OriginalPath:   "C:\\Users\\skibi\\Desktop\\drive-janitor\\delete_me.txt",
			},
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			sampledir := "SampleINFO2Files"
			absolutePath, err := filepath.Abs(sampledir + "/" + test.inputfile)
			if err != nil {
				t.Fatalf("Failed to get absolute path: %v", err)
			}
			data, err := os.ReadFile(absolutePath)
			if err != nil {
				t.Fatalf("Failed to read file: %v", err)
			}
			result, err := Parse(data)
			if err != nil {
				t.Fatalf("Failed to parse file: %v", err)
			}
			if result.Header != test.expected.Header {
				t.Errorf("Expected Header %d, got %d", test.expected.Header, result.Header)
			}
			if result.FileSize != test.expected.FileSize {
				t.Errorf("Expected FileSize %d, got %d", test.expected.FileSize, result.FileSize)
			}
			if result.DeletionTime != test.expected.DeletionTime {
				t.Errorf("Expected DeletionTime %d, got %d", test.expected.DeletionTime, result.DeletionTime)
			}
			if result.FileNameLength != test.expected.FileNameLength {
				t.Errorf("Expected FileNameLength %d, got %d", test.expected.FileNameLength, result.FileNameLength)
			}
			if result.OriginalPath != test.expected.OriginalPath {
				t.Errorf("Expected OriginalPath %s, got %s", test.expected.OriginalPath, result.OriginalPath)
			}
		})
	}
}

func TestParseInvalid(t *testing.T) {
	tests := []struct {
		testname  string
		inputfile string
		expected  Info2
		err       error
	}{
		{
			testname:  "Invalid Input",
			inputfile: "invalid_file.txt",
			expected: Info2{
				Header:         0,
				FileSize:       0,
				DeletionTime:   0,
				FileNameLength: 0,
				OriginalPath:   "",
			},
			err: fmt.Errorf("file invalid"),
		},
		{
			testname:  "empty file",
			inputfile: "empty_file.txt",
			expected: Info2{
				Header:         0,
				FileSize:       0,
				DeletionTime:   0,
				FileNameLength: 0,
				OriginalPath:   "",
			},
			err: fmt.Errorf("file invalid"),
		},
	}
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			sampledir := "SampleINFO2Files"
			absolutePath, err := filepath.Abs(sampledir + "/" + test.inputfile)
			if err != nil {
				t.Fatalf("Failed to get absolute path: %v", err)
			}
			data, err := os.ReadFile(absolutePath)
			if err != nil {
				t.Fatalf("Failed to read file: %v", err)
			}
			result, err := Parse(data)
			if err == nil {
				t.Fatalf("Expected error but got none")
			}
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}

}
