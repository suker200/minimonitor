package data_parser

import (
    "fmt"
    "log"
    "strings"
    "io/ioutil"
    "github.com/suker200/config_parser"
    "github.com/suker200/data_report"
)

func LoadAvg(os string, object_config map[string]map[string]interface{}, object_tag config_parser.Server, messages chan string) {
    f, err := ioutil.ReadFile("/proc/loadavg")

    if err != nil {
        log.Fatal(err)
    }

    content := strings.TrimSpace(string(f))
    fields := strings.Fields(content)

    loadavg1min  := fields[0]
    loadavg5min  := fields[1]
    loadavg15min := fields[2]

    warning_threshold := object_config["loadavg"]["warning"].(int)
    critical_threshold := object_config["loadavg"]["critical"].(int)
    data_report.Pushbullet_report(object_config["email"]["email"].(string), "LoadAvg", loadavg15min, warning_threshold, critical_threshold)

    loadavg := "loadavg" + "," + object_tag.Tag + "," + "type=loadavg1min" + " value=" + loadavg1min + "\n"
    loadavg += "loadavg" + "," + object_tag.Tag + "," + "type=loadavg5min" + " value=" + loadavg5min + "\n"
    loadavg += "loadavg" + "," + object_tag.Tag + "," + "type=loadavg15min" + " value=" + loadavg15min + "\n"

    messages <- loadavg
    //fmt.Println(loadavg)
    //return loadavg
}