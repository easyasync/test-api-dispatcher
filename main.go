package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

// Handler is the main entry point for Lambda. Receives a proxy request and
// returns a proxy response
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// stdout and stderr are sent to AWS CloudWatch Logs
	r := gin.Default()

	r.POST("/test/api/hello", handleMessage)
	ginLambda = ginadapter.New(r)

	return ginLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}

//Message is struct representing http message
type Message struct {
	Message string `json:"message"`
}

func handleMessage(c *gin.Context) {
	message := Message{}

	err := c.ShouldBind(&message)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "invalid request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message.Message,
	})
}
