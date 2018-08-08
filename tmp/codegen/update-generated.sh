#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

vendor/k8s.io/code-generator/generate-groups.sh \
deepcopy \
github.com/aveshagarwal/kube-scheduler-operator/pkg/generated \
github.com/aveshagarwal/kube-scheduler-operator/pkg/apis \
kube-scheduler:v1alpha1 \
--go-header-file "./tmp/codegen/boilerplate.go.txt"
