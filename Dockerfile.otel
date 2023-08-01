FROM otel/opentelemetry-collector-contrib:latest

LABEL maintainer="brittonhayes"
LABEL org.opencontainers.image.source="https://github.com/brittonhayes/aos"
LABEL org.opencontainers.image.description="AOS API specific version of OpenTelemetry collector."

COPY ./config/otel-config.prod.yaml /etc/otel-collector-config.yaml