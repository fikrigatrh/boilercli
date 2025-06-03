
package cmd

import (
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "boilercli",
    Short: "Boilerplate CLI generator",
}

func Execute() {
    cobra.CheckErr(rootCmd.Execute())
}
