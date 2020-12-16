package helmvalues

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestHelmValuesFor(t *testing.T) {
	emptyValues := make(map[string]interface{})
	testCases := []struct {
		name               string
		overrideHelmValues *unstructured.Unstructured
		expectValues       map[string]interface{}
	}{
		{
			name:         "override helm values is null",
			expectValues: emptyValues,
		},
		{
			name: "override helm values is not map",
			overrideHelmValues: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"etcd": "i am string",
				},
			},
			expectValues: emptyValues,
		},
		{
			name: "add extra annotations to etcd service account",
			overrideHelmValues: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"etcd": map[string]interface{}{
						"serviceAccountAnnotations": map[string]interface{}{
							"eks.amazonaws.com/role-arn": "role-id",
						},
					},
				},
			},
			expectValues: map[string]interface{}{
				"serviceAccountAnnotations": map[string]interface{}{
					"eks.amazonaws.com/role-arn": "role-id",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			OverrideHelmValues = testCase.overrideHelmValues
			overrideValues := HelmValuesFor("etcd")
			if diff := cmp.Diff(overrideValues, testCase.expectValues); diff != "" {
				t.Fatalf("diff: %s", diff)
			}
		})
	}
}
