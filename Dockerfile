# base image
FROM golang:1.25.1 AS builder
# set working directory
WORKDIR /app
# copy go module and dependencies
COPY go.mod go.sum ./
RUN go mod download
# copy source code
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
# build the application
RUN CGO_ENABLED=0 go build -o nutech-api ./app
# use minimal base image for final deployment
FROM gcr.io/distroless/static
# set working directory in the container
WORKDIR /
# copy the builder binary from the builder stage
COPY --from=builder /app/nutech-api . 
ENV PORT=8080
# EXPOSE 8080
# run the application
CMD ["/nutech-api"]