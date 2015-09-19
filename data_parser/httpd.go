package main

import (
    //"fmt"
    "log"
    "os/exec"
    //"github.com/suker200/config_parser"
    //"github.com/suker200/data_report"
)

func Httpd(os string, object_config map[string]map[string]interface{}, object_tag config_parser.Server, messages chan string) {
    
    out, err := exec.Command("bash", "-c", "cat `whereis apache2`/logs/access_log | wc -l").Output()

    if err != nil {
        log.Fatal(err)
    }
    
    access_log := string(out[:])
    httpd = "httpd" + "," + object_tag.Tag + "," + "type=httpd" + " value=" + access_log

    messages <- httpd
    //fmt.Println(httpd)
    //return httpd
}