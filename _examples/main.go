package main

import (
	"fmt"
	"log"

	"github.com/stsy/remoteip"
)

func main() {
	ip, err := remoteip.Get("http://www.howtofindmyipaddress.com/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ip)
}

/*
Working URLs:
 http://ipchicken.com
 http://ipinfo.info/index.php
 https://api.ipify.org/
 http://www.howtofindmyipaddress.com/
*/
