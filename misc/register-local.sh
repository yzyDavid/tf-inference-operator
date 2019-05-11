#!/usr/bin/env bash

# register CRD
kubectl create -f deploy/crds/serving_v1_tfinference_crd.yaml
kubectl create -f deploy/nfs_local.yaml
kubectl create -f deploy/nfs_pvc_local.yaml
