/*
Copyright © 2021 SUDIPTO BARAL <sudiptobaral.me@gmail.com>

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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	color "github.com/TwinProduction/go-color"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// motivationCmd represents the motivation command
var motivationCmd = &cobra.Command{
	Use:   "motivation",
	Short: "Give you a single motivation",
	Long:  `Fetch a motivation quote from the api and dispalays to you. Every time it will generate new motivation.`,
	Run: func(cmd *cobra.Command, args []string) {
		showMotivation()
	},
}

// Resposne struct. Holds the json data
// uses nested struct syntax
type Resposne struct {
	Slip struct {
		ID     int    `json:"id"`
		Advice string `json:"advice"`
	}
}

func showMotivation() {
	url := "https://api.adviceslip.com/advice"
	dataBytes := getDataFromAPI(url)
	motivation := Resposne{}

	if err := json.Unmarshal(dataBytes, &motivation); err != nil {
		fmt.Printf("Could to unmarshal the databytes!")
	}
	fmt.Println(color.Ize(color.Green, string(motivation.Slip.Advice)))
}

func getDataFromAPI(url string) []byte {
	var responseByte []byte
	resp, err := http.Get(url)
	if err != nil {
		log.Error(err)
	} else {

		responseByte, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error(err)
		}

	}
	// Read the resp and convert into a byte array

	return responseByte
}

func init() {
	rootCmd.AddCommand(motivationCmd)
}
