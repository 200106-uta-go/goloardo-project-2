#build stage
FROM golang:alpine AS builder 
RUN mkdir /app 
ADD ./populate.go /app
WORKDIR /app
# Adds git to package manager
RUN apk add --no-cache git
# Always docker build relative to the main package directory
RUN go build -o populate ./populate.go

#production stage
FROM alpine:latest AS production
# Copies everything from builder to your new continer
COPY --from=builder /app .
ENTRYPOINT [ "./populate" ]
CMD [ "" ]