FROM golang:1.23-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY .env .env 
RUN CGO_ENABLED=0 GOOS=linux go build -o job_portal main.go
FROM alpine:3.18
WORKDIR /app
COPY --from=build /app/job_portal .
COPY --from=build /app/.env .env
RUN apk --no-cache add ca-certificates tzdata
EXPOSE 8080 9100
CMD ["./job_portal"]
