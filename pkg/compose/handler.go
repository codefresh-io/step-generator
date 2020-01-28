package compose

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"

	"github.com/codefresh-io/step-generator/pkg/steptype"
)

type (
	// Handler - exposed struct that implementd Handler interface
	Handler struct{}
)

// Handle - the function that will be called from the CLI with viper config
// to provide access to all flags
func (g *Handler) Handle(cnf *viper.Viper) error {
	res := steptype.StepType{
		Version: "1.0",
		Kind:    "step-type",
		Metadata: steptype.Metadata{
			Version: cnf.GetString("version"),
		},
		Spec: steptype.Spec{
			Delimiters: steptype.Delimiters{
				Left:  cnf.GetString("leftDelimiter"),
				Right: cnf.GetString("rightDelimiter"),
			},
		},
	}
	specFilePath := cnf.GetString("specFile")
	argumentsJSONFilePath := cnf.GetString("argumentsJsonFile")
	returnsJSONFilePath := cnf.GetString("returnsJsonFile")
	metadataFilePath := cnf.GetString("metadataFile")

	directory := cnf.GetString("directory")
	if directory != "" {
		fmt.Printf("Directory given, using default values inside it")
		specFilePath = path.Join(directory, "spectemplate.yaml.tmpl")
		argumentsJSONFilePath = path.Join(directory, "arguments.json")
		returnsJSONFilePath = path.Join(directory, "returns.json")
		metadataFilePath = path.Join(directory, "metadata.yaml")
	}

	if specFilePath != "" {
		fmt.Printf("Reading spec file %s\n", specFilePath)
		specFileBytes, err := ioutil.ReadFile(specFilePath)
		res.Spec.StepsTemplate = string(specFileBytes)
		if err != nil {
			return err
		}
	}

	if specFilePath != "" {
		fmt.Printf("Reading arguments json file %s\n", specFilePath)
		argumentsJSONBytes, err := ioutil.ReadFile(argumentsJSONFilePath)
		res.Spec.Arguments = string(argumentsJSONBytes)
		if err != nil {
			return err
		}
	}

	if returnsJSONFilePath != "" {
		fmt.Printf("Reading returns json file %s\n", returnsJSONFilePath)
		returnsJSONBytes, err := ioutil.ReadFile(returnsJSONFilePath)
		res.Spec.Returns = string(returnsJSONBytes)
		if err != nil {
			return err
		}
	}

	if metadataFilePath != "" {
		fmt.Printf("Reading description file %s\n", metadataFilePath)
		metadataFileBytes, err := ioutil.ReadFile(metadataFilePath)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(metadataFileBytes, &res.Metadata)
		if err != nil {
			return err
		}
	}

	outFile := cnf.GetString("out")
	resBytes, err := res.Marshal()
	if err != nil {
		return err
	}
	if outFile == "" {
		fmt.Println(string(resBytes))
	} else {
		return ioutil.WriteFile(outFile, resBytes, os.ModePerm)
	}

	return nil
}
