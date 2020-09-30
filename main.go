package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"k8s/tomcat"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		DisableColors: false,
	})
	logrus.SetReportCaller(true)
	err := tomcat.CreateDeployment()
	err = tomcat.CreateService()
	err = tomcat.CreateIngress()
	if err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	fmt.Println("the task is over")
}
