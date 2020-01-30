FROM golang:1.13
RUN mkdir -p discover_services
ADD . discover_services
WORKDIR discover_services
RUN apt-get update
ENV PORT 8080
EXPOSE 8080
RUN go build -o discover_services .
CMD ["./discover_services"]

