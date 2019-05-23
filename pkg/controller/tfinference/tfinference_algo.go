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

func deploymentInSlice(a DeploymentMeta, list []DeploymentMeta) bool {
	for _, b := range list {
		if b.Hash == a.Hash {
			return true
		}
	}
	return false
}

func getDeploymentMetas(models []servingv1.Model, nodes []servingv1.Node, groups []servingv1.Group) []DeploymentMeta {
	deployments := make([]DeploymentMeta, 0)
	NodeMap := make(map[int][]string)

	sort.SliceStable(models, func(i, j int) bool {
		return models[i].Memory > models[j].Memory
	})
	sort.SliceStable(nodes, func(i, j int) bool {
		return nodes[i].Memory > nodes[j].Memory
	})
	sort.SliceStable(groups, func(i, j int) bool {
		return groups[i].ComputingResource > groups[j].ComputingResource
	})
	nodesLeft := make([]servingv1.Node, len(nodes))
	modelsLeft := make([]servingv1.Model, len(models))
	groupsLeft := make([]servingv1.Group,len(groups))
	copy(modelsLeft, models)
	copy(nodesLeft, nodes)
	copy(groupsLeft,groups)

	// grouped models
	for len(groupsLeft) > 0 {
		group := groupsLeft[0]
		models := make([]servingv1.Model,0)
		for _, modelName := range group.GroupedModels {
			for i:= 0; i < len(modelsLeft);i++ {
				if modelsLeft[i].Name == modelName {
					models = append(models,modelsLeft[i])
					copy(modelsLeft[i:],modelsLeft[i+1:])
					modelsLeft = modelsLeft[0:len(modelsLeft)-1]
				}
			}
		}
		nodeCnt := 0
		for len(models) > 0 {
			for nodesLeft[nodeCnt].ComputingResource < groupsLeft[0].ComputingResource {
				nodeCnt++
			}
			nodesLeft[nodeCnt].ComputingResource -= groupsLeft[0].ComputingResource
			for i := 0; i < len(models); i++ {
				nodesLeft[nodeCnt].Memory -= models[i].Memory
				if nodesLeft[nodeCnt].Memory < 0 {
					nodesLeft[nodeCnt].Memory += models[i].Memory
					continue
				} else {
					NodeMap[nodeCnt] = append(NodeMap[nodeCnt], models[i].Name)
					copy(models[i:], models[i+1:])
					models = models[0 : len(models)-1]
					i--
				}
			}
		}
		if len(groupsLeft) > 1 {
			groupsLeft = groupsLeft[1:]
		} else {
			break
		}
	}

	// normal models
	nodeCnt := 0
	for len(modelsLeft) > 0 && len(nodesLeft) > 0 {
		var deployment DeploymentMeta
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
				NodeMap[nodeCnt] = append(NodeMap[nodeCnt], modelsLeft[i].Name)
				copy(modelsLeft[i:],modelsLeft[i+1:])
				modelsLeft = modelsLeft[0:len(modelsLeft)-1]
				i--
			}
		}
		deployment.Models = NodeMap[nodeCnt]
		deployment.Hash = getHashByModels(NodeMap[nodeCnt])
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
		if len(nodesLeft) > 1 {
			nodesLeft = nodesLeft[1:]
		} else {
			break
		}
		nodeCnt++
	}
	return deployments
}
