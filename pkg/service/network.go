package service

import (
	"fmt"
	"os/exec"
)

// NetworkService
type NetworkService interface {
	Scan() string
	GetIp(mac string) string
}

type networkService struct {
	t string
}

func NewNetworkService() NetworkService {
	return &networkService{
		t: "network",
	}
}

// Scan returns a Go array of device for the given interface.
func (*networkService) Scan() string {
	fmt.Print("Scan function\n")
	// apr-scan -I [interface] [gateway]/24
	app := "arp"
	arg0 := "-a"

	cmd := exec.Command(app, arg0)
	// cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	// Print the output
	fmt.Println(string(stdout))
	test := GetInterface()
	return test
}

// GetIp returns a Go string literal with the ip for the given mac.
func (*networkService) GetIp(mac string) string {
	fmt.Print("GetIp function\n")
	// apr-scan -I [interface] [gateway]/24
	app := "arp"

	arg0 := "-a"
	arg1 := "| grep '192'"
	// arg2 := "\n\tfrom"
	// arg3 := "golang"

	cmd := exec.Command(app, arg0, arg1)
	// cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	// Print the output
	fmt.Println(string(stdout))
	return ""
}

// GetInterace returns a Go string literal with the name of the computer default
// interface.
func GetInterface() string {
	/*
		Nash script

		host=google.com

		# get the ip of that host (works with dns and /etc/hosts. In case we get
		# multiple IP addresses, we just want one of them
		host_ip=$(getent ahosts "$host" | awk '{print $1; exit}')

		# only list the interface used to reach a specific host/IP. We only want the part
		# between dev and the remainder (use grep for that)
		ip route get "$host_ip" | grep -Po '(?<=(dev ))(\S+)'
	*/
	return "wlan0"
}
