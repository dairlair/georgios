package main

import (
	"context"
	"encoding/json"
	"fmt"
)

// The input JSON document is automatically converted to this type of object
type Request struct {
	Name string `json:"name"`
}

type ResponseBody struct {
	Greetign string `json:"greeting"`
}

func Handler(ctx context.Context, request *Request) ([]byte, error) {
	// The function logs contain the values of the invocation context and request body
	// fmt.Println("context", ctx)
	// fmt.Println("request", request)

	// The object containing the response body is converted to an array of bytes
	body, err := json.Marshal(&ResponseBody{
		Greetign: fmt.Sprintf("Hello, dear %s!", request.Name),
	})

	if err != nil {
		return nil, err
	}

	// The response body must be returned as an array of bytes
	return body, nil
}
