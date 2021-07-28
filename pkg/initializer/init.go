package initializer

import "github.com/kyma-incubator/reconciler/pkg/reconciler/service"

var Reconciler *service.ComponentReconciler

func init() {
	var err error
	Reconciler, err = service.NewComponentReconciler("reconciler1", ".", true)
	if err != nil {
		panic(err)
	}
}
