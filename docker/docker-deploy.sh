VERSION=dev-1.0.1
DOCKER_TAG=trongnv138/golang:go-storage-$VERSION
cd ..
docker build -t $DOCKER_TAG -f docker/Dockerfile .
docker push $DOCKER_TAG