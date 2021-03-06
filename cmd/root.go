// Code generated by cli-generator; DO NOT EDIT.
package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

var cnf *viper.Viper = viper.New()

var rootCmdOptions struct {
	verbose bool
}

var rootCmd = &cobra.Command{
	Use:     "step-generator",
	Version: "0.11.0",
	Long:    "Quickly generate Codefresh step-types",
	PreRun: func(cmd *cobra.Command, args []string) {

		cnf.Set("verbose", rootCmdOptions.verbose)

	},
}

// Execute - execute the root command
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
	cnf.BindEnv("verbose", "VERBOSE")
	cnf.SetDefault("verbose", false)

	rootCmd.PersistentFlags().BoolVar(&rootCmdOptions.verbose, "verbose", cnf.GetBool("verbose"), "Set to see more logs [$VERBOSE]")
}
