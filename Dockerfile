FROM golang:1.11 as builder

LABEL maintainer "David Ndungu <dndungu@zendesk.com>"

WORKDIR /app

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch

LABEL maintainer "David Ndungu <dndungu@zendesk.com>"

WORKDIR /app

COPY --from=builder /app/app .

COPY page.tpl .

ENV HELLO_CI 80

EXPOSE ${HELLO_CI}

ENTRYPOINT ["/app/app"]
