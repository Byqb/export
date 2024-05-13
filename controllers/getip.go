package controllers

import (
	"log"
	"net"
	
)

func GetIp() string { // get the ip address to be able to use the website on other computers

	ip, err := net.Dial("udp", "8.8.8.8:80")

	if err != nil {
		log.Fatal(err)
	}
	defer ip.Close()

	address := ip.LocalAddr().(*net.UDPAddr)

	return address.IP.String()
}
