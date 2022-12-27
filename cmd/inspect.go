/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"regexp"
)

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("inspect called")
		isNumbers, _ := cmd.Flags().GetBool("numbers")
		isLetter, _ := cmd.Flags().GetBool("letters")
		if isNumbers {
			fmt.Println("Total no of numbers are : ", getNumbersCount(args[0]))
		} else if isLetter {
			fmt.Println("Total no of letters are : ", getLettersCount(args[0]))
		}
	},
}

func getLettersCount(str string) int {
	re := regexp.MustCompile("[^a-z]")
	return len(re.ReplaceAllString(str, ""))
}

func getNumbersCount(str string) int {
	re := regexp.MustCompile("[^0-9]")
	return len(re.ReplaceAllString(str, ""))
}

func init() {
	rootCmd.AddCommand(inspectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inspectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inspectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	inspectCmd.Flags().BoolP("numbers", "n", false, "Returns numbers count from the string")
	inspectCmd.Flags().BoolP("letters", "l", false, "Returns letters count from the string")
}
