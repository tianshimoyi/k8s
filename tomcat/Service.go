package tomcat

import (
	"context"
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

const(
	ServerNAME="client-go-my-tomcat"
	ServerPort=8081
)

func CreateService()error{
	serverClient:=Clientset.CoreV1().Services(NameSpace)
	_,err:=serverClient.Get(context.TODO(),ServerNAME,metav1.GetOptions{})
	if err!=nil{
		fmt.Println("deplyment ",ServerNAME," is not exits")
	}else {
		serverClient.Delete(context.TODO(),ServerNAME,metav1.DeleteOptions{})
		fmt.Println("please wait 10 secends or more time ",ServerNAME," is deleting")
		time.Sleep(time.Second*10)
	}

	service:=&apiv1.Service{
		TypeMeta:metav1.TypeMeta{
			APIVersion:"v1",
			Kind:"Service",
		},
		ObjectMeta:metav1.ObjectMeta{
			Name:ServerNAME,
			Namespace:NameSpace,
		},
		Spec:apiv1.ServiceSpec{
			Selector: map[string]string{
				"app":     "my-tomcat-client-go",
				"version": "v1",
			},
			Ports:[]apiv1.ServicePort{
				{
					Name:"http",
					Port:ServerPort,
					TargetPort:intstr.IntOrString{
						Type:0,
						IntVal:ContainerPORT,
					},
				},
			},
		},
	}

	fmt.Println("begin to create service ",ServerNAME)
	_,err=serverClient.Create(context.TODO(),service,metav1.CreateOptions{})
	if err!=nil{
		fmt.Println("create service err: ",err)
		return err
	}
	fmt.Println("deployment ",ServerNAME," created")
	return nil
}
