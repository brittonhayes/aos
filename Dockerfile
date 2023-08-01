ARG GO_VERSION=1.20

# STAGE 1
FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /src/

COPY go.mod ./
RUN go mod download

COPY . /src/
RUN CGO_ENABLED=0 go build -o /bin/aos cmd/aos/main.go

# STAGE 2
FROM gcr.io/distroless/static-debian11:nonroot

LABEL maintainer="brittonhayes"
LABEL org.opencontainers.image.source="https://github.com/brittonhayes/aos"
LABEL org.opencontainers.image.description="An unnoficial AoS REST API built with Go and Sqlite."
LABEL org.opencontainers.image.licenses="MIT"

COPY --from=builder --chown=nonroot:nonroot /bin/aos /bin/aos

EXPOSE 8080

ENTRYPOINT [ "/bin/aos" ]

CMD ["/bin/aos"]