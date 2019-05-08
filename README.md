# README #

Note: This is work-in-progress.
Current version: 1.0

### Content ###

This repository contains the source files of a privacy preserving distributed
multi-agent planning system called Goa (= Groups of agents), also referred to as
Goaplan.  


### Setup ###

1. clone the repository
2. install the required go packages
    - go get -u
3. install vendor code
    - cd vendor
    - make nanomsg
    - make nanomsg4py
4. export nanomsg library path for python

    ``` shell
    export LD_LIBRARY_PATH="$LD_LIBRARY_PATH":vendor/nanomsg/build/lib/ && \
    export PYTHONPATH="$PYTHONPATH":vendor/nanomsg4py
    ```


### Running the planner ###

To solve a distributed multi-agent planning problem, execute the following steps

1. compile the problem description files (factored MA-PDDL) into the required
   json format

``` shell
    ./translate.sh <src-folder> <trg-folder>
```

2. run the planner with search.sh, either specifying all configurations:

    ``` shell
        ./search.sh \ 
        -t <path/of/task/files/> \
        --search "<search-options>" \
        --heuristic [ff|add|addbf|blind|gc] \
        [--threaded --macros]
    ```

    or using predefined configuration profiles:

    ``` shell
        ./search.sh \ 
        -t <folder-of-task-files> \
        -c <configuration-profile> \
        [--threaded --macros]
    ```

    configuration profiles are defined in driver/configs.py

Example:
``` shell
./search.sh -t ./experiments/taxi01/ \
    --search '"myconfig -s dmt-gus -h ff -l 100 -c 1.41"' \ 
    --heuristic ff \
    --threaded
```



### Running tests ###
``` shell
    go test -v <package>
```

or, from within the `src` folder:

``` shell
    go test ./...
```

### Contact ###

* Tim Schulte <schultet@informatik.uni-freiburg.de>
