ARG GO_VERSION=1.20.6

# STAGE 1
FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /src/

COPY go.mod ./
RUN go mod download

COPY . /src/
RUN CGO_ENABLED=0 go build -o /bin/warhammerd cmd/warhammerd/main.go

# STAGE 2
FROM gcr.io/distroless/static-debian11:nonroot

LABEL maintainer="brittonhayes"

COPY --from=builder --chown=nonroot:nonroot /bin/warhammerd /bin/warhammerd

EXPOSE 8080

ENTRYPOINT [ "/bin/warhammerd" ]

CMD ["/bin/warhammerd"]