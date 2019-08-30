/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"io/ioutil"
	"net/http"
	"time"

	"github.com/briandowns/spinner"
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"github.com/torusresearch/bijson"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info [name of maker]",
	Short: "Info about maker",
	Long:  `Product Hunt ID , Twitter ID and products made by the maker.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("info called")
		makerInfo(args[0])
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Product struct {
	Name      string  `json:"name"`
	Url       string  `json:"url"`
	Image_url string  `json:"image_url"`
	Upvotes   float64 `json:"upvotes"`
}

type Maker struct {
	Producthunt_username string  `json:"producthunt_username"`
	Name                 string  `json:"name"`
	Image_url            string  `json:"image_url"`
	Twitter_username     string  `json:"twitter_username"`
	Rank                 float64 `json:"rank"`
	Upvotes              float64 `json:"upvotes"`
	Comments             float64 `json:"comments"`
	Twitter_followers    float64 `json:"twitter_followers"`
	Products             []Product
	Inputs               bijson.RawMessage
	Products_by_year     bijson.RawMessage
}

func makerInfo(maker string) {
	// https://makerrank.co/@agiche.json
	url := "https://makerrank.co/@" + maker + ".json"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("API has some problems, please check your internet")
	}
	s := spinner.New(spinner.CharSets[35], 200*time.Millisecond) // Build our new spinner
	s.Start()                                                    // Start the spinner
	time.Sleep(1 * time.Second)                                  // Run for some time to simulate work
	s.Stop()
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Invalid Response from API")
	}

	// fmt.Print(string(b))
	var m Maker

	if err := bijson.Unmarshal(b, &m); err != nil {
		fmt.Errorf("Error in Unmarshal of JSON")
	}

	// fmt.Print(m.Name)
	myFigure := figure.NewFigure(m.Name, "", true)
	myFigure.Print()
	fmt.Print("\nProduct Hunt: ", m.Producthunt_username)
	fmt.Print("\nTwitter:      ", m.Twitter_username)
	fmt.Print("\nUpvotes:      ", m.Upvotes)
	fmt.Print("\nProduct          |          URL         | Upvotes \n-------------------------------------------------------------------")
	for _, p := range m.Products {
		fmt.Print("\n", p.Name, "   |  ", p.Url, "  |  ", p.Upvotes)
	}

	// var p []Product
	// if err := bijson.Unmarshal(m.Products, &p); err != nil {
	// 	fmt.Errorf("Error in Unmarshal of JSON")
	// }
	// fmt.Print("\nProducts:", p)
	// // fmt.Printf("\n%+v", m)

}
