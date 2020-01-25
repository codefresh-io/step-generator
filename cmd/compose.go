// Code generated by cli-generator; DO NOT EDIT.
package cmd

import (
	handler "github.com/codefresh/step-generator/pkg/compose"

	"github.com/spf13/cobra"
)

var composeCmdOptions struct {
	leftDelimiter     string
	rightDelimiter    string
	specFile          string
	argumentsJsonFile string
	returnsJsonFile   string
	descriptionFile   string
	version           string
	official          bool
	tags              []string
	maintainerName    string
	maintainerEmail   string
	account           string
	isPublic          bool
	out               string
}

var composeCmd = &cobra.Command{
	Use: "compose",
	Args: func(cmd *cobra.Command, args []string) error {
		var validators []func(cmd *cobra.Command, args []string) error
		validators = append(validators, cobra.ExactArgs(1))
		for _, v := range validators {
			if err := v(cmd, args); err != nil {
				return err
			}
		}
		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		h := &handler.Handler{}
		return h.Handle(cnf)
	},
	Long: "Quickly generate Codefresh step-types",
	PreRun: func(cmd *cobra.Command, args []string) {
		cnf.Set("name", args)
		rootCmd.PreRun(cmd, args)

		cnf.Set("leftDelimiter", composeCmdOptions.leftDelimiter)

		cnf.Set("rightDelimiter", composeCmdOptions.rightDelimiter)

		cnf.Set("specFile", composeCmdOptions.specFile)

		cnf.Set("argumentsJsonFile", composeCmdOptions.argumentsJsonFile)

		cnf.Set("returnsJsonFile", composeCmdOptions.returnsJsonFile)

		cnf.Set("descriptionFile", composeCmdOptions.descriptionFile)

		cnf.Set("version", composeCmdOptions.version)

		cnf.Set("official", composeCmdOptions.official)

		cnf.Set("tags", composeCmdOptions.tags)

		cnf.Set("maintainerName", composeCmdOptions.maintainerName)

		cnf.Set("maintainerEmail", composeCmdOptions.maintainerEmail)

		cnf.Set("account", composeCmdOptions.account)

		cnf.Set("isPublic", composeCmdOptions.isPublic)

		cnf.Set("out", composeCmdOptions.out)

	},
}

func init() {
	cnf.SetDefault("leftDelimiter", "[[")

	composeCmd.PersistentFlags().StringVar(&composeCmdOptions.leftDelimiter, "left-delimiter", cnf.GetString("leftDelimiter"), "Left delimiter to use in spec file")
	cnf.SetDefault("rightDelimiter", "]]")

	composeCmd.PersistentFlags().StringVar(&composeCmdOptions.rightDelimiter, "right-delimiter", cnf.GetString("rightDelimiter"), "Right delimiter to use in spec file")

	composeCmd.PersistentFlags().StringVar(&composeCmdOptions.specFile, "spec-file", cnf.GetString("specFile"), "Path to Codefresh template spec file")

	composeCmd.PersistentFlags().StringVar(&composeCmdOptions.argumentsJsonFile, "arguments-json-file", cnf.GetString("argumentsJsonFile"), "Path to Codefresh JSON-schema arguments file")

	composeCmd.PersistentFlags().StringVar(&composeCmdOptions.returnsJsonFile, "returns-json-file", cnf.GetString("returnsJsonFile"), "Path to Codefresh JSON-schema returns file")

	composeCmd.PersistentFlags().StringVar(&composeCmdOptions.descriptionFile, "description-file", cnf.GetString("descriptionFile"), "Path to step-type description file")
	cnf.SetDefault("version", "0.1.0")

	composeCmd.PersistentFlags().StringVar(&composeCmdOptions.version, "version", cnf.GetString("version"), "Version of the step-type")

	composeCmd.PersistentFlags().BoolVar(&composeCmdOptions.official, "official", cnf.GetBool("official"), "Set if the step-type if an official one")

	composeCmd.PersistentFlags().StringArrayVar(&composeCmdOptions.tags, "tags", cnf.GetStringSlice("tags"), "")

	composeCmd.PersistentFlags().StringVar(&composeCmdOptions.maintainerName, "maintainer-name", cnf.GetString("maintainerName"), "")

	composeCmd.PersistentFlags().StringVar(&composeCmdOptions.maintainerEmail, "maintainer-email", cnf.GetString("maintainerEmail"), "")

	composeCmd.PersistentFlags().StringVar(&composeCmdOptions.account, "account", cnf.GetString("account"), "Name of the account that the step-type is maintained by")

	composeCmd.PersistentFlags().BoolVar(&composeCmdOptions.isPublic, "is-public", cnf.GetBool("isPublic"), "Set to mark step-type as public and accesiable by other accounts as well")

	composeCmd.PersistentFlags().StringVar(&composeCmdOptions.out, "out", cnf.GetString("out"), "Name of the output file, if empty the result will be printed to stdout")
	rootCmd.AddCommand(composeCmd)
}
