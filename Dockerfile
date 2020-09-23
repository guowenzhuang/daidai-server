FROM golang:latest


ENV GOPROXY https://goproxy.cn,direct
WORKDIR /Users/guowenzhuang/go/src/daidai-server
COPY . /Users/guowenzhuang/go/src/daidai-server
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./daidai-server"]


