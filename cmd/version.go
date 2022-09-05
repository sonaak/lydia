/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// open the version file
		projectRoot := getProjectRoot()
		versionFile, err := os.Open(fmt.Sprintf("%s/VERSION", projectRoot))
		if err != nil {
			os.Exit(1)
		}

		// extract the version
		defer versionFile.Close()

		versionBytes, err := ioutil.ReadAll(versionFile)
		if err != nil {
			os.Exit(1)
		}

		fmt.Println(string(versionBytes))
	},
}

func getProjectRoot() string {
	// if the environment variable is defined, then use it
	if root, defined := os.LookupEnv("PROJECT_ROOT"); defined {
		return root
	}

	workingDir, err := os.Getwd()
	// if this fails, we should exit the program
	// as it is important for us to be able to retrieve
	// this value
	if err != nil {
		panic(err)
	}

	return workingDir
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
