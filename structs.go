package main

import (
	"time"
)

type RandomizerResponse struct {
	Categories []array   `json: "categories, omitempty"`
	CreateDate time.Time `json: "created_at"`
	IconUrl    string    `json: "icon_url"`
	Id         string    `json: "id"`
	UpdateAt   time.Time `json: "update_at"`
	Url        string    `json: "url"`
	Value      string    `json: "value"`
}

type ArrayRandomizer struct {
	Data []RandomizerResponse `json: "data"`
}

type array struct {
}
