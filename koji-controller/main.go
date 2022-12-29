package main

import (
	"os"

	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"

	"github.com/odra/cowji/koji-controller/helpers"
    controllers "github.com/odra/cowji/koji-controller/controllers"
)

func init() {
    log.SetLogger(zap.New())
}

func main() {
    ctrlName := helpers.GetEnv("KUBE_CTRL_NAME", "koji-controller")
    logName := helpers.GetEnv("LOG_NAME", ctrlName)
    l := log.Log.WithName(logName)

    // manager setup
    l.Info("manager setup start")
    config, err := config.GetConfig()
    if err != nil {
        l.Error(err, "unable to get kubernetes rest config")
        os.Exit(1)
    }
    mgr, err := manager.New(config, manager.Options{})
    if err != nil {
        l.Error(err, "unable to create a new manager")
        os.Exit(1)
    }

    l.Info("adding all controllers to manager")
    err = controllers.AddAllToManager(mgr)
    if err != nil {
        l.Error(err, "failed to setup configmap controller")
        os.Exit(1)
    }

    l.Info("start manager")
    err = mgr.Start(signals.SetupSignalHandler())
    if err != nil {
        l.Error(err, "unable to start manager")
        os.Exit(1)
    }
}
