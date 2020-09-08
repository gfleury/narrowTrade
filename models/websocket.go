package models

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WSMessage struct {
	MsgID         uint64
	reserved0     []byte
	RIDSize       uint8
	RID           []byte
	PayloadFormat uint8
	PayloadSize   uint32
	Payload       []byte
}

func (w *WSMessage) Parse(data []byte) error {
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.LittleEndian, &w.MsgID)
	if err != nil {
		return err
	}

	w.reserved0 = make([]byte, 2)
	_, err = r.Read(w.reserved0)
	if err != nil {
		return err
	}

	err = binary.Read(r, binary.LittleEndian, &w.RIDSize)
	if err != nil {
		return err
	}

	w.RID = make([]byte, w.RIDSize)
	n, err := r.Read(w.RID)
	if err != nil {
		return err
	}
	if n != int(w.RIDSize) {
		return fmt.Errorf("Unable to read enough bytes Expected: %d Read: %d", w.PayloadSize, n)
	}

	err = binary.Read(r, binary.LittleEndian, &w.PayloadFormat)
	if err != nil {
		return err
	}

	err = binary.Read(r, binary.LittleEndian, &w.PayloadSize)
	if err != nil {
		return err
	}

	w.Payload = make([]byte, w.PayloadSize)
	n, err = r.Read(w.Payload)
	if err != nil {
		return err
	}
	if n != int(w.PayloadSize) {
		return fmt.Errorf("Unable to read enough bytes Expected: %d Read: %d", w.PayloadSize, n)
	}

	return nil
}

func NewWSMessage(data []byte) (*WSMessage, error) {
	m := &WSMessage{}
	err := m.Parse(data)
	return m, err
}

func createSubscriptioWebSocket(authorizationToken, contextID string) (chan *WSMessage, *websocket.Conn, error) {
	header := http.Header{}
	header.Add("Authorization", fmt.Sprintf("Bearer %s", authorizationToken))
	c, _, err := websocket.DefaultDialer.Dial(
		fmt.Sprintf(
			"wss://streaming.saxobank.com/sim/openapi/streamingws/connect?contextId=%s", contextID),
		header)
	if err != nil {
		return nil, nil, err
	}

	msgs := make(chan *WSMessage)

	go func() {
		for {
			_, data, err := c.ReadMessage()
			if err != nil {
				log.Printf("WS %s: error: %v", contextID, err)
				return
			}

			message, err := NewWSMessage(data)
			if err != nil {
				log.Printf("WS %s: error: %v", contextID, err)
				return
			}
			msgs <- message
		}
	}()

	return msgs, c, nil
}
