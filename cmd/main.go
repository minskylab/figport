package main

import (
	"context"
	"flag"

	"github.com/minskylab/figport"
	"github.com/sirupsen/logrus"
)

func main() {
	debugMode := flag.Bool("debug", false, "activate or not: the debug log level")
	flag.Parse()

	ctx := context.Background()
	exporter, err := figport.New(ctx)
	if err != nil {
		logrus.Panic(err.Error())
		return
	}

	if err := exporter.Start(*debugMode); err != nil {
		logrus.Panic(err.Error())
	}
}
