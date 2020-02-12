# FROM golang
# ADD . .
# COPY . . 
# EXPOSE 8080
# CMD ["go", "build", "cmd/horoscope-ms/horoscope-ms.go"]

# build stage
# FROM golang:alpine AS builder
# ENV GOPATH /home/daniaestrada/go/src/
# WORKDIR /home/daniaestrada/go/src/github.com/200106-uta-go/goloardo-project-2
# COPY cmd/horoscope-ms/horoscope-ms.go /usr/local/app
# COPY internal/gethoroscope/gethoroscope.go /usr/local/app

# COPY $GOPATH/github.com/200106-uta-go/goloardo-project-2/cmd/horoscope-ms/horoscope-ms.go $GOPATH/src/github.com/200106-uta-go/goloardo-project-2/cmd/horoscope-ms/horoscope-ms.go
# COPY $GOPATH/github.com/200106-uta-go/goloardo-project-2/internal/gethoroscope/gethoroscope.go $GOPATH/src/github.com/200106-uta-go/goloardo-project-2/internal/gethoroscope/gethoroscope.go
# RUN apk add --no-cache git
# RUN go get -d -v ./...
# RUN go build cmd/horoscope-ms/horoscope-ms.go#

# final stage
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# COPY --from=builder /go/src/app /app
# EXPOSE 8080
# ENTRYPOINT ["./horoscope-ms"]

#-----------------
#build stage
FROM golang:alpine AS builder 
RUN mkdir /app 
ADD . /app
WORKDIR /app
RUN apk add --no-cache git
RUN go get -d
#RUN go get github.com/200106-uta-go/goloardo-project-2/pkg/gethoroscope
RUN go build -o horoscope-ms ./horoscope-ms.go


FROM alpine:latest AS production
COPY --from=builder /app .
EXPOSE 8080
CMD [ "./horoscope-ms" ]