FROM     ubuntu
RUN      apt-get install golang
WORKDIR  docker
CMD      cd docker/
RUN      go build main.go
RUN      ./main