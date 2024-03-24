/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"brandonbraner.com/openalex-cli/src/files"
	"github.com/spf13/cobra"
)

// getValidFilenamesCmd represents the getValidFilenames command
var getValidFilenamesCmd = &cobra.Command{
	Use:   "get-valid-filenames",
	Short: "Get a list of valid filenames to pass to the openalex-cli commands",
	Long:  `Get a list of valid filenames to pass to the openalex-cli commands.`,
	Run: func(cmd *cobra.Command, args []string) {
		filenames := files.GetValidFilenames()
		for _, filename := range filenames {
			cmd.Println(filename)
		}
	},
}

func init() {
	rootCmd.AddCommand(getValidFilenamesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getValidFilenamesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getValidFilenamesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
