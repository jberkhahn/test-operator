package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FoobarSpec defines the desired state of Foobar
type FoobarSpec struct {
	// input string
	Input string `json:"input"`
}

// FoobarStatus defines the observed state of Foobar
type FoobarStatus struct {
	// output string
	Output string `json:"output"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Foobar is the Schema for the foobars API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=foobars,scope=Namespaced
type Foobar struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FoobarSpec   `json:"spec,omitempty"`
	Status FoobarStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FoobarList contains a list of Foobar
type FoobarList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Foobar `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Foobar{}, &FoobarList{})
}
