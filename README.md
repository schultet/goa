# README #

Note: This is work-in-progress.
Current version: 0.1

### Content ###

This repository contains the source files of *GOA*, a privacy preserving
distributed multi-agent planning system. GOA has two execution modes:
distributed and centralized.


### Setup ###
0. install Go if you haven't already. (See: [https://golang.org/doc/install](https://golang.org/doc/install).)
1. clone the repository:

    ``` shell
    mkdir -p $GOPATH/src/github.com/schultet/ \
    && cd $GOPATH/src/github.com/schultet \
    && git clone https://github.com/schultet/goa.git
    ```
2. install the required go packages:

    ``` shell
    cd $GOPATH/src/github.com/schultet/goa/pkg/
    go get [-u]
    ```
3. install the required python modules using pip3:

    ``` shell
    pip3 install -r requirements.txt
    ```


### Running the planner (centralized) ###

To solve a distributed multi-agent planning problem, execute the following
steps:

0. compile the problem description files (factored MA-PDDL) into the required
   json format

    ``` shell
        ./translate.sh <src-folder> <trg-folder>
    ```

1. run the planner with search.sh, either specifying all configuration options:

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

    Note: configuration profiles are defined in scripts/configs.py

Example:

    ``` shell
    ./search.sh -t ./experiments/taxi01/ \
        --search '"myconfig -s dmt-gus -h ff -l 100 -c 1.41"' \ 
        --heuristic ff \
        --threaded
    ```

### Running the planner (distributed) ###

To run the planner in a distributed fashion, run the following commands in
different terminals/shells (one for each agent):

0. Compile the MA-PDDL factors into a local task file for each agent. This is a
   distributed process in which all planning agents participate. E.g.:

    ``` shell
    # Agent 0:

    python3 ./translate/translate.py \
    benchmarks/factored/productionsite/2-4-2-1/domain-0.pddl \
    benchmarks/factored/productionsite/2-4-2-1/problem-0.pddl \
    --agent-url tcp://127.0.0.1:3035 \
    --agent-url tcp://127.0.0.1:3036 \
    --agent-id 0 \
    --output test/0.json \
    --json
    ```
    
    ``` shell
    # Agent 1 (different terminal):

    python3 ./translate/translate.py \
    benchmarks/factored/productionsite/2-4-2-1/domain-1.pddl \
    benchmarks/factored/productionsite/2-4-2-1/problem-1.pddl \
    --agent-url tcp://127.0.0.1:3035 \
    --agent-url tcp://127.0.0.1:3036 \
    --agent-id 1 \
    --output test/1.json \
    --json
    ```

1. Run the planner using the compiled task-files (.json):

    ``` shell
    # Agent 0:

    go run cmd/distributed/main.go \
    --problem test/0.json \
    --agent "0 127.0.0.1 3035" \
    --agent "1 127.0.0.1 3036" \
    -s "mafs -s mafs-g  -h ff -l 1" \
    --heuristic "ff" \
    --planlimit 1
    ```

    ``` shell
    # Agent 1 (different terminal):

    go run cmd/distributed/main.go \
    --problem test/1.json \
    --agent "0 127.0.0.1 3035" \
    --agent "1 127.0.0.1 3036" \
    -s "mafs -s mafs-g  -h ff -l 1" \
    --heuristic "ff" \
    --planlimit 1
    ```

### Docker ###

To use our environment, you can build and run a GOA image:

0. Build the docker image:

    ``` shell
    docker build -t goa .
    ```

1. Run the container:

    ``` shell
    docker run -d -t --name my_goa goa
    ```

2. Execute planner commands, e.g.:

    ``` shell
    # Translate MA-PDDL files
    docker exec -it my_goa /bin/bash translate.sh <path/to/pddl/> <path/to/compiled/task>

    # Search for a plan
    docker exec -it my_goa /bin/bash search.sh -t <path/to/compiled/task> -c mafs
    ```

Note: you can also use the Docker image to plan in a distributed fashion. In
this case, you need to `docker run` two GOA containers, each representing a
single agent. In this case, you have to map host-container ports accordingly,
such that one agent uses port 3035 and the other port 3036 (as an example).

### Running tests ###
``` shell
    go test -v <package>
```
   or

``` shell
    go test ./pkg/...
```

### Contact ###

* Tim Schulte <schultet@informatik.uni-freiburg.de>

### References ###

* will be added soon
