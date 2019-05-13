#!/bin/sh
for products in {4..10}; 
    do for properties in {0..6};
        do
        ./translate.sh \
        benchmarks/factored/productionsite/4-$products-$properties-$properties/ \
        benchmarks/compiled/productionsite/4-$products-$properties-$properties/;
    done; 
done;

for products in {2..10}; 
    do for properties in {0..4};
        do
        ./translate.sh \
        benchmarks/factored/productionsite/4-$products-$properties-0/ \
        benchmarks/compiled/productionsite/4-$products-$properties-0/;
    done; 
done;
