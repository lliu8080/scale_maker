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
)

var testApp *fiber.App

func TestInitSetupSuccess(t *testing.T) {
	testApp = InitialTestSetup()
	assert.NotNil(t, kc.clientSet)
}

func newK8STestClient() {
	var err error
	kc.ctx = context.TODO()
	kc.clientSet = fakek8s.NewSimpleClientset()
	if err != nil {
		log.Fatal("Error: unable to create normal Kubernetes clientSet.")
	}
	kc.dynamicClient = fakeDynamic.NewSimpleDynamicClient(runtime.NewScheme())
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