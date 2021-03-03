docker build -f Dockerfile.client --build-arg BUILD_ID . -t clientapi
docker build -f Dockerfile.server --build-arg BUILD_ID . -t portdomain
docker image prune --filter label=stage=builder --filter label=build=$BUILD_ID