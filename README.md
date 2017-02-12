# Go Benchmark App 
[![Build Status](https://travis-ci.org/mrLSD/go-benchmark-app.svg?branch=master)](https://travis-ci.org/mrLSD/go-benchmark-app)   [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/mrLSD/go-benchmark-app/master/LICENSE)  [![codecov](https://codecov.io/gh/mrLSD/go-benchmark-app/branch/master/graph/badge.svg)](https://codecov.io/gh/mrLSD/go-benchmark-app)  [![Coverage Status](https://coveralls.io/repos/github/mrLSD/go-benchmark-app/badge.svg?branch=master)](https://coveralls.io/github/mrLSD/go-benchmark-app?branch=master)  [![Go Report Card](https://goreportcard.com/badge/github.com/mrLSD/go-benchmark-app)](https://goreportcard.com/report/github.com/mrLSD/go-benchmark-app)  [![GoDoc](https://godoc.org/github.com/mrLSD/go-benchmark-app?status.png)](https://godoc.org/github.com/mrLSD/go-benchmark-app)

_The efficiency and speed of application - our goal and the basic idea._

![Go Benchmark App ](http://letzgro.net/wp-content/uploads/2016/01/banners-4.png)

Application for HTTP-benchmarking via different rules and configurations.

Configurations provided via TOML-config.

Our main aims - compare different web-applications 
that we want to compare with same benchmarks tools, 
same parameters, repeated and distributed in time.
Get average results and compare it with other applications.

This will allow us to analyze and compare the performance 
bottlenecks, and to take appropriate measures.
The efficiency and speed of application - our goal and the 
basic idea.

## Benchmarks tools 
`ab`, `wrk`, `siege`

It should be installed.

## How to configure
All config at `config/main.toml`

````
# Title for describing benchmarks
title = "CMS benchmarks"
# Benchmarks version
version = "0.1"
# Delay via try in seconds
delay = 20
# How much we should try
try = 10

# ab benchmarks parametres
[ab]
concurency = 5000
keepalive = false
requests = 10000

# wrk benchmarks parametres
[wrk]
connections = 5000
duration = 10
threads = 1000

# siege benchmarks parametres
[siege]
concurrent = 100
time = 30

# Applications parametres - list
[[app]]
title = "Application Banchamrs Title"
path = "fool/path/to/app"
url = "http://localhost:5000/test"
````

## How to use
* Installed `Go 1.6+`
* Check is benchmarks tools installed:

	`$ whereis ab`
	
	`$ whereis wrk`
	
	`$ whereis siege` 
* Configure `config/main.toml` in current dirrectory as mentioned before.
* Install application: `$ go install github.com/mrlsd/go-benchmark-app`
* Command-line help: `$ go-benchmark-app -h`
```
Go Benchmark Applications v1.0.0
Options:
  -c FILE
    	load configuration from FILE (default "config/main.toml")
  -v	verbose output
```
* To change config file - run: `$ go-benchmark-app -c path/to/cfg.toml`
* Verbose output: `$ go-benchmark-app -v`

## Tips & Tricks
* When your benchmarks test failed with socket error 
(more resources unawailable), try increase `delay` options
 at config file (for example 60 sec) or/and change system 
 limits to opened files, TCP/IP configuration and other 
 system limitations.
* For flexibility running test applications you can create
recipe for that. For example set at config file:
```
[[app]]
title = "My App"
path = "apps/myapp.sh"
url = "http://localhost:5000/test"
```

and `apps/myapp.sh` file:
```
#!/bin/sh
prepare -to -run
/full/path/to/app -v param -d param -etc
```
in that way you can run `docker`, `nginx` or another
useful commands.

For one application we use one `URL`, because it's simplify
results analyze, interpretation, comparison.

#### License: MIT [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/mrLSD/go-benchmark-app/master/LICENSE)
