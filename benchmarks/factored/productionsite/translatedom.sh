counter=1
for products in {2,3,4,5,6,7,8,9}; 
    do for refinements in {0,1,2,3,4}; 
        do 
            if [[ "$counter" -gt 9 ]]; then
                ./run.sh translate benchmarks/factored/productionsite/4-$products-$refinements-0/ experiments/factored/productionsite$counter/
            else
                ./run.sh translate benchmarks/factored/productionsite/4-$products-$refinements-0/ experiments/factored/productionsite0$counter/
            fi
            counter=$((counter+1))
    done; 
done
