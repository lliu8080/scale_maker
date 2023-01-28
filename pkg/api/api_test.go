package api

import (
	"context"
	"log"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime"
	fakeDynamic "k8s.io/client-go/dynamic/fake"
	fakek8s "k8s.io/client-go/kubernetes/fake"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/config"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/k8s"
)

var testApp *fiber.App

func TestInitSetupSuccess(t *testing.T) {
	testApp = InitialTestSetup()
	assert.NotNil(t, kc.ClientSet)
}

func newK8STestClient() {
	var err error
	kc.Ctx = context.TODO()
	kc.ClientSet = fakek8s.NewSimpleClientset()
	if err != nil {
		log.Fatal("Error: unable to create normal Kubernetes clientSet.")
	}
	kc.DynamicClient = fakeDynamic.NewSimpleDynamicClient(runtime.NewScheme())
	kc.Discovery = k8s.SetupDiscovery(kc)
}

// InitialSetup doc
func InitialTestSetup() *fiber.App {
	config.NewConfig()
	conf := config.AppConfig
	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: conf.Prod, // go run app.go -prod
	})
	newK8STestClient()
	setupRoutesandMiddleware(app, true)

	// Return the configured app
	return app
}
