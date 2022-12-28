package config

import (
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

// gets k8s cluster rest api config
func GetKubeConfig() (*rest.Config, error) {
	return config.GetConfig()
}
