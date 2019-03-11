package terraform_test

import (
	"terrain/terraform"
	"testing"
)

func TestScanFileMulti(t *testing.T) {
	expectedResourceCount := 2
	config := terraform.ScanFile("testdata/multi.tf")

	if len(config.Resources) != expectedResourceCount {
		t.Fatalf("incorrect count of resources in the test file")
	}
}

func TestScanFileEmpty(t *testing.T) {
	expectedResourceCount := 0
	config := terraform.ScanFile("testdata/empty.tf")

	if len(config.Resources) != expectedResourceCount {
		t.Fatalf("test file should have no resources")
	}
}
