package terraform

import (
	gohcl2 "github.com/hashicorp/hcl2/gohcl"
	hcl2parse "github.com/hashicorp/hcl2/hclparse"
)

func ScanFile(filePath string) *Config {
	// Populate an instance of Config
	var config Config

	parser := hcl2parse.NewParser()
	f, _ := parser.ParseHCLFile(filePath) // we ignore the diagnostic errors here, because they don't allow for empty TF files

	diags := gohcl2.DecodeBody(f.Body, nil, &config)
	if diags.HasErrors() {
		panic("diags has errors on decoding body")
	}

	return &config // todo: return decoding errors
}
