// package main
package data_parser

import (
	"fmt"
	"github.com/suker200/config_parser"
	"github.com/suker200/data_report"
	// "io/ioutil"
	"os/exec"
	"regexp"
	"strings"
)

func Disk(os string, object_config map[string]map[string]interface{}, object_tag config_parser.Server, messages chan string) {
	// func main() {
	warning_threshold := object_config["disk"]["warning"].(int)
	critical_threshold := object_config["disk"]["critical"].(int)
	var disk string
	var data = make(map[string]string)
	var re = regexp.MustCompile(`[\n%]`)

	data["centos"] = "df -h | grep '/mapper/' | awk '{print $5,$6}'"
	data["fedora"] = "df -h | grep '/mapper/' | awk '{print $5,$6}'"

	disk_cmd := exec.Command("sh", "-c", data["fedora"])
	disk_out, err := disk_cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	disk_data := strings.TrimSpace(string(disk_out[:]))
	list_disk := strings.Split(disk_data, "\n")
	for _, disk_element := range list_disk {
		disk_info := strings.Split(disk_element, " ")
		fmt.Println(disk_info[0])
		fmt.Println(disk_info[1])
		disk += "disk," + object_tag.Tag + "," + "type=" + disk_info[1] + " value=" + re.ReplaceAllString(disk_info[0], "") + "\n"
		data_report.Pushbullet_report(object_config["email"]["email"].(string), "disk "+disk_info[1], disk_info[0], warning_threshold, critical_threshold)
	}

	// fmt.Println(disk)
	// fields := strings.Fields(content)

	// loadavg1min := fields[0]
	// loadavg5min := fields[1]
	// loadavg15min := fields[2]

	// warning_threshold := object_config["loadavg"]["warning"].(int)
	// critical_threshold := object_config["loadavg"]["critical"].(int)
	// data_report.Pushbullet_report(object_config["email"]["email"].(string), "LoadAvg1Min", loadavg15min, warning_threshold, critical_threshold)

	// // warning_threshold := object_config["loadavg5min"]["warning"].(int)
	// // critical_threshold := object_config["loadavg5min"]["critical"].(int)
	// // data_report.Pushbullet_report(object_config["email"]["email"].(string), "LoadAvg5Min", loadavg15min, warning_threshold, critical_threshold)

	// // warning_threshold := object_config["loadavg15min"]["warning"].(int)
	// // critical_threshold := object_config["loadavg15min"]["critical"].(int)
	// // data_report.Pushbullet_report(object_config["email"]["email"].(string), "LoadAvg15Min", loadavg15min, warning_threshold, critical_threshold)

	// loadavg := "loadavg" + "," + object_tag.Tag + "," + "type=loadavg1min" + " value=" + loadavg1min + "\n"
	// loadavg += "loadavg" + "," + object_tag.Tag + "," + "type=loadavg5min" + " value=" + loadavg5min + "\n"
	// loadavg += "loadavg" + "," + object_tag.Tag + "," + "type=loadavg15min" + " value=" + loadavg15min + "\n"

	messages <- disk
	// fmt.Println(content)
	//return loadavg
}
