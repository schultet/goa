#! /usr/bin/env python3
# -*- coding: utf-8 -*-
"""This module contains functions to create a job file containing all possible
combinations of configurations and benchmarks. The resulting file <out>
can then be submitted with

> qsub <out>
> qgki <out>  # gki members only

Authors:
    Tim Schulte <schultet@cs.uni-freiburg.de>

"""

import sys
import os
import argparse


header = """#!/bin/bash
#SBATCH -D /home/schultet/dimap
#SBATCH -e joblog/std.err.%j
#SBATCH -o joblog/std.out.%j
#SBATCH -J {jobname}
#SBATCH -p {queue}
#SBATCH --time={timelimit}
#SBATCH -a 1-{jobs}
"""

def createjobs(logfolder, configs, tasks):
    jobs = []
    # TODO: if more than 10 agents participate we need a greater shift
    startport, shift = 3035, 10 
    for config in configs:
        for task in tasks:
            expfolder = os.path.abspath(task)
            task = "__{}__-__{}__".format(os.path.basename(os.path.dirname(expfolder)),
                                  os.path.basename(expfolder))
            jobs += [{
                "taskid": len(jobs)+1,
                "expfolder": expfolder,
                "config": config,
                "logfolder": os.path.abspath(logfolder),
                "task": task,
                "port": len(jobs) * shift + startport,
            }]
    return jobs


def writejobs(jobs, outf, queue, timeout=None, memout=None, dryrun=False): 
    jobtmpl = (
        "\nif [ {taskid} -eq $SLURM_ARRAY_TASK_ID ]; then\n"
        "    ./search.sh -t {expfolder} -c {config} --port {port}"
        "    --logfile {logfile}\n"
        "    exit $?\n"
        "fi\n"
    )
    with open(outf, 'w') as f:
        if dryrun:
            outstream = sys.stdout
        else:
            outstream = f
        print(header.format(
                queue=queue,
                timelimit=timeout if timeout else 0,
                #memout=memout if memout else 0,
                jobs=len(jobs),
                jobname='BALLERN',
            ), file=outstream)
        for job in jobs:
            logfile = '{logfolder}/{config}-{task}-$SLURM_ARRAY_JOB_ID.$SLURM_ARRAY_TASK_ID.res'.format(**job)
            job['logfile'] = logfile
            print(jobtmpl.format(**job), file=outstream)


def parse_args():
    parser = argparse.ArgumentParser(description='create jobs for grid')
    parser.add_argument('-t', '--tasks', nargs='+', metavar='T ...', 
        default=[],
        help='List of folders containing PDDL files, e.g. "rovers/p01 rovers/p02 ..."')
    parser.add_argument('-a', '--all', action='store_true')
    parser.add_argument('-q', '--queue', 
        default='test_core.q', 
        choices=['cpu_ivy', 'gki_cpu-ivy', 'test_cpu-ivy'],
        help='Select the queue. Queues with "_core" suffix use a single core.')
    parser.add_argument('--out', default='job.q', 
            help='the name of the output file, def: job.q')
    parser.add_argument('--logfolder', default='log', 
            help='path to the folder for the log files, def: log/')
    parser.add_argument('-c', '--configs', nargs='+', metavar='C',
            help='list of configs', default=[])
    parser.add_argument('--timeout', default=None, metavar='h:mm:s',
            help=('jobs timeout, "0:30:0" sets the timeout to 30 minutes'
                  'default value is "None", i.e. no timeout.'))
    parser.add_argument('--memout', default=None, metavar='<mem>[G|M]',
            help=('jobs memory cap, "<mem>G" ("<mem>M") sets the cap to <mem>'
                  ' (an integer) many GByte (MByte).' 
                  ' Default value is "None", i.e. no memory limit.'))
    return parser.parse_args(sys.argv[1:])


def create_folders(*folders):
    for f in folders:
        p = os.path.join(os.getcwd(), f)
        if not os.path.exists(p):
            os.makedirs(p)
            print("created folder: %s" % p)


if __name__ == '__main__':
    args = parse_args()

    if args.all:
        from files import experiments
        tasks = []
        for d in experiments:
            tasks.append(d)
            for p in range(len(experiments[d])):
                tasks.append(str(p))
    else:
        tasks = args.tasks

    from configs import configs
    for conf in args.configs:
        if conf not in configs:
            print('invalid config {}'.format(conf))
            raise KeyError

    # create `joblog` and `logfolder` folder if it doesn't exist
    create_folders('joblog', args.logfolder)

    jobs = createjobs(args.logfolder, args.configs, tasks)
    writejobs(jobs, args.out, args.queue, args.timeout, args.memout)
    print('jobfile {} written (to cwd)'.format(args.out))
