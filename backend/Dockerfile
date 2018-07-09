FROM golang:alpine as builder

ENV SRC=/go/src/github.com/ComputePractice2018/todolist/backend
RUN mkdir -p '${SRC}'
ADD . ${SRC}

RUN apk add --no-cache git

RUN go get -u github.com/gorilla/mux

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/todolist github.com/ComputePractice2018/todolist/backend/

FROM alpine:3.7

WORKDIR /opt
COPY --from=builder /go/bin/todolist /opt/todolist

EXPOSE 8080

ENV NAME "Egor"
ENTRYPOINT /opt/todolist -name ${NAME}