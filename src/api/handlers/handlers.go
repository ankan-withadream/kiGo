package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"kiGo/internal/config"
	"kiGo/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var aiclient = &http.Client{}

func Handle_empty(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		c.Writer.Write([]byte("Get Response"))
	case http.MethodPost:
		fmt.Fprint(c.Writer, "POST request")
	default:
		fmt.Fprintf(c.Writer, "Unsupported method: %s", c.Request.Method)
	}
}

func Handle_kigo(c *gin.Context) {
	fmt.Println(c.Request.Method)

	// if GET method, send this
	if c.Request.Method == http.MethodGet {
		c.Writer.Write([]byte("Ei cholche Go :)"))
	}

	// if POST method, do this
	if c.Request.Method == http.MethodPost {

		// read req body as json
		jsondata, err := io.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println("Error reading request body: ", err)
		}

		// json to payload data
		data := &services.Payload{}
		err = json.Unmarshal(jsondata, data)
		if err != nil {
			fmt.Println("Error parsing json data: ", err)
		}

		// request to AI server and get response
		resp, err := services.Ask_AI(config.APP_CONFIG.AI_CLIENT_API_ENDPOINT, data, aiclient)
		if err != nil {
			fmt.Println("Error from the AI server: ", err)
		}

		// send response to client
		// fmt.Fprint(c.Writer, resp)
		// println(resp)
		c.JSON(http.StatusOK, resp)

	}
}
