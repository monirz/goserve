FROM golang:1.19.3-alpine AS build
ENV CGO_ENABLED=0
ENV GO111MODULE=on
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /server

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /server .
ENV PORT=8090
EXPOSE $PORT
CMD ["./server"]