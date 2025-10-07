package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new to-do list",
	Long: `Create a new to-do list with given name.
For example:

todo create ExampleList`,
	Run: createList,
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func createList(cmd *cobra.Command, args []string) {
	listname := args[0]
	err := rootSvc.Create(listname)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created list: %s.\n", listname)
}
