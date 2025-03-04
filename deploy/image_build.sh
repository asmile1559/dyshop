#!/bin/bash

cd $(dirname $0)

image() {
    image_name="jiny14/dyshop-$1"
    go build -o bin/$1 .
    docker build -t $image_name .
}

cd ../app

cd auth
image auth

cd ../user
image user

cd ../product
image product

cd ../cart
image cart

cd ../order
image order

cd ../checkout
image checkout

cd ../payment
image payment

cd ../frontend
image frontend

cd $(dirname $0)
