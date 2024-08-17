/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/silvioguiso/tri/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all todo items",
	Long:  `Lists all todo items so you can iterate over whatever still needs to be done. Filtering optional`,
	Run: func(cmd *cobra.Command, args []string) {
		readfile()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func readfile() {
	items, err := todo.ReadItems("/Users/silvi/.tridos.json")

	if err != nil {
		log.Printf("%v", err)
	}

	fmt.Println(items)
}
