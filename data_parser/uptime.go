package data_parser

import (
    "fmt"
    "strings"
    "io/ioutil"
    "github.com/suker200/config_parser"
    //"github.com/suker200/data_report"
)

func Uptime(os string, object_config map[string]map[string]interface{}, object_tag config_parser.Server, messages chan string) {
    f, err := ioutil.ReadFile("/proc/uptime")
    
    if err != nil {
        panic(err)
    }

    content := strings.TrimSpace(string(f))
    fields := strings.Fields(content)
    uptime_total := fields[0]
    uptime_idle := fields[1]

    warning_threshold := object_config["uptime"]["warning"].(int)
    critical_threshold := object_config["uptime"]["critical"].(int)
    data_report.Pushbullet_report(object_config["email"]["email"].(string), "Uptime", uptime_total, warning_threshold, critical_threshold)

    uptime := "uptime_total" + "," + "object_tag.Tag" + "," + "type=uptime" + " value=" + uptime_total
    uptime += "\nuptime_idle" + "," + "object_tag.Tag" + "," + "type=uptime" + " value=" + uptime_idle

    messages <- uptime
    //fmt.Println(uptime)
    //return uptime
}