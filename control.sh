#!/bin/sh
CURRENT_DIR=$(cd $(dirname $0); pwd)
#echo $CURRENT_DIR
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


function doc() {
    tmpfile="./log/.godoc"
    PID=$(cat $tmpfile)
    echo "godoc  pid is "$PID
    if [ "$param" == "stop" ]; then
        if [ ! $PID_EXIST ];then
            echo the process $PID is not exist

        else
            echo the process $PID exist
            echo "kill godoc"
            kill -9 $PID
            
        fi
        exit 1
    else
        echo "godoc controller start"
    fi
    
    

    PID_EXIST=$(ps aux | awk '{print $2}'| grep -w $PID)
    if [ ! $PID_EXIST ];then
        echo the process $PID is not exist
        echo "start godoc"
        nohup godoc -http=:8001 > ./log/godoc.log &
        echo $! > $tmpfile
    else
        echo the process $PID exist
        echo "restart godoc"
        kill -9 $PID
        doc
    fi
    echo "http://localhost:8001/pkg/gDAG/"
}

if [ "$1" == "build" ]; then
    build
elif [ "$1" == "doc" ]; then
    doc
elif [ "$1" == "clean" ]; then
    clean
else
    echo "Invalid function name"
fi