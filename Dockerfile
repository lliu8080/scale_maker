ARG go_image_version=1.21
ARG almalinux_version=9
FROM golang:$go_image_version AS build

# Building the binary of the App
WORKDIR /go/src/scale_maker

# Copy all the Code and stuff to compile everything
COPY . .

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

RUN go test ./...

# If need to use venders file
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .
# RUN mkdir -p bin && \
#	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o bin/app ./main.go

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

# Moving the binary to the 'final Image' to make it smaller
FROM almalinux:$almalinux_version

WORKDIR /app
RUN dnf -y update
RUN dnf -y install stress-ng iperf3 fio

# Create the `public` dir and copy all the assets into it
RUN mkdir ./assets
COPY ./assets ./assets
COPY ./pkg/api/templates ./templates

# Copy app binary from build image.
COPY --from=build /go/src/scale_maker/app .

# Exposes port 3000 because our program listens on that port
EXPOSE 3000

CMD ["./app"]
