package api

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	_ "nuc.lliu.ca/gitea/app/scale_maker/docs" // doc import for swagger
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/config"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/k8s"
	ctrl "sigs.k8s.io/controller-runtime"
)

// type k8sTemplates struct {
// 	cpuLoadTestPod *yamlutil.YAMLOrJSONDecoder
// }

var kc k8s.KClient

// var kt k8sTemplates

func newK8SClient() {
	var err error
	// config, err := rest.InClusterConfig()
	// if err != nil {
	// 	return nil, err
	// }

	// return kubernetes.NewForConfig(config)
	kc.Ctx = context.Background()
	config := ctrl.GetConfigOrDie()
	kc.ClientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal("Error: unable to create normal Kubernetes clientSet.")
	}
	kc.DynamicClient = dynamic.NewForConfigOrDie(config)
}

// func loadDefaultK8STemplates() {
// 	//var err error
// 	cpuLoadTestPodTemplate := "./templates/cpu_load_test_pod.yaml"
// 	cpuLoadTestPodFile, err := ioutil.ReadFile(cpuLoadTestPodTemplate)
// 	if err != nil {
// 		log.Fatal("Error: can not load cpuLoadTestPodTemplate with error " + err.Error())
// 	}
// 	kt.cpuLoadTestPod = yamlutil.NewYAMLOrJSONDecoder(bytes.NewReader(cpuLoadTestPodFile), 100)
// }

// InitialSetup doc
func InitialSetup() *fiber.App {
	config.NewConfig()
	conf := config.AppConfig
	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: conf.Prod, // go run app.go -prod
	})
	//newK8SClient()
	setupRoutesandMiddleware(app, false)

	// Return the configured app
	return app
}
