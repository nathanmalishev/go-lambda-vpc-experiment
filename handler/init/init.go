package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
	// Importing the plugins enables collection of AWS resource information at runtime.
	// Every plugin should be imported after "github.com/aws/aws-xray-sdk-go/xray" library.

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-xray-sdk-go/xray"
)

type person struct {
	name string
	age  int
}

func handleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	SQL_URI := os.Getenv("MY_SQL_URI")
	if SQL_URI == "" {
		log.Println("No SQL_URI defined")
		return events.APIGatewayProxyResponse{Body: string("No SQL_URI defined"), StatusCode: 500}, nil
	}

	fmt.Println(os.Getenv("MY_SQL_URI"))
	db, err := xray.SQL("mysql", os.Getenv("MY_SQL_URI"))
	defer db.Close()

	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	_, err = db.Exec(ctx, "CREATE DATABASE mydb")
	if err != nil {
		//return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}
	_, err = db.Exec(ctx, "USE mydb")
	if err != nil {
		//return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	_, err = db.Exec(ctx, `CREATE TABLE IF NOT EXISTS customers (name text, age integer, address text)`)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}
	return events.APIGatewayProxyResponse{Body: string("Init success"), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}
