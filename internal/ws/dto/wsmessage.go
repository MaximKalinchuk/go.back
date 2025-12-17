package wsdto

import "encoding/json"

type WsMessage struct {
	Type string          `json:"type" binding:"required"`
	Data json.RawMessage `json:"data" binding:"required"`
}

type Coorditanes struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
