package main

import (
	"context"
	"encoding/json"
	"fmt"
)

func Handler(ctx context.Context, input []byte) (*Response, error) {
	var event Event
	err := json.Unmarshal(input, &event)
	if err != nil {
		return &Response{
			Version: event.Version,
			Session: event.Session,
			Result: Result{
				Text:       fmt.Sprintf("an error has occurred when parsing event: %v", err),
				EndSession: false,
			},
		}, nil
	}

	text := "Hello! I'll repeat anything you say to me."
	if event.Request.OriginalUtterance != "" {
		text = event.Request.OriginalUtterance
	}

	return &Response{
		Version: event.Version,
		Session: event.Session,
		Result: Result{
			Text:       text,
			EndSession: false,
		},
	}, nil
}
