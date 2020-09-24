package tomcat

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	exv1beta "k8s.io/api/networking/v1"
	"time"
)
const (
	IngressNAME="client-go-ingress-my-tomcat"
	Host="www.mytomcat.com"
)

var (
	PathType=exv1beta.PathTypePrefix
)

func CreateIngress()error{
	ingressClinet:=Clientset.NetworkingV1().Ingresses(NameSpace)
	_,err:=ingressClinet.Get(context.TODO(),IngressNAME,metav1.GetOptions{})
	if err!=nil{
		fmt.Println("ingress ",IngressNAME," is not exits")
	}else{
		ingressClinet.Delete(context.TODO(),IngressNAME,metav1.DeleteOptions{})
		fmt.Println("please wait 10 secends or more time ",IngressNAME," is deleting")
		time.Sleep(time.Second*10)
	}

	ingress:=&exv1beta.Ingress{
		TypeMeta:metav1.TypeMeta{
			APIVersion:"networking.k8s.io/v1",
			Kind:"Ingress",
		},
		ObjectMeta:metav1.ObjectMeta{
			Name:IngressNAME,
			Namespace:NameSpace,
			Annotations: map[string]string{
				"nginx.ingress.kubernetes.io/rewrite-target":"/",
			},
		},
		Spec:exv1beta.IngressSpec{
			Rules:[]exv1beta.IngressRule{
				{
					Host:Host,
					IngressRuleValue:exv1beta.IngressRuleValue{
						HTTP:&exv1beta.HTTPIngressRuleValue{
							Paths:[]exv1beta.HTTPIngressPath{
								{
									PathType:&PathType,
									Path: "/",
									Backend:exv1beta.IngressBackend{
										Service:&exv1beta.IngressServiceBackend{
											Name:ServerNAME,
											Port:exv1beta.ServiceBackendPort{
												Number:ServerPort,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println("ingress ",IngressNAME," is creating")
	_,err=ingressClinet.Create(context.TODO(),ingress,metav1.CreateOptions{})
	if err!=nil{
		fmt.Println("creating ingress err: ",err)
		return err
	}
	fmt.Println("ingress ",IngressNAME," create!" )
	return nil
}
