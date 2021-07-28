package initializer

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
)

type CustomAction struct {
	name            string
	receivedVersion string
}

func (a *CustomAction) Run(version string, kubeClient *kubernetes.Clientset) error {
	if kubeClient == nil {
		return fmt.Errorf("kubeClient is expected but was nil")
	}
	fmt.Printf("Action '%s': received version '%s'\n", a.name, version)
	a.receivedVersion = version
	return nil
}
