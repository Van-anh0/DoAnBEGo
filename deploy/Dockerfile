FROM golang:1.18 as builder

WORKDIR /code
ADD go.mod .
ADD go.sum .
ADD . .
RUN go mod download


RUN go build -o server .

FROM ubuntu:20.04
WORKDIR /code
COPY --from=builder /code/server .
ADD resources ./resources

CMD "/code/server"
