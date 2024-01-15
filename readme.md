Tool to add a healthcheck for "FROM SCRATCH" docker images.

### Usage

```bash
./healthchecker -port 3000 -path /healthcheck
./healthchecker -h
```

### Multi-stage build example

Since the image is a "FROM SCRATCH" image, you can just build your image from the healthchecker image :

```dockerfile
# main application build
FROM golang:1.21.5 as builder
WORKDIR /build
COPY ./go.* ./
COPY ./ ./
RUN CGO_ENABLED=0 go build -o ./bin/main ./

FROM ermineaweb/healthchecker:latest
WORKDIR /app
COPY --from=builder /build/bin/main ./main

# you can surcharge the healthcheck
HEALTHCHECK --interval=5s --timeout=3s --start-period=5s --retries=3 \
CMD [ "/app/healthchecker", "-port", "3000", "-path", "/healthcheck" ]

ENTRYPOINT ["./main"]
```

### Build locally then copy the binary

```bash
docker build -t healthchecker:latest .
```

Then in the dockerfile add these lines :

```dockerfile
WORKDIR /app
COPY --from=healthchecker:latest /build/bin/healthchecker ./healthchecker
```

### Import the binary from docker hub

Add to the dockerfile these lines :

```dockerfile
WORKDIR /app
COPY --from=ermineaweb/healthchecker:latest /build/bin/healthchecker ./healthchecker
```
