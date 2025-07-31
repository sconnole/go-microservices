#!/bin/bash

set -e

IMAGE_NAME="job-creator"       
IMAGE_TAG="local"                 
K3D_CLUSTER_NAME="go-microservices"  

echo "Building Docker image: $IMAGE_NAME:$IMAGE_TAG"
docker build -t $IMAGE_NAME:$IMAGE_TAG .

echo "Importing image into k3d cluster: $K3D_CLUSTER_NAME"
k3d image import $IMAGE_NAME:$IMAGE_TAG -c $K3D_CLUSTER_NAME

