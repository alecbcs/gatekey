package gatekey

import (
	"io/ioutil"
	"log"
	"net/http"
)

// NewKey submits a request to Gatekey to generate and return another key.
func NewKey(location string, username string, password string, jid string, mid string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", (location + "create/?mid=" + jid + "&jid=" + mid), nil)
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	return s
}
