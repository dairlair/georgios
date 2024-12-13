package main

import (
	"context"
	"encoding/json"
	"fmt"
)

type Event struct {
	Version string   `json:"version"`
	Session struct{} `json:"session"`
	Request struct {
		OriginalUtterance string `json:"original_utterance"`
	} `json:"request"`
}

type Response struct {
	Version string   `json:"version"`
	Session struct{} `json:"session"`
	Result  struct {
		Text       string `json:"text"`
		EndSession bool   `json:"end_session"`
	} `json:"response"`
}

func ParrotHandler(ctx context.Context, event []byte) (*Response, error) {
	var input Event
	err := json.Unmarshal(event, &input)
	if err != nil {
		return &Response{
			Version: input.Version,
			Session: input.Session,
			Result: struct {
				Text       string `json:"text"`
				EndSession bool   `json:"end_session"`
			}{
				Text:       fmt.Sprintf("an error has occurred when parsing event: %v", err),
				EndSession: false,
			},
		}, nil
		// return nil, fmt.Errorf("an error has occurred when parsing event: %v", err)
		// return nil, fmt.Errorf("an error has occurred when parsing event: %v", err)
	}

	text := "Hello! I'll repeat anything you say to me."
	if input.Request.OriginalUtterance != "" {
		text = input.Request.OriginalUtterance
	}

	return &Response{
		Version: input.Version,
		Session: input.Session,
		Result: struct {
			Text       string `json:"text"`
			EndSession bool   `json:"end_session"`
		}{
			Text:       text,
			EndSession: false,
		},
	}, nil
}
