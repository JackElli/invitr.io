# pulls from golang alpine
FROM golang:alpine as builder

# install git for dependencies
RUN apk update && apk add --no-cache git

# set workpath to anything with root
WORKDIR /app

#copy the users microservice (has to be local dir so users/)
COPY . .

WORKDIR /app/services/users
# get dependecies
# and build
RUN go get -d -v
RUN go build -o /users

FROM scratch

COPY --from=builder /users /users

EXPOSE 3200

ENTRYPOINT [ "/users" ]