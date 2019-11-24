#!/usr/bin/env bash
protoc --proto_path=api/authentication --proto_path=third_party --go_out=plugins=grpc:proto-gen/authentication/rest_credentials rest_credentials.proto
protoc --proto_path=api/authentication --proto_path=third_party --grpc-gateway_out=logtostderr=true:proto-gen/authentication/rest_credentials rest_credentials.proto
protoc --proto_path=api/authentication --proto_path=third_party --swagger_out=logtostderr=true:api/authentication/swagger rest_credentials.proto
