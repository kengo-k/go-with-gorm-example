FROM golang:alpine3.16

WORKDIR /app
RUN \
  apk add --no-cache gcc libc-dev make

ENTRYPOINT [ "/bin/ash" ]