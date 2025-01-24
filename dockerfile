# build stage
FROM golang:1.23-alpine3.19 AS build-env
RUN apk update && apk upgrade && apk add --no-cache bash git openssh g++ glib-static

ARG VERSION
ARG COMMIT

WORKDIR /hr-service
COPY . /hr-service

ENV GO111MODULE=on

# RUN ls -l

#RUN go mod vendor

RUN go build -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT}" -o app .

# final stage
FROM alpine
RUN apk update && apk upgrade && apk add --no-cache libgcc libstdc++

WORKDIR /app

COPY --from=build-env hr-service/app /app/
COPY --from=build-env hr-service/environment/. /app/environment/.
ENTRYPOINT ["./app"]
