package main

import (
	"fmt"
	"k8s/tomcat"
)

func init(){
	//test.CreateDeployment()
	//test.CreateServer()
	//test.GetPodMessage()
	//test.CreatePod()
	err:=tomcat.CreateDeployment()
	err=tomcat.CreateService()
	err=tomcat.CreateIngress()
	if err!=nil{
		panic(err.Error())
	}
}

func main(){
	//fmt.Println(" deployment create allready!")
	fmt.Println("the task is over")
}