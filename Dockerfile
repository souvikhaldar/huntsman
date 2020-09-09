# builder image
FROM golang:1.13-alpine3.11 as builder
RUN mkdir /huntsman
ADD . /huntsman
WORKDIR /huntsman
RUN go build -o huntsman .

# the final image
FROM alpine
COPY --from=builder /huntsman/huntsman .
ENTRYPOINT [ "./huntsman" ]
