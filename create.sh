#! /usr/bin/bash

NAME=$1
cd webook && mkdir $NAME
cd $NAME && mkdir _internal
touch main.go wire.go
cd _internal
mkdir domian repository service
cd repository && mkdir dao

