FROM golang:1.22.0-bullseye as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the Go app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dentistMigrate ./dentists/migrations/dentistMigrate.go


######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN apk add --no-cache postgresql-client

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/dentistMigrate .

COPY --from=builder /app/entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/entrypoint.sh

EXPOSE 5000

ENTRYPOINT ["entrypoint.sh"]

# Command to run the executable
CMD ["./main"]