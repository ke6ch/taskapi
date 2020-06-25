# Using offical images
FROM golang:latest
LABEL maintainer="ke6ch <19749261+ke6ch@users.noreply.github.com>"

WORKDIR /go/src/app
COPY . .

# Add library
RUN go get github.com/labstack/echo \
  && go get github.com/labstack/echo/middleware \
  && go get github.com/go-sql-driver/mysql

EXPOSE 1323

CMD ["go", "run", "main.go"]
