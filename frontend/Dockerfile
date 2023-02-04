# build step
FROM node:19-alpine3.16 as build
LABEL maintainer "Vitor Kretzer <vitor.kretzer@gamil.com>"

WORKDIR /app/frontend
COPY package.json ./
COPY package-lock.json ./
COPY . .
RUN yarn install
EXPOSE 8000
