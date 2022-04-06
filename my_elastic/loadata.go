package my_elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
)

func LoadData() {
	var spaceCrafts []map[string]interface{}
	pageNumber := 0
	for {
		response, err := http.Get("http://stapi.co/api/v1/rest/spacecraft/search?pageSize=100&pageNumber=" + strconv.Itoa(pageNumber))
		if err != nil {
			fmt.Println("got error performing get request on http://stapi.co/api/v1/rest/spacecraft/ ", err)
		}
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("got error reading response body ", err)
		}
		defer response.Body.Close()
		var result map[string]interface{}
		json.Unmarshal(body, &result)

		page := result["page"].(map[string]interface{})
		totalPages := int(page["totalPages"].(float64))

		crafts := result["spacecrafts"].([]interface{})

		for _, craftInterface := range crafts {
			craft := craftInterface.(map[string]interface{})
			spaceCrafts = append(spaceCrafts, craft)
		}

		pageNumber++
		if pageNumber >= totalPages {
			break
		}
	}

	for _, data := range spaceCrafts {
		uid, _ := data["uid"].(string)
		jsonString, err := json.Marshal(data)
		if err != nil {
			fmt.Println("error when marshalling data ", err)
		}
		request := esapi.IndexRequest{Index: "stsc", DocumentID: uid, Body: strings.NewReader(string(jsonString))}
		request.Do(context.Background(), es)
	}
	print(len(spaceCrafts), " spacecraft read\n")
}
