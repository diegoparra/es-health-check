package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/aws/aws-lambda-go/lambda"
)

var resp string

func getResponse(site string) string {
	resp, err := http.Get(site)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)
	re := regexp.MustCompile("green")
	fmt.Println(re.FindString(bodyString))
	return re.FindString(bodyString)
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (string, error) {
	//resp := "green"
	if resp = getResponse(os.Getenv("ES_ENDPOINT")); resp == "green" {
		fmt.Println("ES Cluster Status: ", resp)
	} else if resp == "yellow" {
		fmt.Println("WARNING: Es cluster is in yellow state")
	} else {
		fmt.Println("CRITICAL: ES cluster is in Red state")
	}

	return fmt.Sprintf(resp), nil
}

func main() {
	lambda.Start(Handler)
}
