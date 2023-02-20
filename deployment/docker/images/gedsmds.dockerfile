FROM golang:1.20.1-alpine3.17 as builder

RUN apk add --no-cache \
	bash \
	make

WORKDIR /gedsmds
ADD . /gedsmds/
COPY ./env.secret /gedsmds/env

RUN make tidy
RUN make build-mds

FROM alpine:latest

ENV GEDSMDS_SERVER_PORT=50001

WORKDIR /gedsmds
COPY --from=builder /gedsmds/build/linux/gedsmds .

EXPOSE $GEDSMDS_SERVER_PORT
CMD ["./gedsmds"]
