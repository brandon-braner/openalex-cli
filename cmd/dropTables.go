/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"brandonbraner.com/openalex-cli/src/database"
	"fmt"

	"github.com/spf13/cobra"
)

// dropTablesCmd represents the dropTables command
var dropTablesCmd = &cobra.Command{
	Use:   "drop-tables",
	Short: "Drop all tables and the openalex schema from your database",
	Long:  `Drop all tables and the openalex schema from your database.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dropTables called")
		database.DropTables()
	},
}

func init() {
	rootCmd.AddCommand(dropTablesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dropTablesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dropTablesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
