#!/bin/sh
SCRIPT_DIR=$(dirname "$0")
BASE_DIR=$SCRIPT_DIR/../../..
cd $BASE_DIR
BASE_DIR=$(pwd)
SCRIPT_DIR=benchmarks/factored/productionsite
EXPERIMENTS_DIR=$BASE_DIR"/experiments"
# Translate scenarios where x products have to be created.
for products in {2..10}; 
    # Every product being created has y properties. 
    # It takes y+1 steps to finish the product.
    do for properties in {0..4};
        do
        # scenario where 2 agents are involved in the production.
        #SCENARIO=2-$products-$properties
        #./run.sh translate $SCRIPT_DIR/$SCENARIO $EXPERIMENTS_DIR/productionsite-$SCENARIO;
        # scenario where 4 agents are involved in the production.
        SCENARIO=4-$products-$properties
        ./translate.sh $SCRIPT_DIR/$SCENARIO $EXPERIMENTS_DIR/productionsite-$SCENARIO;
    done; 
done;
