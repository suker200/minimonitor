package data_report

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func pushbullet_report() {
	uri_path := "https://api.pushbullet.com/v2/pushes"
	var query = []byte(`{"type": "note", "title": "Note Title", "body": "Note Body", "email": "xxxxxxxx"}`)
	req, err := http.NewRequest("POST", uri_path, bytes.NewBuffer(query))
	req.Header.Set("Access-Token", "xxxxxxx")
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
		fmt.Println(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
