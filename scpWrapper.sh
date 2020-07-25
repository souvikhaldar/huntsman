#!/bin/bash
echo "hint: scpWrapper.sh binary username ip"
echo "binary: " $1
echo "username: " $2
echo "IP: " $3
scp -r static $2@$3:~
scp $1 $2@$3:~
