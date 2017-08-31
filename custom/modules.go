package main

import (
	_ "github.com/adamryman/logspout-syslog-json/adapters/json"
	_ "github.com/adamryman/logspout-syslog-json/adapters/syslog"
	_ "github.com/gliderlabs/logspout/transports/tcp"
)
