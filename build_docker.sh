#!/bin/bash

set -e
VERSION=$(cat version.txt)
IMAGE_NAME="josephwesselman/project-imgs:go-server-$VERSION"

echo "building image:  $IMAGE_NAME"
docker build -t $IMAGE_NAME .

docker push $IMAGE_NAME