#!/usr/bin/env bash

# revert register-local.sh
kubectl delete -f deploy/crds/serving_v1_tfinference_crd.yaml
kubectl delete -f deploy/nfs_local.yaml
kubectl delete -f deploy/nfs_pvc_local.yaml
