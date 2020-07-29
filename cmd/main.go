package main

import (
	"context"

	"github.com/minskylab/figport"
	"github.com/sirupsen/logrus"
	// "github.com/minskylab/figport"
)

func main() {
	figport, err := figport.NewDefault(context.Background())
	if err != nil {
		logrus.Panic(err.Error())
	}

	if err := figport.Start(); err != nil {
		logrus.Panic(err.Error())
	}
}
