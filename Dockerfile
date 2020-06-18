FROM golang:1.14.2 AS base
    WORKDIR ${GOPATH}/src/github.com/victorabarros/work-at-olist/
    COPY . ./
    # RUN go mod vendor && \
    RUN go build main.go

    CMD ["./main"]
