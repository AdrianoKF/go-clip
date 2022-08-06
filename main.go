/*
Copyright © 2022 Adrian Rumpold <a.rumpold@gmail.com>
*/
package main

import (
	"fmt"

	"github.com/AdrianoKF/go-clip/cmd"
	"github.com/AdrianoKF/go-clip/internal/util"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version will output the current build information",
	Long:  ``,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Version: %s\n", version)
		fmt.Printf("Commit: %s\n", commit)
		fmt.Printf("Date: %s\n", date)
		fmt.Printf("Built by: %s\n", builtBy)
	},
}

func main() {
	util.InitializeLogging(true)

	cmd.RootCmd.AddCommand(versionCmd)
	cmd.Execute()
}
