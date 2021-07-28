package main

import (
	"context"
	"fmt"
	"github.com/tobiscr/poc-embedded-reconciler-1/pkg/initializer"
	"time"
)

func main() {
	fmt.Println("Starting poc-embedded-reconciler-1")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	if err := initializer.Reconciler.StartRemote(ctx); err != nil {
		panic(err)
	}
}
