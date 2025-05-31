ARG GOLANG_VERSION=1.24@sha256:81bf5927dc91aefb42e2bc3a5abdbe9bb3bae8ba8b107e2a4cf43ce3402534c6

FROM golang:${GOLANG_VERSION} AS go-builder

COPY go.mod go.sum main.go /ginapp/
WORKDIR /ginapp/
ENV GOOS=linux
RUN go build -o /ginapp/ginapp main.go

FROM debian@sha256:bd73076dc2cd9c88f48b5b358328f24f2a4289811bd73787c031e20db9f97123
COPY --from=go-builder /ginapp/ginapp /usr/local/bin/

EXPOSE 8080
CMD ["ginapp"]