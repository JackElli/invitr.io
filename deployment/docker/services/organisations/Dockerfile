# pulls from golang alpine
FROM golang:alpine as builder

# install git for dependencies
RUN apk update && apk add --no-cache git

# set workpath to anything with root
WORKDIR /app

#copy the organisations microservice (has to be local dir so organisations/)
COPY . .

WORKDIR /app/services/organisations
# get dependencies
# and build
RUN go get -d -v
RUN go build -o /organisations

FROM scratch

COPY --from=builder /organisations /organisations

EXPOSE 3203

ENTRYPOINT [ "/organisations" ]
