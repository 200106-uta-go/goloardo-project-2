#build stage
FROM golang:alpine AS builder 
RUN mkdir /app 
ADD . /app
#WORKDIR /app
RUN mkdir /app/badger
#Adds basics package, this includes gcc which was used in badger's compression algorithm.
RUN apk add build-base
#The --no-cache option allows to not cache the index locally, which is useful for keeping containers small.
#This is executing apk update in the beginning and rm -rf /var/cache/apk/* in the end.
RUN apk add --no-cache git && apk update
#Obtain all the dependency to run the functions.
RUN go get -d
RUN go build -o dbserver /app/dbserver.go

#execute stage
FROM alpine:latest AS production
COPY --from=builder . .
EXPOSE 8081
ENTRYPOINT [ "./dbserver" ]
CMD [ "" ]
#CMD ["/bin/sh"]