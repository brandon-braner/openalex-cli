/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"brandonbraner.com/openalex-cli/src/database"
	"fmt"

	"github.com/spf13/cobra"
)

// dropIndexesCmd represents the dropIndexes command
var dropIndexesCmd = &cobra.Command{
	Use:   "drop-indexes",
	Short: "Drop all indexes from your database",
	Long:  `Drop all indexes to make new inserts faster. Don't forget to readd them when done.'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dropIndexes called")
		database.DropIndexes()
	},
}

func init() {
	rootCmd.AddCommand(dropIndexesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dropIndexesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dropIndexesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
