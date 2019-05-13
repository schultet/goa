#!/bin/sh
# Generate scenarios where x products have to be created.
for products in {2..10}; 
    # Every product being created has y properties. 
    # It takes y+1 steps to finish the product.
    do for properties in {0..4};
        do
        # scenario where 4 agents are involved in the production.
        ./gen.py 4 $products $properties 0 4-$products-$properties-0/;
    done; 
done;
