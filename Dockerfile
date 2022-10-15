FROM golang:alpine3.16

WORKDIR /app
RUN \
  apk add --no-cache gcc libc-dev sqlite make

ENTRYPOINT [ "/bin/ash" ]
