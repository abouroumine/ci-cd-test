FROM golang:1.17-alpine as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o ci_cd_test

FROM alpine
RUN mkdir /app
WORKDIR /app
COPY --from=builder /build/ /app/

CMD ["./ci_cd_test"]
