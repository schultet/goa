#!/bin/sh
# Generate scenarios where x products have to be created.
for products in {4..10}; 
    # Every product being created has y properties. 
    # It takes y+1 steps to finish the product.
    do for properties in {0..6};
        do
        # scenario where 2 agents are involved in the production.
        #./gen.py 2 $products $properties 0 2-$products-$properties/;
        # scenario where 4 agents are involved in the production.
        ./gen.py 4 $products $properties $properties 4-$products-$properties-$properties/;
    done; 
done;
