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
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// loggedCmd represents the logged command
var loggedCmd = &cobra.Command{
	Use:   "logged",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := "/logged"
		user := User{Username: username, Password: password}

		// bytes.Buffer is both a Reader and a Writer
		buf := new(bytes.Buffer)
		err := user.ToJSON(buf)
		if err != nil {
			fmt.Println("JSON:", err)
			return
		}

		req, err := http.NewRequest(http.MethodGet, SERVER+PORT+endpoint, buf)
		if err != nil {
			fmt.Println("GetAll – Error in req: ", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		c := &http.Client{
			Timeout: 15 * time.Second,
		}

		resp, err := c.Do(req)
		if err != nil {
			fmt.Println("Do:", err)
			return
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Println(resp)
			return
		}

		var users = []User{}
		SliceFromJSON(&users, resp.Body)
		data, err := PrettyJSON(users)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print(data)
	},
}

func init() {
	rootCmd.AddCommand(loggedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loggedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loggedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
