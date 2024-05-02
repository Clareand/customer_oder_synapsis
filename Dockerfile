#################################
# Multi Stage 
# STEP 1 build executable binary
#################################

FROM golang:1.18.1-alpine as builder

RUN apk update && apk add --no-cache git


WORKDIR /system-conf
COPY . ./

RUN go mod tidy
RUN go mod vendor
RUN cd cmd/api && CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o customer_order_synapsis .

#COPY ./.env.sample ./cmd/api/.env


#############################
# STEP 2 build a small image
#############################
# FROM scratch
FROM quay.io/giantswarm/golang:1.18.1-alpine3.15

COPY --from=builder /customer_order_synapsis/cmd/api/ /customer_order_synapsis/cmd/api/
COPY ./certificate ./certificate
COPY ./cmd/api/certificate ./cmd/api/certificate

EXPOSE 8080
ENTRYPOINT ["/customer_order_synapsis/cmd/api/customer_order_synapsis"]
CMD ["./customer_order_synapsis"]
