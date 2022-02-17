/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <num>",
	Short: "Delete the task at given index",
	Run: func(cmd *cobra.Command, args []string) {
		index := args[0]
		client := &http.Client{}
		req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8080/tasks/%s", index), nil)
		if err != nil {
			log.Fatalln(err)
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
			return
		}
		defer resp.Body.Close()
		
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
