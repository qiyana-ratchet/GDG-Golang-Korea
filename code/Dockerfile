#FROM golang:1.12-alpine3.9 AS build-env
FROM golang:latest

ADD pm-exporter /bin/
ADD please_parse_this.xml .

WORKDIR /app
#WORKDIR /collector

COPY pm-exporter /app
#COPY collector /app/collector
#COPY data_parse.xml /app
