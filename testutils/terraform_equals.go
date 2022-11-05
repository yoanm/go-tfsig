package testutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/andreyvit/diff"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func LoadGoldenFile(filename string) (*string, error) {
	fp := filepath.Join("testdata", filename+".golden.tf")
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		return nil, fmt.Errorf("golden file '%s' doesn't exist", fp)
	}
	content, err := ioutil.ReadFile(fp)
	if err != nil {
		return nil, fmt.Errorf("failed to load golden file %s: %v", fp, err)
	}

	c := string(content)

	return &c, nil
}

func EnsureFileContentEquals(file *hclwrite.File, expected string) error {
	actual := string(file.Bytes())
	if actual != expected {
		return fmt.Errorf("\n- expected\n+ actual\n\n%v", diff.LineDiff(expected, actual))
	}

	return nil
}

func EnsureBlockFileEqualsGoldenFile(block *hclwrite.Block, goldenFile string) error {
	hclFile := hclwrite.NewEmptyFile()

	if block != nil {
		hclFile.Body().AppendBlock(block)
	}

	return EnsureFileEqualsGoldenFile(hclFile, goldenFile)
}

func EnsureFileEqualsGoldenFile(f *hclwrite.File, goldenFile string) error {
	expected, err := LoadGoldenFile(goldenFile)
	if err != nil {
		return err
	}

	return EnsureFileContentEquals(f, *expected)
}
