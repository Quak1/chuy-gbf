FROM node:24-alpine AS frontend-builder
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm ci
COPY frontend/ ./
RUN npm run build

FROM golang:1.26-alpine AS backend-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY main.go ./
COPY --from=frontend-builder /app/dist ./dist
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

FROM alpine:3.24
WORKDIR /app
COPY --from=backend-builder /app/server .
RUN mkdir /data
ENV DB_PATH=/data/data.db
EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=1s --start-period=30s --retries=3 \
  CMD curl --fail http://localhost:8080/api/ping || exit 1

CMD ["./server"]
