package main

import (
	"encoding/json"
	"sort"
  _ "github.com/aws/aws-xray-sdk-go/xray"

  // Importing the plugins enables collection of AWS resource information at runtime.
  // Every plugin should be imported after "github.com/aws/aws-xray-sdk-go/xray" library.

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nathanmalishev/go-lambda-example/repository"
)

type AlphabeticalPersonList []repository.Person

func (a AlphabeticalPersonList) Len() int           { return len(a) }
func (a AlphabeticalPersonList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a AlphabeticalPersonList) Less(i, j int) bool { return a[i].Name < a[j].Name }

// PeopleResponse
type PeopleResponse struct {
	People []repository.Person `json:"data"`
}

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	list := repository.GetPeople()
	sort.Sort(AlphabeticalPersonList(list))
	body, err := json.Marshal(PeopleResponse{list})

	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Unable to marshal JSON", StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}
