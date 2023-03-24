FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY /competency-constructor/swagger/v1 .


CMD ["./competency-constructor"]