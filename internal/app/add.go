package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new item to selected list",
	Long: `Add a new item to selected list. You can also specify the item's priority.
For example:

todo add "ExampleItem" medium`,
	Run: addItem,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addItem(cmd *cobra.Command, args []string) {
	itemname := args[0]
	priority := args[1]
	listname, err := rootSvc.AddToSelected(itemname, priority)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\"%s\" added to %s.\n", itemname, listname)
}
