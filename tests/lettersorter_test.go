package tests

import (
	"testing"

	"github.com/IndBrony/lettersorter"
)

func TestSortString(t *testing.T) {
	basicSortStringTesting(t, "osama", "aaoms")
	basicSortStringTesting(t, "omama", "aaomm")
	basicSortStringTesting(t, "test", "estt")
}
func basicSortStringTesting(t *testing.T, input, expectedOutput string) {
	output := lettersorter.SortString(input)
	if output != expectedOutput {
		t.Errorf("Test Fail, expecting %v but got %v", expectedOutput, output)
	}
}

func TestSortStringConcurrent(t *testing.T) {
	basicSortStringTestingConcurrent(t, "osama", "aaoms")
	basicSortStringTestingConcurrent(t, "omama", "aaomm")
	basicSortStringTestingConcurrent(t, "test", "estt")
}
func basicSortStringTestingConcurrent(t *testing.T, input, expectedOutput string) {
	output := lettersorter.SortStringConcurrent(input)
	if output != expectedOutput {
		t.Errorf("Test Fail, expecting %v but got %v", expectedOutput, output)
	}
}

func BenchmarkSortStringConcurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lettersorter.SortStringConcurrent("osam ghkyun hfdaa sd")
	}
}
func BenchmarkSortString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lettersorter.SortString("osam ghkyun hfdaa sd")
	}
}
