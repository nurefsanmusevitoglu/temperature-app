package customerror

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func New(status int) events.APIGatewayProxyResponse {
	body := map[string]string{
		"message": http.StatusText(status),
	}

	bodyBytes, _ := json.Marshal(body)

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(bodyBytes),
	}
}
