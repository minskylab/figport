package main

import (
	"context"

	"github.com/minskylab/figport"
	"github.com/sirupsen/logrus"
)

func main() {
	exporter, err := figport.New(context.Background())
	if err != nil {
		logrus.Panic(err.Error())
		return
	}

	if err := exporter.Start(); err != nil {
		logrus.Panic(err.Error())
	}
}
