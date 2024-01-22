package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getRandomizer(c http.ResponseWriter, d *http.Request) {

	var arrayFinal ArrayRandomizer

	for i := 0; i < counts.contador; i++ {
		response, err := http.Get(endpoints.apiurl)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(responseData))

		var randomizerResponse RandomizerResponse
		json.Unmarshal(responseData, &randomizerResponse)
		fmt.Println(randomizerResponse)

		arrayFinal.Data = append(arrayFinal.Data, randomizerResponse)
	}

	c.Header().Set("Content-Type", "application/json")
	json.NewEncoder(c).Encode(arrayFinal)
}
