package data_parser

import (
    "fmt"
    "strings"
    "io/ioutil"
    "github.com/suker200/config_parser"
    //"github.com/suker200/data_report"
)

func NetIO(os string, object_config map[string]map[string]interface{}, object_tag config_parser.Server, messages chan string) {
    f, err := ioutil.ReadFile("/proc/net/dev")
    
    if err != nil {
        panic(err)
    }

    netio := ""
    lines := strings.Split(string(f), "\n")
    for _, line := range lines {
        if strings.Contains(line, ":") {
            content := strings.TrimSpace(line)
            fields := strings.Fields(content)

            dev := strings.Replace(fields[0], ":" , "", -1)
            in  := fields[1]
            out := fields[9]
            netio += "\nin" + "," + "object_tag.Tag" + "," + "type=netio" + " interface=" + dev + " value=" + in
            netio += "\nout" + "," + "object_tag.Tag" + "," + "type=netio" + " interface=" + dev + " value=" + out
        }
    }

    netio = strings.TrimPrefix(netio, "\n")

    messages <- netio
    //fmt.Println(netio)
    //return netio
}