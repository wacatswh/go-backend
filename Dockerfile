FROM --platform=linux/arm64 golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-backend

FROM scratch

COPY --from=builder /go-backend /go-backend

EXPOSE 8080

ENTRYPOINT [ "/go-backend" ]