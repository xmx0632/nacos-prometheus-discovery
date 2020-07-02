#!/bin/sh
echo "exec ${BASH_SOURCE[0]}"

export BASE_DIR=`cd $(dirname $0); pwd`
echo $BASE_DIR

TARGET_DIR=target

rm -rf $TARGET_DIR
mkdir -p $TARGET_DIR/bin

cp -rf conf $TARGET_DIR/bin
cd $TARGET_DIR/bin

go get github.com/karalabe/xgo
xgo --targets=darwin/amd64,linux/amd64,windows/amd64,windows/386  gitee.com/xmx0632/nacos-prometheus-discovery


# build all platform
# xgo gitee.com/xmx0632/nacos-prometheus-discovery

#The supported targets are:
#
#Platforms: android, darwin, ios, linux, windows
#Achitectures: 386, amd64, arm-5, arm-6, arm-7, arm64, mips, mipsle, mips64, mips64le

#xgo -go 1.6.1 gitee.com/xmx0632/nacos-prometheus-discovery
#xgo -out iris-v0.3.2 gitee.com/xmx0632/nacos-prometheus-discovery
#xgo --branch release-branch.go1.4 golang.org/x/tools/cmd/goimports
#xgo --pkg cmd/goimports golang.org/x/tools

# return
cd $BASE_DIR