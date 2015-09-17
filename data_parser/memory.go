package data_parser

import (
	"fmt"
	"github.com/suker200/config_parser"
	// "github.com/suker200/data_report"
	"log"
	"os/exec"
	"strconv"
)

func Mem(os string, object_config map[string]map[string]interface{}, object_tag config_parser.Server, messages chan string) {
	var data = make(map[string][]byte)
	var return_value string
	data["centos"] = []byte("cat /proc/meminfo  | grep 'MemAvailable:' | awk '{print $2}' | tr '\n' ' ' | sed 's/ //'g")
	data["ubuntu"] = []byte("cat /proc/meminfo  | grep 'MemAvailable:' | awk '{print $2}' | tr '\n' ' ' | sed 's/ //'g")
	data["fedora"] = []byte("cat /proc/meminfo  | grep 'MemAvailable:' | awk '{print $2}' | tr '\n' ' ' | sed 's/ //'g")

	cmd := string(data[os][:])
	out, err := exec.Command("bash", "-c", cmd).Output()
	string_convert := string(out[:])
	string_to_int, err := strconv.Atoi(string_convert)
	if err == nil {
		out := string_to_int / 1028
		result := strconv.Itoa(out)
		fmt.Println(result)
		// data_report.Influxdb_report("memory", object_tag.Tag, result)
		return_value = "memory" + "," + object_tag.Tag + " value=" + result
		// return return_value
	} else {
		log.Fatal(err)
	}
	// return return_value
	messages <- return_value
}
