FROM golang:1.21.5 as builder
WORKDIR /build
COPY go.* ./
COPY ./cmd ./cmd
RUN CGO_ENABLED=0 go build -o ./bin/healthchecker ./cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /build/bin/healthchecker ./healthchecker

HEALTHCHECK --interval=5s --timeout=3s --start-period=5s --retries=3 \
CMD [ "/app/healthchecker", "-port", "3000", "-path", "/healthcheck" ]