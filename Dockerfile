#FROM golang:1.12-alpine3.9 AS build-env
FROM golang:latest

ADD pm-exporter /bin/
ADD collector/parse_this_1.xml .
ADD collector/parse_this_2.xml .
ADD collector/parse_this_3.xml .
ADD collector/parse_this_4.xml .
ADD collector/parse_this_5.xml .

WORKDIR /app
#WORKDIR /collector

COPY pm-exporter /app
#COPY collector /app/collector
#COPY data_parse.xml /app
