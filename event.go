package main

type Event struct {
	Version string   `json:"version"`
	Session struct{} `json:"session"`
	Request Request  `json:"request"`
}
