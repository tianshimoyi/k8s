package tomcat

import (
	"context"
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
	//"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

const (
	MOUNTPATH="/usr/local/tomcat/webapps/ROOT"
	DIRNAME="/tmp/tomcat"
	RAPLICAS=3
	ContainerPORT=8080
	TaintKEY="isadmin"
	TaintValue="true"
	TaintTIME=60*60*24*360
	NameSpace=apiv1.NamespaceDefault
)

var(
	MOUNTTYPE=apiv1.HostPathDirectoryOrCreate
	Clientset=getContext()
)

func CreateDeployment()error{


	dplyname:="test-client-go-tomcat-dply"
	deployment,err:=Clientset.AppsV1().Deployments(NameSpace).Get(context.TODO(),dplyname,metav1.GetOptions{})
	if err!=nil{
		fmt.Println("deplyment ",dplyname," is not exits")
	}else {
		Clientset.AppsV1().Deployments(NameSpace).Delete(context.TODO(),dplyname,metav1.DeleteOptions{})
		fmt.Println("please wait 10 secends or more time ",dplyname," is deleting")
		time.Sleep(time.Second*10)
	}
	deploymentClient:=Clientset.AppsV1().Deployments(NameSpace)
	deployment=&appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      dplyname,
			Namespace: NameSpace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: func(x int32) *int32 { return &x }(RAPLICAS),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":     "my-tomcat-client-go",
					"version": "v1",
				},
			},
			Template:apiv1.PodTemplateSpec{
				ObjectMeta:metav1.ObjectMeta{
					Labels: map[string]string{
						"app":     "my-tomcat-client-go",
						"version": "v1",
					},
				},
				Spec:apiv1.PodSpec{
					Containers:[]apiv1.Container{
						{
							Name:"my-tomcat-v1",
							Image:"my-tomcat:v1",
							ImagePullPolicy:apiv1.PullIfNotPresent,
							VolumeMounts:[]apiv1.VolumeMount{
								{
									MountPath:MOUNTPATH,
									Name: "test1",
								},
							},
							Ports:[]apiv1.ContainerPort{
								{
									Name:"http",
									ContainerPort:ContainerPORT,
								},
							},
						},
					},
					Volumes:[]apiv1.Volume{
						{
							Name:"test1",
							VolumeSource:apiv1.VolumeSource{
								HostPath:& apiv1.HostPathVolumeSource{
									Path:DIRNAME,
									Type:&MOUNTTYPE,
								},
							},
						},
					},
					Tolerations:[]apiv1.Toleration{
						{
							Key:TaintKEY,
							Operator:apiv1.TolerationOpEqual,
							Effect: apiv1.TaintEffectNoExecute,
							Value:TaintValue,
							TolerationSeconds:func(x int64)*int64{return &x}(TaintTIME),
						},
					},
				},
			},
		},

	}
	fmt.Println("begin to create deplyment ",dplyname)
	_,err=deploymentClient.Create(context.TODO(),deployment,metav1.CreateOptions{})
	if err!=nil{
		fmt.Println("create deployment err: ",err)
		return err
	}
	fmt.Println("deployment ",dplyname," created")
	return nil
}
