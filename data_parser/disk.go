package data_parser

import (
	"github.com/suker200/minimonitor/config_parser"
	"github.com/suker200/minimonitor/data_report"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func Disk(os string, function string, object_config map[string]map[string]interface{}, object_tag config_parser.Server, messages chan string) {
	var disk string
	var data = make(map[string]string)
	var re = regexp.MustCompile(`[\n%a-zA-Z ]`)

	data["centos"] = "df -h | grep '/mapper/' | grep -v 'resolv.conf\\|hostname\\|hosts' | awk '{print $5,$6}'"
	data["fedora"] = "df -h | grep '/mapper/' | grep -v 'resolv.conf\\|hostname\\|hosts' | awk '{print $5,$6}'"

	disk_cmd := exec.Command("sh", "-c", data[os])
	disk_out, err := disk_cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	disk_data := strings.TrimSpace(string(disk_out[:]))
	list_disk := strings.Split(disk_data, "\n")
	for _, disk_element := range list_disk {
		disk_info := strings.Split(disk_element, " ")
		disk += "disk," + object_tag.Tag + "," + "type=" + disk_info[1] + " value=" + re.ReplaceAllString(disk_info[0], "") + "\n"
		data_report.Pushbullet_report(function, object_config, "disk "+disk_info[1], disk_info[0])
	}

	messages <- disk
}
