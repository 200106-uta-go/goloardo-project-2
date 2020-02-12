FROM golang

ADD . .

COPY . . 

# EXPOSE 8080

CMD ["go", "run", "main.go"]