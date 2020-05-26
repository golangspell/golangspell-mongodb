package cmd

import (
	"fmt"

	"github.com/golangspell/golangspell-mongodb/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "golangspell-mongodb-version",
	Short: "golangspell-mongodb version number",
	Long:  `Shows the golangspell-mongodb current installed version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("golangspell-mongodb v%s -- HEAD\n", config.Version)
	},
}
