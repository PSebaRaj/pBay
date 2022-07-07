#!/bin/sh

# build / exit if any compilation errors
(cd users && go build) || exit 1
(cd products && go build) || exit 1
(cd messages && go build) || exit 1
(cd payments && go build) || exit 1
(cd shipping && go build) || exit 1
(cd advertisement && go build) || exit 1


# start services
echo '***' Starting the users microservice '***'
./users/users &
sleep 1

echo '***' Starting the products microservice '***'
./products/products &
sleep 1

echo '***' Starting the messages microservice '***'
./messages/messages &
sleep 1

echo '***' Starting the payments microservice '***'
./payments/payments &
sleep 1

echo '***' Starting the shipping microservice '***'
./shipping/shipping &
sleep 1

echo '***' Starting the advertisements microservice '***'
./advertisement/advertisement

# 5 backgrounded processes
wait
wait
wait
wait
wait

