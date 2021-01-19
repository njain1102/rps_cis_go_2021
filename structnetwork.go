package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Date struct{
	day int
	month int
	year int
}

type Config struct{
	subnet string
	gateway string
}
type ConfigFrom struct{
	network string
}

type IPAM struct{
	driver string
	options string
    config [2]Config

}
type Network struct{
	name string
	id string
	created Date
	scope string
	driver string
	enableIPv6 bool
    ipam IPAM
	internal bool
	attachable bool
	ingress bool
    configFrom ConfigFrom
	configOnly bool
	containers string
	options string
	labels string

}


func main() {
    var config [2]Config
    for i:=0;i<len(config);i++ {
		config[i] = Config{"172.20.0.0/16"+strconv.Itoa(rand.Int()),
			"172.20.0.1"+strconv.Itoa(rand.Int())}
	}

	ipam:=IPAM{"default","",config}
	configFrom:=ConfigFrom{""}
    network:=Network{"virtusa_network","fa3d377e197ab5eec3e44951b1f3de60b906c1924d06ab7fe20349874dd7707e",
    	Date{11,11,2020},"local","bridge",
    	false,ipam,false,false,false,
    	configFrom,false,"","",""}

    fmt.Println(network)
    fmt.Printf("Netowrk Details %+v",network)
    fmt.Println("\nEnableIPV6 Status",network.enableIPv6)
}
