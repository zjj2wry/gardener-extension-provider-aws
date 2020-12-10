package helmvalues

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

const (
	MachineControllerManager = "machine-controller-manager"
	CloudControllerManager   = "cloud-controller-manager"
)

var OverrideHelmValues *unstructured.Unstructured

func Load(v *unstructured.Unstructured) {
	OverrideHelmValues = v
}

func HelmValuesFor(name string) map[string]interface{} {
	emptyValues := make(map[string]interface{})
	if OverrideHelmValues == nil {
		return emptyValues
	}
	data := OverrideHelmValues.UnstructuredContent()

	if len(data) == 0 {
		return emptyValues
	}

	values, ok := data[name]
	if !ok {
		return emptyValues
	}
	overrideValues, ok := values.(map[string]interface{})
	if !ok {
		return emptyValues
	}
	return overrideValues
}
