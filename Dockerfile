FROM golang:latest AS builder

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app cmd/*.go

FROM gcr.io/distroless/static-debian12
COPY --from=builder /go/bin/app /
ENTRYPOINT ["/app"]
