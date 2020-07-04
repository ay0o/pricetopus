FROM golang:alpine AS builder
RUN apk add git ca-certificates
WORKDIR /app
COPY . .
RUN go mod download \
 && CGO_ENABLED=0 go build -o /app/main cmd/pricetopus/*

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/main pricetopus

ENV PRICETOPUS_EMAIL_SERVER=""
ENV PRICETOPUS_EMAIL_SERVER_PORT=""
ENV PRICETOPUS_EMAIL_USER=""
ENV PRICETOPUS_EMAIL_PASSWORD=""
ENV PRICETOPUS_EMAIL_TO=""
ENV PRICETOPUS_PRODUCT_URL=""
ENV PRICETOPUS_PRODUCT_PRICE=""

USER 12345
ENTRYPOINT [ "/pricetopus" ]
