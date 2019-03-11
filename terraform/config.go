package terraform

import (
	"github.com/hashicorp/hcl2/hcl"
	"github.com/imdario/mergo"
	"io"
	"path/filepath"
)

type Config struct {
	Resources []Resource `json:"resources" hcl:"resource,block"`
}

type Resource struct {
	SourceFile string   `json:"source"`
	Type       string   `json:"type" hcl:"type,label"`
	Name       string   `json:"name" hcl:"name,label"`
	Config     hcl.Body `hcl:",remain"` // This must be here, else the HCL parser fails
}

type ConfigScanner interface {
	ListScannedFiles() []string
	Run(io.Writer) error
}

func Run(writer interface{}) error {
	files := FindTfFiles("*.tf")

	_, err := BuildConfig(files)
	if err != nil {
		return err
	}

	// writer.Write(config)

	return nil
}

func BuildConfig(files []string) (*Config, error) {
	full := &Config{}
	partial := make(chan *Config, len(files))

	for _, file := range files {
		go BuildPartialConfig(file, partial)
	}

	// Todo: Make this a loop over a range. Need to fix the deadlock issue
	for i := 0; i < len(files); i++ {
		if err := mergo.Merge(full, <-partial, mergo.WithAppendSlice); err != nil {
			return nil, err
		}
	}

	return full, nil
}

func BuildPartialConfig(file string, partial chan *Config) {
	partial <- ScanFile(file)
}

func FindTfFiles(glob string) []string {
	files, _ := filepath.Glob(glob)
	return files
}
