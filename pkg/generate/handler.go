package generate

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
)

type (
	// Handler - exposed struct that implementd Handler interface
	Handler struct{}
)

// Handle - the function that will be called from the CLI with viper config
// to provide access to all flags
func (g *Handler) Handle(cnf *viper.Viper) error {
	res := StepType{
		Version: "1.0",
		Kind:    "step-type",
		Metadata: Metadata{
			Version:  cnf.GetString("version"),
			Official: cnf.GetBool("official"),
			IsPublic: cnf.GetBool("isPublic"),
			Tags:     cnf.GetStringSlice("tags"),
		},
		Spec: Spec{
			Delimiters: Delimiters{
				Left:  cnf.GetString("left-delimiter"),
				Right: cnf.GetString("right-delimiter"),
			},
		},
	}

	acc := cnf.GetString("account")
	name := cnf.GetStringSlice("name")[0]
	if acc != "" {
		res.Metadata.Name = fmt.Sprintf("%s/%s", acc, name)
	} else {
		res.Metadata.Name = name
	}
	specFilePath := cnf.GetString("spec-file")
	if specFilePath != "" {
		specFileBytes, err := ioutil.ReadFile(specFilePath)
		res.Spec.StepsTemplate = string(specFileBytes)
		if err != nil {
			return err
		}
	}

	argumentsJSONFilePath := cnf.GetString("arguments-json-file")
	if specFilePath != "" {
		argumentsJSONBytes, err := ioutil.ReadFile(argumentsJSONFilePath)
		res.Spec.Arguments = string(argumentsJSONBytes)
		if err != nil {
			return err
		}
	}

	returnsJSONFilePath := cnf.GetString("returns-json-file")
	if specFilePath != "" {
		returnsJSONBytes, err := ioutil.ReadFile(returnsJSONFilePath)
		res.Spec.Returns = string(returnsJSONBytes)
		if err != nil {
			return err
		}
	}

	descriptionFilePath := cnf.GetString("description-file")
	if specFilePath != "" {
		descriptionBytes, err := ioutil.ReadFile(descriptionFilePath)
		res.Metadata.Description = string(descriptionBytes)
		if err != nil {
			return err
		}
	}
	maintainerName := cnf.GetString("maintainerName")
	maintainerEmail := cnf.GetString("maintainerEmail")
	if maintainerName != "" && maintainerEmail != "" {
		res.Metadata.Maintainers = []Maintainer{
			Maintainer{
				Name:  maintainerName,
				Email: maintainerEmail,
			},
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