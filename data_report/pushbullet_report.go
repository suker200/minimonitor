package data_report

import (
	"bytes"
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func Pushbullet_report(email string, title_message string, result string, warning int, critical int) {
	// convert result value to int
	var notification string
	var check bool
	var re = regexp.MustCompile(`[\n%]`)
	// result_int, err := strconv.Atoi(strings.Replace(result, "\n", "", -1))
	result_int, err := strconv.Atoi(re.ReplaceAllString(result, ""))
	if result_int > critical && err == nil {
		// critical case, send notification
		notification = "[CRITICAL] " + title_message + " " + strings.Replace(result, "\n", "", -1) + " greater than " + strconv.Itoa(critical)
		// data_report.Pushbullet_report(object_config["email"]["email"].(string), "CCU", notification)
		check = true
	} else if result_int > warning && err == nil {
		// warning case, send notification
		notification = "[WARNING] " + title_message + " " + strings.Replace(result, "\n", "", -1) + " greater than " + strconv.Itoa(warning)
		// data_report.Pushbullet_report(object_config["email"]["email"].(string), "CCU", notification)
		check = true
	}

	if check {
		list_email := strings.Split(email, ",")
		for _, reciever := range list_email {
			uri_path := "https://api.pushbullet.com/v2/pushes"
			email_reciever := reciever
			message := "{'type': 'note', 'title':'" + title_message + "', 'body':'" + notification + "', 'email':'" + email_reciever + "'}"
			message_format := strings.Replace(message, "'", "\"", -1)
			req, err := http.NewRequest("POST", uri_path, bytes.NewBuffer([]byte(message_format)))
			req.Header.Set("Access-Token", "xxx")
			req.Header.Set("Content-Type", "application/json")
			if err != nil {
				panic(err)
				fmt.Println(err)
			}
			client := &http.Client{}
			resp, err := client.Do(req)
			response, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(response))
		}
	}
}
