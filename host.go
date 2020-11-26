package utils

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"

	"github.com/wuyingsong/utils"
)

const (
	uuid_file = "/sys/devices/virtual/dmi/id/product_uuid"
)

func getHostname() string {
	if hostname, err := os.Hostname(); err == nil {
		return hostname
	}
	return "unknown"
}

func GetHostID() string {
	var id string
	if data, err := ioutil.ReadFile(uuid_file); err == nil {
		if len(data) > 0 {
			dataStr := strings.TrimSpace(string(data))
			id = utils.Md5(dataStr)
		}
	} else {
		id = utils.Md5(getHostname())
	}
	if len(id) > 32 {
		return id[:32]
	}
	return id
}

func GetEngineID(port int) string {
	host_id := GetHostID()
	engine_id := utils.Md5(fmt.Sprintf("%s:%s", host_id, port))
	return engine_id
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "unknown"
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}
	return "unknown"
}
