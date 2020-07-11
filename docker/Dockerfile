FROM alpine:latest

RUN apk update && apk add --no-cache libc6-compat
COPY ./go-chat /go/src/

CMD ["/go/src/go-chat"]