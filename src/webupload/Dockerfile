FROM golang:latest 
EXPOSE 9090
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app
VOLUME ["/tmp"]
RUN go build main.go
CMD ["/app/main"]