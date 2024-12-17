package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dairlair/georgios/pkg/yandexalicesdk"
)

func Handler(ctx context.Context, input []byte) (*yandexalicesdk.Response, error) {
	var event yandexalicesdk.Event
	err := json.Unmarshal(input, &event)
	if err != nil {
		return &yandexalicesdk.Response{
			Version: event.Version,
			Session: event.Session,
			Result: yandexalicesdk.Result{
				Text:       fmt.Sprintf("an error has occurred when parsing event: %v", err),
				EndSession: false,
			},
		}, nil
	}

	text := "Hello! I'll repeat anything you say to me."
	if event.Request.OriginalUtterance != "" {
		text = event.Request.OriginalUtterance
	}

	return &yandexalicesdk.Response{
		Version: event.Version,
		Session: event.Session,
		Result: yandexalicesdk.Result{
			Text:       text,
			EndSession: false,
		},
	}, nil
}
