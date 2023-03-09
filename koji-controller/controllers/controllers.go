package controllers

import (
    corev1 "k8s.io/api/core/v1"
    ctrl "sigs.k8s.io/controller-runtime"
    "sigs.k8s.io/controller-runtime/pkg/reconcile"
    "sigs.k8s.io/controller-runtime/pkg/client"

    configmapCtrl "github.com/odra/kloudji/koji-controller/controllers/configmap"
)

// add one controller to a Manager
func AddToManager(mgr ctrl.Manager, r reconcile.Reconciler, obj client.Object) error {
    return ctrl.NewControllerManagedBy(mgr).
           For(obj).
           Complete(r)
}

// Adds all controllers to a manager
func AddAllToManager(mgr ctrl.Manager) error {
    var err error
    // configmap controller
    r := configmapCtrl.New(mgr.GetClient())
    obj := &corev1.ConfigMap{}
    if err = AddToManager(mgr, r, obj); err != nil {
        return err
    }

    return nil
}
