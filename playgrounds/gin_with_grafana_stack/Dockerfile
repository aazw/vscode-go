ARG GOLANG_VERSION=1.24

FROM golang:${GOLANG_VERSION} AS go-builder

COPY go.mod go.sum main.go /ginapp/
WORKDIR /ginapp/
ENV GOOS=linux
RUN go build -o /ginapp/ginapp main.go

FROM debian
COPY --from=go-builder /ginapp/ginapp /usr/local/bin/

EXPOSE 8080
CMD ["ginapp"]