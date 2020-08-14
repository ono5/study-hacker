FROM golang:latest

ENV SRC_DIR=/go/src/github.com/ono5/study-hacker/api

WORKDIR $SRC_DIR

ADD ./api $SRC_DIR

RUN go get -u google.golang.org/grpc \
    && go get -u github.com/golang/protobuf/protoc-gen-go \
    && go get go.mongodb.org/mongo-driver/mongo \
    && go get -u github.com/labstack/echo/... \
    && go get github.com/oxequa/realize

CMD ["go", "run", "main.go"]
