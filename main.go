package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "spotHeroProject/docs" // docs is generated by Swag CLI, you have to import it.
	"spotHeroProject/routes"
)

var gorillaMuxLambda *gorillamux.GorillaMuxAdapter

func init() {
	fmt.Printf("Mux start-")

	router := routes.Init()
	router.PathPrefix("/v2/swagger").Handler(httpSwagger.WrapHandler)

	gorillaMuxLambda = gorillamux.New(router)
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return gorillaMuxLambda.Proxy(req)
}

// @title Spot Hero
// @version 2.0
// @description Implement spot hero
// @contact.name Mohammad Dehghanpour
// @contact.email m.dehghanpour10@gmail.com
// @host rxzgqi6zfc.execute-api.us-west-2.amazonaws.com/api
// @BasePath /
func main() {
	lambda.Start(Handler)
}
