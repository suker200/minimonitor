package data_parser

import (
    "fmt"
    "log"
    "strings"
    "io/ioutil"
    "github.com/suker200/config_parser"
    //"github.com/suker200/data_report"
)

func Memory(os string, object_config map[string]map[string]interface{}, object_tag config_parser.Server, messages chan string) {
    f, err := ioutil.ReadFile("/proc/meminfo")
    
    if err != nil {
        log.Fatal(err)
    }

    memory_usage := ""
    memory_limit := ""
    lines := strings.Split(string(f), "\n")
    for _, line := range lines {
        fields := strings.SplitN(line, ":", 2)

        if strings.Compare(fields[0], "MemTotal") == 0 {
            memory_limit = strings.TrimSpace(fields[1])
            memory_limit = strings.Replace(memory_limit, " kB", "", -1)
        }

        if strings.Compare(fields[0], "Active") == 0 {
            memory_usage = strings.TrimSpace(fields[1])
            memory_usage = strings.Replace(memory_usage, " kB", "", -1)
        }
    }

    warning_threshold := object_config["memory"]["warning"].(int)
    critical_threshold := object_config["memory"]["critical"].(int)
    data_report.Pushbullet_report(object_config["email"]["email"].(string), "Memory", memory_usage, warning_threshold, critical_threshold)

    memory := "memory" + "," + "object_tag.Tag" + "," + "type=memory_usage" + " value=" + memory_usage + "\n"
    memory += "memory" + "," + "object_tag.Tag" + "," + "type=memory_limit" + " value=" + memory_limit + "\n"

    messages <- memory
    //fmt.Println(memory)
    //return memory
}