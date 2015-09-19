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
    access_log = "access_log" + "," + object_tag.Tag + "," + "type=access_log" + " value=" + access_log

    messages <- access_log
    //fmt.Println(access_log)
    //return access_log
}