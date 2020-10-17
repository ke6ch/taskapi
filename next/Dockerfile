# Using offical images
FROM node:12.16.1-alpine
LABEL maintainer="ke6ch <19749261+ke6ch@users.noreply.github.com>"

WORKDIR /usr/src/app

# デバッグ用にcurlを追加
RUN apk update && apk add curl

# APIのURLを設定する
ARG BASE_URL=http://localhost:80
ENV BASE_URL $BASE_URL

# デバッグ用に環境変数の値を出力
RUN echo $BASE_URL

# Nextjsの環境構築
COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build
RUN npx next telemetry disable

EXPOSE 3000

CMD ["npm", "start"]
