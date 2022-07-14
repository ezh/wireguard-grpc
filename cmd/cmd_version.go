package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints version",
	Long:  "Prints the program’s version information.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		git := ""
		r := strings.TrimSpace(Revision)

		if len(r) > 0 {
			git = fmt.Sprintf(" (r%s) %s", r, BuildCommit)
		}

		fmt.Printf("%s Version: %s%s built at %s", AppName, Version, git, BuildDate) //nolint:forbidigo
	},
}
