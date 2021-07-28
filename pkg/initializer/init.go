package initializer

import (
	"github.com/kyma-incubator/reconciler/pkg/reconciler/service"
	"time"
)

var Reconciler *service.ComponentReconciler

func init() {
	var err error
	Reconciler, err = service.NewComponentReconciler("reconciler1", ".", true)
	if err != nil {
		panic(err)
	}

	preAct := &CustomAction{
		name: "pre-action of reconciler-1",
	}
	instAct := &CustomAction{
		name: "install-action of reconciler-1",
	}
	postAct := &CustomAction{
		name: "post-action of reconciler-1",
	}
	//configure reconciler
	Reconciler.WithRetry(1, 1*time.Second).
		WithPreInstallAction(preAct).
		WithInstallAction(instAct).
		WithPostInstallAction(postAct)

}
