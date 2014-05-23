consul-discovery
================

This Go package is a service discovery client for [Consul](http://www.consul.io).

Related: see [consul-kv](https://github.com/armon/consul-kv) for a key-value client package to Consul.

Documentation
=============

The full documentation is available on [Godoc](http://godoc.org/github.com/drnic/consul-discovery)

Usage
=====

Below is an example of using the Consul KV client:

```go
client, _ := consuldiscovery.NewClient(consuldiscovery.DefaultConfig())
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
