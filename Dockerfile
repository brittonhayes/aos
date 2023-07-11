FROM golang:alpine AS builder
WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 go build -o /bin/warhammerd cmd/warhammerd/main.go

ARG DATABASE_READONLY_URL
ARG READONLY_USER
ARG PGDATABASE
ARG PGHOST
ARG PGPASSWORD
ARG PGPORT  
ARG PGUSER

ARG PORT
ENV PORT ${PORT}

EXPOSE 8080

ENTRYPOINT [ "/bin/warhammerd" ]

CMD ["/bin/warhammerd"]