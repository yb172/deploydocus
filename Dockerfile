# Build container. Here we build golang app
# To build the app we need go, so we base this container
# on a heavy-weight golang image
FROM golang:alpine AS build-go
ARG version
# Download deps (do it as a separate step to let Docker cache downloaded deps)
RUN apk add git
WORKDIR /src
COPY go.mod .
RUN go mod download -x
# Copy sources and run build
ADD . .
RUN go build -ldflags "-X main.Version=${version}" -o server

# Run container. Here we run app
# We don't need anything additional to run the app
# so we base this container on a lightweight alpine image
FROM alpine
WORKDIR /app
# Copy built app from build container
COPY --from=build-go /src/server /app/server
# timezone 
RUN apk add --no-cache tzdata
# Start the server
ENTRYPOINT ./server start --env $ENV
