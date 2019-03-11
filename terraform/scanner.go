package terraform

import (
	"fmt"
	gohcl2 "github.com/hashicorp/hcl2/gohcl"
	hcl2parse "github.com/hashicorp/hcl2/hclparse"
	"reflect"
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

	populateConfigWithFileRefs(&config)

	return &config // todo: return decoding errors
}

func populateConfigWithFileRefs(c *Config) {
	fmt.Println("running func")

	root := reflect.Indirect(reflect.ValueOf(c))

	// For each field in Config
	for i := 0; i < root.NumField(); i++ {
		fmt.Printf("%d: ", i)
		field := root.Type().Field(i).Type
		fmt.Println(field.Elem().Name())

		for j := 0; j < field.Elem().NumField(); j++ {
			fmt.Printf("  %d: ", j)
			innerfield := field.Elem().Field(j)
			fmt.Println(innerfield.Name)
		}
	}

	// reflect.ValueOf(&n).Elem().FieldByName("N").SetInt(7)
	fmt.Printf("\n\n\n\n")
}
