// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/yzyDavid/tf-inference-operator/pkg/apis/serving/v1.TfInference":       schema_pkg_apis_serving_v1_TfInference(ref),
		"github.com/yzyDavid/tf-inference-operator/pkg/apis/serving/v1.TfInferenceSpec":   schema_pkg_apis_serving_v1_TfInferenceSpec(ref),
		"github.com/yzyDavid/tf-inference-operator/pkg/apis/serving/v1.TfInferenceStatus": schema_pkg_apis_serving_v1_TfInferenceStatus(ref),
	}
}

func schema_pkg_apis_serving_v1_TfInference(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TfInference is the Schema for the tfinferences API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/yzyDavid/tf-inference-operator/pkg/apis/serving/v1.TfInferenceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/yzyDavid/tf-inference-operator/pkg/apis/serving/v1.TfInferenceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/yzyDavid/tf-inference-operator/pkg/apis/serving/v1.TfInferenceSpec", "github.com/yzyDavid/tf-inference-operator/pkg/apis/serving/v1.TfInferenceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_serving_v1_TfInferenceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TfInferenceSpec defines the desired state of TfInference",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_serving_v1_TfInferenceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TfInferenceStatus defines the observed state of TfInference",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
