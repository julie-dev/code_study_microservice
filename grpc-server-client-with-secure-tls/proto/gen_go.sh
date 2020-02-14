#!/bin/bash

protoc ./echo.proto --go_out=plugins=grpc:./
