package app

import (
	"github.com/soulgeo/todolist/internal/todo"
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
	name := args[0]
	items := make([]todo.Item, 0)
	list := todo.TodoList{Name: name, Items: items}
	err := rootSvc.Create(list)
	if err != nil {
		panic(err)
	}
}
