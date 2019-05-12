package tfinference

import (
	servingv1 "github.com/yzyDavid/tf-inference-operator/pkg/apis/serving/v1"
	"hash/fnv"
	"sort"
)

type DeploymentMeta struct {
	Hash     uint32   `json:"hash"`
	Models   []string `json:"models"`
	Replicas int32    `json:"replicas"`
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

func removeModel(s []servingv1.Model, i int) []servingv1.Model {
	return append(s[:i], s[i+1:]...)
}

func removeNode(s []servingv1.Node, i int) []servingv1.Node {
	return append(s[:i], s[i+1:]...)
}

func deploymentInSlice(a DeploymentMeta, list []DeploymentMeta) bool {
	for _, b := range list {
		if b.Hash == a.Hash {
			return true
		}
	}
	return false
}

func getDeploymentMetas(models []servingv1.Model, nodes []servingv1.Node) []DeploymentMeta {
	// TODO: real impl
	deployments := make([]DeploymentMeta, 0)

	sort.SliceStable(models, func(i, j int) bool {
		return models[i].Memory > models[j].Memory
	})
	sort.SliceStable(nodes, func(i, j int) bool {
		return nodes[i].Memory > nodes[j].Memory
	})
	nodesLeft := make([]servingv1.Node, 0)
	modelsLeft := make([]servingv1.Model, 0)
	copy(modelsLeft, models)
	copy(nodesLeft, nodes)
	for len(nodesLeft) > 0 {
		var deployment DeploymentMeta
		names := make([]string, 0)
		for i := 0; i < len(modelsLeft); i++ {
			nodesLeft[0].Memory -= modelsLeft[i].Memory
			nodesLeft[0].ComputingResource -= modelsLeft[i].ComputingResource
			if nodesLeft[0].Memory < 0 {
				nodesLeft[0].Memory += modelsLeft[i].Memory
				continue
			}
			if nodesLeft[0].ComputingResource < 0 {
				nodesLeft[0].ComputingResource += modelsLeft[i].ComputingResource
				continue
			} else {
				names = append(names, modelsLeft[i].Name)
				removeModel(modelsLeft, i)
				i--
			}
		}
		deployment.Models = names
		deployment.Hash = getHashByModels(names)
		if deploymentInSlice(deployment, deployments) {
			for _, dep := range deployments {
				if dep.Hash == deployment.Hash {
					dep.Replicas++
				}
			}
		} else {
			deployment.Replicas = 1
			deployments = append(deployments, deployment)
		}
		removeNode(nodesLeft, 0)
	}
	return deployments
}
