#!/bin/bash

#Build linux-version first and determine the build-version
go build .
BUILD_VERSION=$(./space-cli -v | cut -f3 -d ' ')

mkdir linux && mkdir windows && mkdir darwin

echo 'linux build'
GOOS=linux GOARCH=amd64 go build -ldflags '-s -w -extldflags "-static"' .
zip space-cli.zip space-cli
mv ./space-cli.zip ./linux/
cp ./linux/space-cli.zip ./linux/space-cli_v$BUILD_VERSION.zip 
rm space-cli

echo 'darwin build'
GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w -extldflags "-static"' .
zip space-cli.zip space-cli
mv ./space-cli.zip ./darwin/
cp ./darwin/space-cli.zip ./darwin/space-cli_v$BUILD_VERSION.zip 
rm space-cli

echo 'windows build'
GOOS=windows GOARCH=amd64 go build -ldflags '-s -w -extldflags "-static"' .
zip space-cli.zip space-cli.exe
mv ./space-cli.zip ./windows/
cp ./windows/space-cli.zip ./windows/space-cli_v$BUILD_VERSION.zip 
rm space-cli.exe

# echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
# docker push spaceuptech/space-cli:latest
JWT_TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.GZObi26RixeVO44D_rcEdIPO9nsLpJm6OR7tNECtO0E"
#Upload as versioned build

curl -H "Authorization: Bearer $JWT_TOKEN" -F "file=@./darwin/space-cli_v$BUILD_VERSION.zip" -F 'fileType=file' -F 'makeAll=false' -F 'path=/darwin' https://spaceuptech.com/v1/api/downloads/files
curl -H "Authorization: Bearer $JWT_TOKEN" -F "file=@./windows/space-cli_v$BUILD_VERSION.zip" -F 'fileType=file' -F 'makeAll=false' -F 'path=/windows' https://spaceuptech.com/v1/api/downloads/files
curl -H "Authorization: Bearer $JWT_TOKEN" -F "file=@./linux/space-cli_v$BUILD_VERSION.zip" -F 'fileType=file' -F 'makeAll=false' -F 'path=/linux' https://spaceuptech.com/v1/api/downloads/files
#Upload as latest build
curl -H "Authorization: Bearer $JWT_TOKEN" -F 'file=@./darwin/space-cli.zip' -F 'fileType=file' -F 'makeAll=false' -F 'path=/darwin' https://spaceuptech.com/v1/api/downloads/files
curl -H "Authorization: Bearer $JWT_TOKEN" -F 'file=@./windows/space-cli.zip' -F 'fileType=file' -F 'makeAll=false' -F 'path=/windows' https://spaceuptech.com/v1/api/downloads/files
curl -H "Authorization: Bearer $JWT_TOKEN" -F 'file=@./linux/space-cli.zip' -F 'fileType=file' -F 'makeAll=false' -F 'path=/linux' https://spaceuptech.com/v1/api/downloads/files

rm -rf windows darwin linux