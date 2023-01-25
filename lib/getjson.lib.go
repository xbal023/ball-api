package lib

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	)
	
var output map[string]interface{}
func FetchJson(url string) interface{} {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error());
		os.Exit(1);
	}
	data, err := ioutil.ReadAll(res.Body);
	if err != nil {
		log.Fatal(err.Error())
	}
	json.Unmarshal(data, &output)
	return output
}