#!/bin/sh

cd $HOME/Documents/HouseGuard-EnvironmentManager/src

git pull

go clean

go build

if [ -f exeEnvironmentManager ];
then
    echo "EVM File found"
    if [ -f $HOME/Documents/Temp/exeEnvironmentManager ];
    then
        echo "EVM old removed"
        rm -f $HOME/Documents/Temp/exeEnvironmentManager
    fi
    mv exeEnvironmentManager $HOME/Documents/Temp/exeEnvironmentManager
fi