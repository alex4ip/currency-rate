FROM golang:1.22 AS build
WORKDIR /go/src
COPY main.go .
COPY go.mod .
COPY go ./go

ENV CGO_ENABLED=1
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o app .

#FROM scratch AS runtime
#COPY --from=build /go/src/app ./
EXPOSE 8080/tcp
ENTRYPOINT ["/go/src/app"]
