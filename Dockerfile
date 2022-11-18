FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/gitlab.com/idoko/HyperSkill/
WORKDIR /go/src/gitlab.com/idoko/HyperSkill
RUN go mod download
COPY . /go/src/gitlab.com/idoko/HyperSkill
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/HyperSkill gitlab.com/idoko/HyperSkill

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/gitlab.com/idoko/HyperSkill/build/HyperSkill /usr/bin/HyperSkill

EXPOSE ${PORT} ${PORT}
ENTRYPOINT ["/usr/bin/HyperSkill"]