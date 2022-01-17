#!/bin/bash

# 需要先安装和配置好kubectl
function checkvar {
    VAR_VALUE=$(echo "${!1}")
    if [[ -z "$VAR_VALUE" ]]; then
      echo "$1 env variable is not set"
      exit 1
    fi
}

if [[ -z "$1" ]]; then
  echo "Usage: $(basename ${0}) project [apply/delete]"
  exit 1
fi

checkvar "DOCKER_IMAGE"

project_dir=$1
command=$2

script_dir="$( cd "$(dirname "$0")" ; pwd -P )"
curr_dir=$(pwd -P)
project_dir="$( cd ${project_dir} ; pwd -P )"

echo "Deploying docker image ${DOCKER_IMAGE} to k8s \n"

manifest=${project_dir}/k8s.yml
manifest_name=$(echo "${manifest}" | sed "s/.*\///")
echo "${manifest_name} \n"

# replace environment variables
envsubst < ${manifest} > ${manifest_name}-gen.yaml
echo "Applying manifest definition \n"
cat ${manifest_name}-gen.yaml

# apply  the config
echo "Using kubeconfig $(kubectl config view -o template --template='{{ index . "current-context" }}') \n"
kubectl ${command} -f ${manifest_name}-gen.yaml -n boomer
rm -f *-gen.yaml

echo "\n Finish."