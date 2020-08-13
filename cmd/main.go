package main

import (
	"context"
	"flag"

	"github.com/minskylab/figport"
	"github.com/sirupsen/logrus"
)

func main() {
	debugMode := flag.Bool("debug", false, "activate or not the debug log level")
	flag.Parse()

	exporter, err := figport.New(context.Background())
	if err != nil {
		logrus.Panic(err.Error())
		return
	}

	if err := exporter.Start(*debugMode); err != nil {
		logrus.Panic(err.Error())
	}
}
