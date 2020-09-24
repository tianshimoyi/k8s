package tomcat

import (
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func getHomeDir()string{
	if home:=os.Getenv("HOME");home != ""{
		return home
	}
	return os.Getenv("USERPROFILE")
}

func getContext() (*kubernetes.Clientset){
	var kubeconfig *string
	if home := getHomeDir(); home != ""{
		kubeconfig=flag.String("kubeconfig",filepath.Join(home,".kube","config"),"input kubeconfig addr")
	}else{
		kubeconfig=flag.String("kubeconfig","","input kubeconfig addr")
	}
	flag.Parse()
	config,err:=clientcmd.BuildConfigFromFlags("",*kubeconfig)
	if err!=nil{
		fmt.Println("解析配置文件出错",err)
		return nil
	}
	clientset:=kubernetes.NewForConfigOrDie(config)
	return clientset
}
