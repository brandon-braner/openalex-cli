package cmd

import (
	"brandonbraner.com/openalex-cli/src/database"
	"fmt"

	"github.com/spf13/cobra"
)

// createIndexesCmd represents the createIndexes command
var createIndexesCmd = &cobra.Command{
	Use:   "create-indexes",
	Short: "Create indexes in the openalex database",
	Long:  `Create indexes on the openalex database. This should be run after all data has been imported as it will speed up insert queries significantly.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createIndexes called")
		database.CreateIndexes()
	},
}

func init() {
	rootCmd.AddCommand(createIndexesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createIndexesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createIndexesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
