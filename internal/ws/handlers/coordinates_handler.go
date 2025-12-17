package wshandlers

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	wsdto "go.back/internal/ws/dto"
	"go.back/pkg/customerror"
)

func HandleCoordinates(conn *websocket.Conn, data json.RawMessage) error {
	var coordinates wsdto.Coorditanes

	err := json.Unmarshal(data, &coordinates)

	if err != nil {
		return customerror.ParsingError
	}

	return nil
}
