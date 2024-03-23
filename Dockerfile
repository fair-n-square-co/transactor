FROM golang:1.21 as build
WORKDIR /app
COPY . .
ENV GOOS linux
ENV CGO_ENABLED 0
RUN go build -v -o app ./cmd/transactions

# Copy the binary to distroless image
FROM gcr.io/distroless/base as prod
COPY --from=build /app .

EXPOSE 8080
CMD ["./app"]
