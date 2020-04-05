# offical golang docker image
FROM golang:latest
# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name
ENV SRC_DIR=/go/src/github.com/curtis-turner/gomortgage
# Add binary to /go/bin
COPY . $SRC_DIR
RUN cd $SRC_DIR; go build -o gomortgage; cp $SRC_DIR/gomortgage /go/bin/
ENTRYPOINT [ "gomortgage" ]