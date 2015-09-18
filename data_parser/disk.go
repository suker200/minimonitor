package data_parser

import (
	"github.com/suker200/config_parser"
	"github.com/suker200/data_report"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func Disk(os string, object_config map[string]map[string]interface{}, object_tag config_parser.Server, messages chan string) {
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
		log.Fatal(err)
	}

	disk_data := strings.TrimSpace(string(disk_out[:]))
	list_disk := strings.Split(disk_data, "\n")
	for _, disk_element := range list_disk {
		disk_info := strings.Split(disk_element, " ")
		disk += "disk," + object_tag.Tag + "," + "type=" + disk_info[1] + " value=" + re.ReplaceAllString(disk_info[0], "") + "\n"
		data_report.Pushbullet_report(object_config["email"]["email"].(string), "disk "+disk_info[1], disk_info[0], warning_threshold, critical_threshold)
	}

	messages <- disk
}
