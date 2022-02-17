/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	//"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)
type JsonTask struct {
	Priority int `json:"priority"`
	Data string `json:"task"`
}
// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		priority, _ := strconv.Atoi(args[1])
		newTask := JsonTask{
			Data: task,
			Priority: priority,
		}
		postBody, _ := json.MarshalIndent(newTask, "", "  ")
		responseBody := bytes.NewBuffer(postBody)
		resp, err := http.Post("http://localhost:8080/tasks","application/json", responseBody)
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body)
		fmt.Println("Response: ",sb)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
