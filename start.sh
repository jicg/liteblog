#!/bin/bash
set -e
cd /app/

if [ ! -d "/app/assert" ] ; then
    mkdir /app/assert
fi

if [ ! -f "/app/conf/app.conf" ] ; then
    touch /app/conf/app.conf
    sed -i 's/^appname.*/appname=liteblog/g' conf/app.conf
    sed -i 's/^httpport.*/httpport=8080/g' conf/app.conf
    sed -i 's/^runmode.*/runmode=prov/g' conf/app.conf

fi

liteblog