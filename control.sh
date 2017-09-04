#!/bin/bash

WORKSPACE=$(cd $(dirname $0)/; pwd)
cd $WORKSPACE

app="../revel_available -importPath revel_available -srcPath ./ -runMode prod"
pidfile=./pid/revel_available.pid
logfile=/dev/null

function check_pid() {
    if [ -f $pidfile ];then
        pid=`cat $pidfile`
        if [ -n $pid ]; then
            running=`ps -p $pid|grep -v "PID TTY" |wc -l`
            return $running
        fi
    fi
    return 0
}

function start() {
    check_pid
    running=$?
    if [ $running -gt 0 ];then
        echo -n "$app now is running already, pid="
        cat $pidfile
        return 1
    fi

    nohup ./$app > $logfile 2>&1 &
    echo $! > $pidfile
    echo "$app started..., pid=$!"
}

function stop() {
    pid=`cat $pidfile`
    kill $pid
    echo "$app stoped..."
}

function restart() {
    stop
    sleep 1
    start
}

function status() {
    check_pid
    running=$?
    if [ $running -gt 0 ];then
        echo started
    else
        echo stoped
    fi
}


function help() {
    echo "$0 start|stop|restart|status"
}

if [ "$1" == "" ]; then
    help
elif [ "$1" == "stop" ];then
    stop
elif [ "$1" == "start" ];then
    start
elif [ "$1" == "restart" ];then
    restart
elif [ "$1" == "status" ];then
    status
else
    help
fi