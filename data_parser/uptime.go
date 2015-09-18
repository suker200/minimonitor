package data_parser

import (
    "fmt"
    "log"
    "strings"
    "io/ioutil"
    "github.com/suker200/config_parser"
    "github.com/suker200/data_report"
)

func Uptime(os string, object_config map[string]map[string]interface{}, object_tag config_parser.Server, messages chan string) {
    f, err := ioutil.ReadFile("/proc/uptime")
    
    if err != nil {
        log.Fatal(err)
    }

    content := strings.TrimSpace(string(f))
    fields := strings.Fields(content)
    uptime_total := fields[0]
    uptime_idle := fields[1]

    warning_threshold := object_config["uptime"]["warning"].(int)
    critical_threshold := object_config["uptime"]["critical"].(int)
    data_report.Pushbullet_report(object_config["email"]["email"].(string), "Uptime", uptime_total, warning_threshold, critical_threshold)

    uptime := "uptime" + "," + object_tag.Tag + "," + "type=uptime_total" + " value=" + uptime_total
    uptime += "uptime" + "," + object_tag.Tag + "," + "type=uptime_idle" + " value=" + uptime_idle

    messages <- uptime
    //fmt.Println(uptime)
    //return uptime
}