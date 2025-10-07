package app

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a to-do list",
	Long: `Delete an existing to-do list with given name.
For example:

todo delete ExampleList`,
	Run: deleteList,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteList(cmd *cobra.Command, args []string) {
	listname := args[0]

	var confirm string
	fmt.Printf(
		"Are you sure you want to delete this list: %s? [Y/n]: ",
		listname,
	)
	fmt.Scanln(&confirm)
	confirm = strings.ToLower(strings.TrimSpace(confirm))

	if confirm != "y" {
		return
	}

	err := rootSvc.Delete(listname)
	if err != nil {
		panic(err)
	}
	fmt.Printf("List %s deleted.\n", listname)
}
