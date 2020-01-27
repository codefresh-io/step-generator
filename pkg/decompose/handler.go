package decompose

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

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
	filepath := cnf.GetString("file")
	if filepath == "" {
		return fmt.Errorf("No path to file")
	}
	fileBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	st, err := steptype.UnmarshalStepType(fileBytes)
	if err != nil {
		return err
	}

	outdir := cnf.GetString("outDirectory")
	if outdir == "" {
		outdir = fmt.Sprintf("generated-%s", st.Metadata.Name)
	}
	err = createDirIfNotExist(outdir)
	if err != nil {
		return err
	}
	marshalAndDump(st.Metadata, path.Join(outdir, "metadata.yaml"))
	if err != nil {
		return err
	}

	dumpTofFile([]byte(st.Spec.Arguments), path.Join(outdir, "arguments.json"))
	if err != nil {
		return err
	}

	dumpTofFile([]byte(st.Spec.Returns), path.Join(outdir, "returns.json"))
	if err != nil {
		return err
	}

	dumpTofFile([]byte(st.Spec.StepsTemplate), path.Join(outdir, "spectemplate.yaml.tmpl"))
	if err != nil {
		return err
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

func marshalAndDump(data interface{}, path string) error {
	bytes, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	return dumpTofFile(bytes, path)
}

func dumpTofFile(bytes []byte, path string) error {
	return ioutil.WriteFile(path, bytes, os.ModePerm)
}
