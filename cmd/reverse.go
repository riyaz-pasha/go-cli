/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// reverseCmd represents the reverse command
var reverseCmd = &cobra.Command{
	Use:   "reverse",
	Short: "reverse of the string",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("reverse called")
		fmt.Println(reverse(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(reverseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reverseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reverseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func reverse(input string) (result string) {
	for _, c := range input {
		result = string(c) + result
	}
	return result
}
