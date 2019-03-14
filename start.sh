#!/usr/bin/env bash

echo 'unzipping database data'
unzip pgdata.zip

echo 'running docker-compose'
docker-compose up -d