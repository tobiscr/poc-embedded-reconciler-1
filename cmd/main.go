package main

import (
	"context"
	"fmt"
	"github.com/kyma-incubator/reconciler/pkg/reconciler/service"
	"time"
)

var reconciler *service.ComponentReconciler

func init() {
	var err error
	reconciler, err = service.NewComponentReconciler("reconciler1", ".", true)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Starting poc-embedded-reconciler-1")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	if err := reconciler.StartRemote(ctx); err != nil {
		panic(err)
	}
}
