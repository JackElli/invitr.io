# pulls from golang alpine
FROM golang:alpine as builder

# install git for dependencies
RUN apk update && apk add --no-cache git

# set workpath to anything with root
WORKDIR /app

#copy the qr-codes microservice (has to be local dir so qr-codes/)
COPY . .

# get dependecies
# and build
RUN go get -d -v
RUN go build -o /qr-codes

FROM scratch

COPY --from=builder /qr-codes /qr-codes

EXPOSE 3200

ENTRYPOINT [ "/qr-codes" ]