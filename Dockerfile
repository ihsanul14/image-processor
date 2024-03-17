FROM golang:alpine

COPY . /go/src/image-processor
WORKDIR /go/src/image-processor

RUN apk --no-cache add \
    build-base \
    cmake \
    git \
    linux-headers \
    gcc \
    ffmpeg \
    musl-dev \
    pkgconf \
    pkgconfig \
    opencv-dev \
    && rm -rf /var/cache/apk/*

RUN go mod tidy

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -o /image-processor .

EXPOSE 30001

ENTRYPOINT ["/image-processor"]