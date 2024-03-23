# pulls from golang alpine
FROM golang:alpine as builder

# install git for dependencies
RUN apk update && apk add --no-cache git

# set workpath to anything with root
WORKDIR /app

#copy the invites microservice (has to be local dir so invites/)
COPY . .

WORKDIR /app/services/invites
# get dependecies
# and build
RUN go get -d -v
RUN go build -o /invites

FROM scratch

COPY --from=builder /invites /invites

EXPOSE 3200

ENTRYPOINT [ "/invites" ]
