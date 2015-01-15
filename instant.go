package tflcountdown

import (
	"log"
	"net/http"
	"net/url"
)

type InstantAPI struct {
	url *url.URL
}

func MakeInstantAPI(urlStr string) (*InstantAPI, error) {
	url, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	return &InstantAPI{
		url: url,
	}, nil
}

func (api *InstantAPI) MakeRequest(r Request) (chan Message, chan error) {
	if r.ReturnList == nil {
		f := NewDefaultFieldMap()
		r.ReturnList = &f
	}

	fields := r.ReturnList

	cloneUrl := *api.url
	cloneUrl.RawQuery = r.Encode().Encode()
	cloneUrlStr := cloneUrl.String()

	log.Println("Requesting", cloneUrlStr)

	resp, err := http.Get(cloneUrlStr)
	if err != nil {
		msgChan := make(chan Message)
		close(msgChan)

		errChan := make(chan error)
		errChan <- err
		close(errChan)

		return msgChan, errChan
	}

	msgChan, inErrChan := DecodeFromJson(resp.Body, *fields)
	errChan := make(chan error)

	go func() {
		for {
			err := <-inErrChan
			resp.Body.Close()
			errChan <- err
		}
	}()

	return msgChan, errChan
}
