#!/usr/bin/env bash

# remove CR & CRD
kubectl delete -f deploy/crds/serving_v1_tfinference_cr.yaml
kubectl delete -f deploy/crds/serving_v1_tfinference_crd.yaml
