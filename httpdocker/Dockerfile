FROM golang as builder

LABEL "Author"="kkishore"

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# Final Stage

FROM scratch

COPY --from=builder /app/httpdocker /app/

EXPOSE 9999

ENTRYPOINT ["/app/httpdocker"]