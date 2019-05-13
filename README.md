# README #

Note: This is work-in-progress.
Current version: 1.0

### Content ###

This repository contains the source files of *Goa*, a privacy preserving
distributed multi-agent planning system.


### Setup ###
0. install Go if you haven't already. See [https://golang.org/doc/install](https://golang.org/doc/install)
1. clone the repository. Either using the go tool `go get`:
        
    ``` shell
    go get github.com/schultet/goa/src
    ```

    or manually using `git clone`, making sure it is located in the $GOPATH correctly. I.e.:

    ``` shell
    mkdir -p $GOPATH/src/github.com/schultet/
    cd $GOPATH/src/github.com/schultet
    git clone https://github.com/schultet/goa.git
    ```
    


2. install the required go packages

    ``` shell
    cd $GOPATH/src/github.com/schultet/goa/src/
    go get [-u]
    ```
3. install vendor code

    ``` shell
    cd $GOPATH/src/github.com/schultet/goa/vendor/
    make nanomsg
    make nanomsg4py
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

or

``` shell
    go test ./src/...
```

### Contact ###

* Tim Schulte <schultet@informatik.uni-freiburg.de>
