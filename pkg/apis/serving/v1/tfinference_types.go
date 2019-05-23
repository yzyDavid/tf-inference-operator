package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Added by Zhenyun Yu, types here are necessary for specs
type Model struct {
	Name string `json:"name"`
    Memory int32 `json:"memory"`  // in MB
    ComputingResource int32 `json:"computing_resource"`  // 100 for each core
}

type Node struct {
	Memory int32 `json:"memory"`
	ComputingResource int32 `json:"computing_resource"`
}

// If a model is grouped, its specs for computing resource will be overwritten by computing resource of group
type Group struct {
	GroupedModels []string `json:"grouped_models"`
	ComputingResource int32 `json:"computing_resource"`
}

// TfInferenceSpec defines the desired state of TfInference
// +k8s:openapi-gen=true
type TfInferenceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
    Models []Model `json:"models"`
    Groups []Group `json:"groups"`
    Nodes []Node `json:"nodes"`
}

// TfInferenceStatus defines the observed state of TfInference
// +k8s:openapi-gen=true
type TfInferenceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Deployments []string `json:"deployments"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TfInference is the Schema for the tfinferences API
// +k8s:openapi-gen=true
type TfInference struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TfInferenceSpec   `json:"spec,omitempty"`
	Status TfInferenceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TfInferenceList contains a list of TfInference
type TfInferenceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TfInference `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TfInference{}, &TfInferenceList{})
}
