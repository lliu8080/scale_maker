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

var kc k8s.KClient

func newK8SClient() {
	var err error
	kc.Ctx = context.Background()
	config := ctrl.GetConfigOrDie()
	kc.ClientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal("Error: unable to create normal Kubernetes clientSet.")
	}
	kc.DynamicClient = dynamic.NewForConfigOrDie(config)
	kc.Discovery = kc.ClientSet.Discovery()
}

// InitialSetup doc
func InitialSetup() *fiber.App {
	config.NewConfig()
	conf := config.AppConfig
	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: conf.Prod, // go run app.go -prod
	})

	newK8SClient()
	setupRoutesandMiddleware(app, false)

	// Return the configured app
	return app
}
