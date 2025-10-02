package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listsCmd = &cobra.Command{
	Use:   "lists",
	Short: "Show a list of all to-do lists",
	Long:  `Display a list of the names of all user-created to-do lists`,
	Run:   showLists,
}

func init() {
	rootCmd.AddCommand(listsCmd)
}

func showLists(cmd *cobra.Command, args []string) {
	lists, err := rootSvc.GetAllLists()
	if err != nil {
		panic(err)
	}
	for _, l := range lists {
		fmt.Printf("- %s\n", l)
	}
}
