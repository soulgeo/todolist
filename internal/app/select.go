package app

import (
	"github.com/spf13/cobra"
)

var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Set a list as selected",
	Long: `Set a list as selected. Other commands perform actions on selected list by default
For example:

todo select ExampleList`,
	Run: selectList,
}

func init() {
	rootCmd.AddCommand(selectCmd)
}

func selectList(cmd *cobra.Command, args []string) {
	listname := args[0]
	err := rootSvc.Select(listname)
	if err != nil {
		panic(err)
	}
}
