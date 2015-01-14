package tflcountdown

import (
	"bufio"
	"encoding/json"
	"io"
)

func Decode(inp *TflArray, fields FieldMap) (Message, error) {
	if inp.Len() < 1 {
		return nil, ERROR_INVALID
	}

	messageTypeId := inp.AsTflMessageType()
	inp.Rewind()

	decoder, ok := registeredTflMessageTypes[messageTypeId]
	if !ok {
		return nil, ERROR_INVALID
	}

	return decoder.Decode(inp, fields)
}

func DecodeFromJson(baseRd io.Reader, fields FieldMap) (chan Message, chan error) {
	msgChan := make(chan Message)
	errChan := make(chan error)

	go func(rd *bufio.Reader, msgChan chan<- Message, errChan chan<- error) {
		for {
			line, err := rd.ReadBytes('\n')
			if err != nil && err != io.EOF {
				errChan <- err
				break
			} else if err == io.EOF {
				errChan <- nil
				break
			}

			intf := []interface{}{}
			err = json.Unmarshal(line, &intf)
			if err != nil {
				errChan <- err
				break
			}

			if msg, err := Decode(NewTflArray(intf), fields); err != nil {
				errChan <- err
				return
			} else {
				msgChan <- msg
			}
		}
		close(msgChan)
		close(errChan)
	}(bufio.NewReader(baseRd), msgChan, errChan)
	return msgChan, errChan
}
