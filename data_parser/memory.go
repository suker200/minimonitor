package data_parser

import (
    "fmt"
    "log"
    "strconv"
    "strings"
    "io/ioutil"
    "github.com/suker200/config_parser"
    "github.com/suker200/data_report"
)

func Memory(os string, object_config map[string]map[string]interface{}, object_tag config_parser.Server, messages chan string) {
    f, err := ioutil.ReadFile("/proc/meminfo")
    
    if err != nil {
        log.Fatal(err)
    }

    str_memory_usage := ""
    lines := strings.Split(string(f), "\n")
    for _, line := range lines {
        fields := strings.SplitN(line, ":", 2)

        if strings.Compare(fields[0], "Active") == 0 {
            str_memory_usage = strings.TrimSpace(fields[1])
            str_memory_usage = strings.Replace(str_memory_usage, " kB", "", -1)
            int_memory_usage, _ := strconv.Atoi(str_memory_usage)
            int_memory_usage = int_memory_usage/1024
            str_memory_usage = strconv.Itoa(int_memory_usage) + " MB"
        }
    }

    warning_threshold := object_config["memory"]["warning"].(int)
    critical_threshold := object_config["memory"]["critical"].(int)
    data_report.Pushbullet_report(object_config["email"]["email"].(string), "Memory", str_memory_usage, warning_threshold, critical_threshold)

    memory := "memory" + "," + object_tag.Tag + "," + "type=memory_usage" + " value=" + str_memory_usage

    messages <- memory
    //fmt.Println(memory)
    //return memory
}