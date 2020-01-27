package steptype

import (
	"fmt"

	"github.com/gobuffalo/packr"
	"gopkg.in/yaml.v2"
)

var (
	ArgumentsDefault    string   = ""
	ReturnsDefault      string   = ""
	SpecTempalteDefault string   = ""
	MetadataDeault      Metadata = Metadata{}
)

func init() {
	box := packr.NewBox("./defaults")
	{
		args, err := box.FindString("arguments.json")
		if err != nil {
			panic(fmt.Errorf("Failed to find file: arguments.json with error: %s", err.Error()))
		}
		ArgumentsDefault = args
	}

	{
		args, err := box.FindString("returns.json")
		if err != nil {
			panic(fmt.Errorf("Failed to find file: returns.json with error: %s", err.Error()))
		}
		ReturnsDefault = args

	}

	{
		args, err := box.FindString("spectemplate.yaml.tmpl")
		if err != nil {
			panic(fmt.Errorf("Failed to find file: spectemplate.yaml.tmpl with error: %s", err.Error()))
		}
		SpecTempalteDefault = args
	}

	{
		args, err := box.FindString("metadata.yaml")
		if err != nil {
			panic(fmt.Errorf("Failed to find file: metadata.yaml with error: %s", err.Error()))
		}
		err = yaml.Unmarshal([]byte(args), &MetadataDeault)
		if err != nil {
			panic(fmt.Errorf("Failed create defailt metadata with error: %s", err.Error()))
		}
	}
}
