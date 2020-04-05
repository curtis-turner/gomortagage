# offical golang docker image
FROM golang:latest
# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name
ENV SRC_DIR=/go/src/github.com/curtis-turner/gomortage
# Add binary to /go/bin
ADD . $SRC_DIR
ENTRYPOINT ["gomortgage"]