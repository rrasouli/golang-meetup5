#!/bin/bash
while true
do
tr -dc 'A-Za-z0-9!"#$%&'\''()*+-./:;<=>?@[\]^_`{|}~' </dev/urandom | head -c 13
done


