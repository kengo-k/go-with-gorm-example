#!/bin/sh
docker run -it --rm --name sample-go-with-gorm -v "${PWD}/src/:/app" sample-go-with-gorm:v1.0.0
