#!/bin/bash
set -e
cd /app/

if [ ! -d "/app/assert" ] ; then
    mkdir /app/assert
fi

liteblog