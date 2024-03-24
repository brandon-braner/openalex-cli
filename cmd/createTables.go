/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"brandonbraner.com/openalex-cli/src/database"
	"fmt"
	"github.com/spf13/cobra"
)

// createTablesCmd represents the createTables command
var createTablesCmd = &cobra.Command{
	Use:   "create-tables",
	Short: "Creates the openalex tables in your databse",
	Long: `Creates the openalex tables in your database. 
	This command should be run once when you first set up openalex.
	The sql definition is a replica of https://raw.githubusercontent.com/ourresearch/openalex-documentation-scripts/main/openalex-pg-schema.sql
	without the indexes. Those should be added after to make queries insertion faster.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createTables called")
		database.CreateTables()
	},
}

func init() {
	rootCmd.AddCommand(createTablesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createTablesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createTablesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
