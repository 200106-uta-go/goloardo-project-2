#build stage
FROM golang:alpine AS builder 
RUN mkdir /app 
ADD . /app
WORKDIR /app
# Adds git to package manager
RUN apk add --no-cache git
# Go get dependency 
RUN go get -d
# Always docker build relative to the main package directory
RUN go build -o horoscoped ./horoscoped-web.go


FROM alpine:latest AS production
# Copies everything from builder to your new continer
RUN mkdir /app 
WORKDIR /app
COPY --from=builder /app /app
EXPOSE 8080
ENTRYPOINT [ "./horoscoped" ]
CMD [ "" ]