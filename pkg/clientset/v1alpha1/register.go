package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	api "vdesjardins/acm-manager/pkg/api/v1alpha1"
)

var (
	schemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = schemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(api.GroupVersion,
		&api.Certificate{},
	)

	metav1.AddToGroupVersion(scheme, api.GroupVersion)
	return nil
}
