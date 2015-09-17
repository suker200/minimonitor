package data_parser

import (
	"fmt"
	"github.com/suker200/config_parser"
	// "github.com/suker200/data_report"
	"log"
	"os/exec"
)

// var Ubuntu = []byte("/usr/sbin/ss -a | grep -v '*\\|127.0.0.1\\|Address' | awk '{print $NF}' |  cut -d: -f1 | sed -e '/^$/d' | sort -r | uniq -c  | sort -n | tail -n 10   | awk '{print $2,$1}' | tr ' ' ':' | tr '\\n' ','  | sed 's/:/-:-/g' | sed 's/,/-,-/g' | sed 's/,-$//g'")

func Ccu(os string, object_config map[string]map[string]interface{}, object_tag config_parser.Server, messages chan string) {
	var data = make(map[string][]byte)
	var return_value string
	data["centos"] = []byte("/usr/sbin/ss -a | grep -v '*\\|127.0.0.1\\|Address' | sed -e '/^$/d'  | wc -l")
	data["ubuntu"] = []byte("/usr/sbin/ss -a | grep -v '*\\|127.0.0.1\\|Address' | sed -e '/^$/d'  | wc -l")
	data["fedora"] = []byte("/usr/sbin/ss -a | grep -v '*\\|127.0.0.1\\|Address' | sed -e '/^$/d'  | wc -l")

	cmd := string(data[os][:])
	out, err := exec.Command("bash", "-c", cmd).Output()
	result := string(out[:])
	if err == nil {
		fmt.Println(result)
		// data_report.Influxdb_report("ccu", object_tag.Tag, result)
		return_value = "ccu" + "," + object_tag.Tag + " value=" + result
	} else {
		log.Fatal(err)
	}
	// return return_value
	messages <- return_value
}
