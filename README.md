consul-discovery
================

This Go package is a service discovery client for [Consul](http://www.consul.io).

Related: see [consul-kv](https://github.com/armon/consul-kv) for a key-value client package to Consul.

Documentation
=============

The full documentation is available on [Godoc](http://godoc.org/github.com/drnic/consul-discovery)

Usage
=====

Below is an example of using the Consul Discovery client:

```go
client, _ := consuldiscovery.NewClient(consuldiscovery.DefaultConfig())
criticalChecks, _ := client.HealthByState("critical")

services, _ := client.CatalogServices()
serviceNodes, _ := client.CatalogServiceByName(services[0].Name)
```

Also see examples folder for runnable examples.

First, run the `consul agent` command:

```
consul agent -data-dir /tmp/consul -server -bootstrap -config-dir testconfig
```

Now the examples can be run on this example local consul node:

```
$ go run examples/catalog.go
Services: consuldiscovery.CatalogServices{consuldiscovery.CatalogService{Name:"consul"...}

consul: consuldiscovery.CatalogServiceByName{consuldiscovery.CatalogServiceNode{Node:"drnic.local"...}
simple_service: consuldiscovery.CatalogServiceByName{consuldiscovery.CatalogServiceNode{Node:"drnic.local"...}
```

Development
===========

To run the tests, first run a consul server with API port 8500:

```
consul agent -data-dir /tmp/consul -server -bootstrap -config-dir testconfig
```

The run tests:

```
go test
```
