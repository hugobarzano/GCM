#!/usr/bin/env bash

docker container prune --force
docker volume prune --force
docker network prune --force
docker image prune -a
docker system prune --volumes --force && docker rmi $(docker image ls -aq) #DO all in one single command
*/10  * * * *   root    docker system prune --volumes --force && docker rmi $(docker image ls -aq)
