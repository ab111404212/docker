# Version 1.0
# 基础镜像
FROM     ubuntu:14.04
# 维护者信息
MAINTAINER 18217007736@163.com

RUN      apt-get update && apt-get install golang
WORKDIR  docker
CMD      cd docker/deploy/
RUN      go build main.go
RUN      ./main