package app

import (
	"strconv"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an item from selected list",
	Long: `Remove an item at given index from selected list
For example:

todo remove 2`,
	Run: removeItem,
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

func removeItem(cmd *cobra.Command, args []string) {
	index, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}
	err = rootSvc.RemoveFromSelected(index - 1)
	if err != nil {
		panic(err)
	}
}
