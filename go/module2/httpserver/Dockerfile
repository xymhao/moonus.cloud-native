FROM golang:1.17 AS build
WORKDIR /httpserver/
COPY . .
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o httpserver httpserver.go

FROM busybox
COPY --from=build /httpserver/httpserver /httpserver/httpserver
EXPOSE 80
ENV ENV local
WORKDIR /httpserver/
ENTRYPOINT ["./httpserver -port 80"]