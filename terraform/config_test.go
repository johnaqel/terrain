package terraform_test

import (
	"terrain/terraform"
	"testing"
)

func TestFindTfFiles(t *testing.T) {
	expectedFiles := 2
	glob := "testdata/*.tf"

	files := terraform.FindTfFiles(glob)

	if len(files) != expectedFiles {
		t.Errorf("incorrect number of test Terraform files")
	}
}

func TestBuildConfig(t *testing.T) {
	files := []string{
		"testdata/empty.tf",
		"testdata/single.tf",
		"testdata/multi.tf",
	}

	expectedResourceType := "aws_key_pair"
	expectedResourceCount := 3

	config, err := terraform.BuildConfig(files)

	if err != nil {
		t.Fatalf("errors while building the full config: %s", err)
	}

	if len(config.Resources) != expectedResourceCount {
		t.Fatalf("did not find all expected terraform resources")
	}

	var unfound = true
	for _, res := range config.Resources {
		if res.Type == expectedResourceType {
			unfound = false
		}
	}

	if unfound {
		t.Fatalf("unable to find the expected terraform resource")
	}
}
