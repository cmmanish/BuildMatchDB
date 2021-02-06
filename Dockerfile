FROM golang
ADD . /go/src/simple
RUN go install simple
ENTRYPOINT /go/bin/simple
EXPOSE 8080

FROM php:7.4-cli
COPY . /usr/src/myapp
WORKDIR /usr/src/myapp
CMD [ "php", "./your-script.php" ]