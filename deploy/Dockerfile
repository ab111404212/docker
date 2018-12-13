# Version 1.0
# 基础镜像
FROM     ab111404212/docker-ubuntu
# 维护者信息
MAINTAINER 18217007736@163.com
#RUN      apt-get update && apt-get -f install  golang
RUN      mkdir docker
RUN      cd docker/deploy/
RUN      go build main.go
CMD      /bin/bash