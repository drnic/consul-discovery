package main

import (
	"fmt"

	"github.com/drnic/consul-discovery"
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	client, _ := consuldiscovery.NewClient(consuldiscovery.DefaultConfig())

	services, err := client.CatalogServices()
	panicIf(err)
	fmt.Printf("Services: %#v\n\n", services)

	for _, service := range services {
		serviceNodes, err := client.CatalogServiceByName(service.Name)
		panicIf(err)
		fmt.Printf("%s: %#v\n", service.Name, serviceNodes)
	}
}
