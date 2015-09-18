package data_parser

import (
    "fmt"
    "strings"
    "io/ioutil"
    "github.com/suker200/config_parser"
    //"github.com/suker200/data_report"
)

func LoadAvg(os string, object_config map[string]map[string]interface{}, object_tag config_parser.Server, messages chan string) {
    f, err := ioutil.ReadFile("/proc/loadavg")

    if err != nil {
        panic(err)
    }

    content := strings.TrimSpace(string(f))
    fields := strings.Fields(content)

    loadavg1min  := fields[0]
    loadavg5min  := fields[1]
    loadavg15min := fields[2]

    warning_threshold := object_config["loadavg1min"]["warning"].(int)
    critical_threshold := object_config["loadavg1min"]["critical"].(int)
    data_report.Pushbullet_report(object_config["email"]["email"].(string), "LoadAvg1Min", loadavg15min, warning_threshold, critical_threshold)

    warning_threshold := object_config["loadavg5min"]["warning"].(int)
    critical_threshold := object_config["loadavg5min"]["critical"].(int)
    data_report.Pushbullet_report(object_config["email"]["email"].(string), "LoadAvg5Min", loadavg15min, warning_threshold, critical_threshold)

    warning_threshold := object_config["loadavg15min"]["warning"].(int)
    critical_threshold := object_config["loadavg15min"]["critical"].(int)
    data_report.Pushbullet_report(object_config["email"]["email"].(string), "LoadAvg15Min", loadavg15min, warning_threshold, critical_threshold)

    loadavg := "loadavg1min" + "," + "object_tag.Tag" + "," + "type=loadavg" + " value=" + loadavg1min
    loadavg += "\nloadavg5min" + "," + "object_tag.Tag" + "," + "type=loadavg" + " value=" + loadavg5min
    loadavg += "\nloadavg15min" + "," + "object_tag.Tag" + "," + "type=loadavg" + " value=" + loadavg15min

    messages <- loadavg
    //fmt.Println(loadavg)
    //return loadavg
}