#!/bin/bash

go version
go build -o ./tmp/$JOB_NAME ./cmd/server
