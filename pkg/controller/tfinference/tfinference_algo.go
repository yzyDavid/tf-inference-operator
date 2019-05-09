package tfinference

import (
	servingv1 "github.com/yzyDavid/tf-inference-operator/pkg/apis/serving/v1"
	"hash/fnv"
)

type DeploymentMeta struct {
	Hash uint32 `json:"hash"`
	Models []string `json:"models"`
	Replicas int32 `json:"replicas"`
}

func getHashByModels(models []string) uint32 {
	hasher := fnv.New32a()
	for _, model := range models {
		_, err := hasher.Write([]byte(model))
		if err != nil {
            // TODO
		}
	}
	return hasher.Sum32()
}

func getDeploymentMetas(models []servingv1.Model, nodes []servingv1.Node) []DeploymentMeta {
	// TODO: real impl
	var deployment DeploymentMeta
	var names []string
    for _, model := range models {
		names = append(names, model.Name)
	}
    deployment.Models = names
    deployment.Hash = getHashByModels(names)
    deployment.Replicas = 3
    return []DeploymentMeta{deployment}
}
