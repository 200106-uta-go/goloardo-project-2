
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
RUN go build -o horoscope-ms ./horoscope-ms.go


FROM alpine:latest AS production
# Copies everything from builder to your new continer
COPY --from=builder /app .
EXPOSE 8080
CMD [ "./horoscope-ms" ]