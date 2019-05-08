#!/usr/bin/env bash

# start local operator
export OPERATOR_NAME=tf-inference-operator
operator-sdk up local --namespace=default
