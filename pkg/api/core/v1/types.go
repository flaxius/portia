package v1

import (
	metav1 "github.com/flaxius/portia/apimachinery/pkg/apis/meta/v1"
)

type ServiceAccount struct {
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
}
