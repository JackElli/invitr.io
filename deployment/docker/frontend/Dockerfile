FROM node:19.7-alpine as BUILD_IMAGE

ENV NODE_ENV development

WORKDIR /app

COPY ["package.json", "package-lock.json*", "./"]

RUN npm install
RUN npm install next

COPY . .

