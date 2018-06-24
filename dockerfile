# start a golang base image for building
FROM golang:alpine as builder

# scaffold fs tree
RUN mkdir -p /go/src/market
WORKDIR /go/src/market
RUN mkdir vendor

#copy the source files
COPY main.go /go/src/market
COPY vendor/* vendor/

#disable crosscompiling
ENV CGO_ENABLED=0
#compile linux only
ENV GOOS=linux

#build lightweight static-linked executable sans bloat
RUN go build -ldflags="-s -w" -a -installsuffix cgo -o market .

#############################################################################

# start with scratch (no layers)
FROM scratch

# copy our assets
COPY --from=builder /go/src/market/market .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# expose port for process
EXPOSE 8080
# run it!
CMD ["./market"]