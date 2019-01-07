#!/bin/bash
readonly name="MiniCloud"
readonly logs="info.log"
case $1 in
        start)
            nohup ./$name >> info.log 2>&1 &
            echo "server start..."
            sleep 1
        ;;
        stop)
            killall $name
            echo "server stop..."
		    sleep 1
		;;
		restart)
		    echo "server stop..."
		    killall $name
		    if [ "$?" != 0 ]; then
		        exit 3
		    fi
		    sleep 1
		    echo "server start...";
		    nohup ./$name >> info.log 2>&1 &
		    if [ "$?" != 0 ]; then
		        exit 4
		    fi
		    echo "server running..."
		    sleep 1
		;;
		build)
		    echo "compile start..."
		    go build -o $name main.go
		    echo "compile end..."
		;;
		*)
            echo "$0 {start|stop|restart|build}"
		    exit 5
		;;
esac

