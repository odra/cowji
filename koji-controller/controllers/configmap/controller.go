package configmap

import (
	"context"

	"github.com/odra/kloudji/koji-controller/helpers"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type Reconciler struct {
    client client.Client
}

func New(client client.Client) *Reconciler {
    return &Reconciler{client}
}

var _ reconcile.Reconciler = &Reconciler{}

func (r * Reconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
    var err error
    l := log.FromContext(ctx)
    obj := &corev1.ConfigMap{}

    err = r.client.Get(ctx, req.NamespacedName, obj)
    if err != nil {
        if errors.IsNotFound(err) {
            l.Error(nil, "could not find ConfigMap")
            return reconcile.Result{}, nil
        } else {
            l.Error(
                err,
                "unable to fetch the ConfigMap object: %s",
                helpers.KeyToString(req.NamespacedName),
            )

            return reconcile.Result{}, err
        }
    }

    l.Info(
        "Reconciling ConfigMap for %s",
        helpers.KeyToString(req.NamespacedName),
    )
    
    if obj.Labels == nil {
        obj.Labels = map[string]string{}
    }
    val, ok := obj.Labels["foo"]
    if ok && val == "bar" {
        return reconcile.Result{}, nil
    }

    obj.Labels["foo"]  = "bar"
    err = r.client.Update(ctx, obj)
    if err != nil {
        return reconcile.Result{}, err
    }

    return reconcile.Result{}, nil
}
