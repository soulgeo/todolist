package app

import (
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
	err := rootSvc.Delete(listname)
	if err != nil {
		panic(err)
	}
}
