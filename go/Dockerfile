# Using offical images
FROM golang:latest
LABEL maintainer="k <19749261+keb0@users.noreply.github.com>"

WORKDIR /go/src/app
COPY . .

# Add library
RUN go get github.com/labstack/echo \
  && go get github.com/labstack/echo/middleware \
  && go get github.com/go-sql-driver/mysql

CMD ["go", "run", "main.go"]
