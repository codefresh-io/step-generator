package generate

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/codefresh-io/step-generator/pkg/steptype"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type (
	// Handler - exposed struct that implementd Handler interface
	Handler struct{}
)

// Handle - the function that will be called from the CLI with viper config
// to provide access to all flags
func (g *Handler) Handle(cnf *viper.Viper) error {
	name := cnf.GetStringSlice("name")[0]
	dir := fmt.Sprintf("%s-%s", cnf.GetString("outDir"), name)

	if err := createDirIfNotExist(dir); err != nil {
		return err
	}

	res := steptype.StepType{
		Version:  "1.0",
		Kind:     "step-type",
		Metadata: steptype.MetadataDeault,
		Spec: steptype.Spec{
			Arguments: steptype.ArgumentsDefault,
			Delimiters: steptype.Delimiters{
				Left:  cnf.GetString("leftDelimiter"),
				Right: cnf.GetString("rightDelimiter"),
			},
			Returns:       steptype.ReturnsDefault,
			StepsTemplate: steptype.SpecTempalteDefault,
		},
	}

	{
		account := cnf.GetString("account")
		if account != "" {
			res.Metadata.Name = fmt.Sprintf("%s/%s", account, name)
		} else {
			res.Metadata.Name = name
		}

		for _, tag := range cnf.GetStringSlice("tags") {
			res.Metadata.Tags = append(res.Metadata.Tags, tag)
		}

		res.Metadata.Official = cnf.GetBool("official")
		maintainerName := cnf.GetString("maintainerName")
		maintainerEmail := cnf.GetString("maintainerEmail")
		if maintainerName != "" && maintainerEmail != "" {
			res.Metadata.Maintainers = []steptype.Maintainer{
				steptype.Maintainer{
					Name:  maintainerName,
					Email: maintainerEmail,
				},
			}
		}

		res.Metadata.IsPublic = cnf.GetBool("isPublic")
		version := cnf.GetString("version")
		if version != "" {
			res.Metadata.Version = version
		}

	}

	{
		if err := ioutil.WriteFile(filepath.Join(dir, "arguments.json"), []byte(res.Spec.Arguments), os.ModePerm); err != nil {
			fmt.Printf("Failed to write arguments.json file with error: %s", err.Error())
		}
	}

	{
		if err := ioutil.WriteFile(filepath.Join(dir, "spectemplate.yaml.tmpl"), []byte(res.Spec.StepsTemplate), os.ModePerm); err != nil {
			fmt.Printf("Failed to write spectemplate.yaml.tmpl file with error: %s", err.Error())
		}
	}

	{
		if err := ioutil.WriteFile(filepath.Join(dir, "returns.json"), []byte(res.Spec.Returns), os.ModePerm); err != nil {
			fmt.Printf("Failed to write returns.json file with error: %s", err.Error())
		}
	}

	{
		meta, err := yaml.Marshal(res.Metadata)
		if err != nil {
			fmt.Printf("Failed to marshal metadata to valid yaml file file with error: %s", err.Error())
		}
		if err := ioutil.WriteFile(filepath.Join(dir, "metadata.yaml"), meta, os.ModePerm); err != nil {
			fmt.Printf("Failed to write metadata.yaml file with error: %s", err.Error())
		}
	}
	return nil
}

func createDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
