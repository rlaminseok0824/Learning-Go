/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		SERVER := viper.GetString("server")
		PORT := viper.GetString("port")

		number, _ := cmd.Flags().GetString("tel")
		if number == "" {
			fmt.Println("Number is empty")
			return
		}
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			fmt.Println("Name is empty")
			return
		}
		surname, _ := cmd.Flags().GetString("surname")
		if surname == "" {
			fmt.Println("Surname is empty")
			return
		}

		URL := "http://" + SERVER + ":" + PORT + "/insert/"
		URL = URL + "/" + name + "/" + surname + "/" + number

		// Send request to server
		data, err := http.Get(URL)
		if err != nil {
			fmt.Println("**", err)
			return
		}

		// Check HTTP Status Code
		if data.StatusCode != http.StatusOK {
			fmt.Println("Status code:", data.StatusCode)
			return
		}

		// Read data
		responseData, err := io.ReadAll(data.Body)
		if err != nil {
			fmt.Println("*", err)
			return
		}

		fmt.Print(string(responseData))
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)
	insertCmd.Flags().StringP("name", "n", "", "Name value")
	insertCmd.Flags().StringP("surname", "s", "", "Surname value")
	insertCmd.Flags().StringP("tel", "t", "", "Telephone value")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// insertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// insertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
