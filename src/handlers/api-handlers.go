package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"kiGo/config"
	"kiGo/src/utills"
	"net/http"
)

var aiclient = &http.Client{}

func Handle_empty(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Get Response"))
	case http.MethodPost:
		fmt.Fprint(w, "POST request")
	default:
		fmt.Fprintf(w, "Unsupported method: %s", r.Method)
	}
}

func Handle_kigo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, http.MethodGet)

	// if GET method, send this
	if r.Method == http.MethodGet {
		w.Write([]byte("Ei cholche Go :)"))
	}

	// if POST method, do this
	if r.Method == http.MethodPost {

		// read req body as json
		jsondata, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Errorf("Error reading request body: ", err)
		}

		// json to payload data
		data := &utills.Payload{}
		err = json.Unmarshal(jsondata, data)
		if err != nil {
			fmt.Errorf("Error parsing json data: ", err)
		}

		// request to AI server and get response
		resp, err := utills.Ask_AI(config.APP_CONFIG.AI_CLIENT_API_ENDPOINT, data, aiclient)
		if err != nil {
			fmt.Errorf("Error from the AI server: ", err)
		}

		// send response to client
		fmt.Fprint(w, resp)

	}
}
