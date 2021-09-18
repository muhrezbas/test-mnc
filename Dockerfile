FROM golang:1.13-alpine3.10 AS builder

RUN apk update && apk upgrade && apk --no-cache --update add ca-certificates tzdata

# Working Directory
WORKDIR $GOPATH/src/mnc

# Copying Data
COPY . .

# View Of Current Directory
RUN echo $PWD && ls -la

# Fetch dependencies.
RUN go get -d -v

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/mnc .

FROM scratch

# LABEL Maintainer
LABEL maintainer="muhrezbasuki@gmail.com" product="test" category="mnc"

# Copy the executable.
COPY --from=builder /go/bin/mnc /go/bin/mnc
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/go/bin/mnc"]
