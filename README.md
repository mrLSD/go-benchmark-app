# Go Benchmark App 
[![Build Status](https://travis-ci.org/mrLSD/go-benchmark-app.svg?branch=master)](https://travis-ci.org/mrLSD/go-benchmark-app)   [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/mrLSD/go-benchmark-app/master/LICENSE)  [![Coverage Status](https://coveralls.io/repos/github/mrLSD/go-benchmark-app/badge.svg?branch=master)](https://coveralls.io/github/mrLSD/go-benchmark-app?branch=master)

_The efficiency and speed of application - our goal and the basic idea._

[![Build Status](http://letzgro.net/wp-content/uploads/2016/01/banners-4.png)]

Application for HTTP-benchmarking via different rules and configurations.

Configurations provided via TOML-config.

Our main aims - compare different web-applications 
that we want to compare with same benchmarks tools, 
same parametres, repeated and distributed in time.
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
````

## How to use
* Check is benchmarks tools installed:

	`$ whereis ab`
	
	`$ whereis wrk`
	
	`$ whereis siege` 
* Configure `config/main.toml` as mentioned before.
* Build application: `$ make build`
* Run application: `$ make run` or after build `$ ./go-benchmark-app`

#### License: MIT [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/mrLSD/go-benchmark-app/master/LICENSE)
