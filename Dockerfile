# Use the offical golang image to create a binary.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.16beta1-buster as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN go build -mod=readonly -v -o server


# Now build the react app
# Use the official lightweight Node.js 12 image.
# https://hub.docker.com/_/node
FROM node:12-slim as nodebuilder
WORKDIR /tmp

# Copy application dependency manifests to the container image.
# A wildcard is used to ensure copying both package.json AND package-lock.json (when available).
# Copying this first prevents re-running npm install on every code change.
COPY package*.json ./

# Install production dependencies.
# If you add a package-lock.json, speed your build by switching to 'npm ci'.
# RUN npm ci --only=production
RUN npm ci --only=production

COPY . ./
RUN npm run build
RUN ls /tmp/build

# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /server
COPY --from=nodebuilder /tmp/build /build
# COPY ./_build /_build
# COPY ./data /data

RUN ls /
# RUN ls /_build

# Run the web service on container startup.
CMD ["/server"]




# # WORKING CONFIG
# FROM golang
# WORKDIR /go/src/app
# COPY . .

# RUN go build -v -o /app .

# ENV PORT 8080

# CMD ["/app"]
