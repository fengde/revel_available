#!/bin/bash

WORKSPACE=$(cd $(dirname $0)/; pwd)
cd $WORKSPACE

app=revel_available
src=revel_available

function build() {
    dest="/tmp/$app"
    revel build $src $dest
    if [ $? -ne 0 ]; then
        exit $?
    fi

    mkdir -p $app/src/$src/app
    cp -fr $dest/$src $app 
    cp -fr $dest/src/github.com $app/src
    cp -fr control.sh $app/src/
    cp -fr $dest/src/$src/conf $dest/src/$src/public $app/src/$src
    cp -fr $dest/src/$src/app/views $app/src/$src/app
    rm -fr $dest
}

function pack() {
    build
    file_dir="$app"
    echo "tar $file_dir.tar.gz <= $file_dir"
    tar -zcf $file_dir.tar.gz $file_dir
    rm -fr $file_dir
}

function help() {
    echo "$0 build|pack"
}

if [ "$1" == "" ]; then
    help
elif [ "$1" == "build" ];then
    build
elif [ "$1" == "pack" ];then
    pack
else
    help
fi