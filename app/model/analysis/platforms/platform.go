package model_platform

import "encoding/json"

type PostingPlatform struct {
	Platforms json.RawMessage `json:"platforms_data"`
}