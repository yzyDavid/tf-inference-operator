#!/usr/bin/env bash

# revert create-local.sh
kubectl delete -f deploy/crds/serving_v1_tfinference_cr.yaml
