package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show the currently selected list",
	Long:  `Display all items of list that is currently selected.`,
	Run:   showList,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func showList(cmd *cobra.Command, args []string) {
	list, err := rootSvc.GetSelectedList()
	if err != nil {
		panic(err)
	}
	for i, item := range list.Items {
		fmt.Printf("%d: %s (Priority: %s)\n", i+1, item.Title, item.Priority)
	}
}
