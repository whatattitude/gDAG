#!/bin/sh
CURRENT_DIR=$(cd $(dirname $0); pwd)
echo $CURRENT_DIR
action=$1
param=$2

function build() {
    echo $param
    if [ "$param" == "" ]; then
        echo "build all app"
        cd $CURRENT_DIR"/app"
        for file in $(ls -d *); do
            cd $file
            pwd
            go build -o $CURRENT_DIR"/bin/"$file"-prod"
            cd ..
        done 
    else 
        cd $CURRENT_DIR"/app/"$param
        pwd
        go build -o $CURRENT_DIR"/bin/"$param"-prod"
        cd $CURRENT_DIR
        
    fi
}

function clean() {
    echo $CURRENT_DIR"/bin"
    cd $CURRENT_DIR"/bin" && rm -rf ./*-prod
}



if [ "$1" == "build" ]; then
    build
elif [ "$1" == "clean" ]; then
    clean
else
    echo "Invalid function name"
fi