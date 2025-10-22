FROM golang:latest AS builder

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=1 go build -o /go/bin/app cmd/*.go

FROM gcr.io/distroless/base-debian12
COPY --from=builder /go/bin/app /
ENTRYPOINT ["/app"]
