#!/usr/bin/env python3
# -*- encoding: utf-8 -*-
import os
import argparse
import subprocess
from glob import glob


def translatefolder(src, trg, **kw):
    python = kw.get("python", "python3")
    translate = kw.get("translate", "./translate/translate.py")
    port = int(kw.get("port", 3035))
    host = kw.get("host", "127.0.0.1")
    
    # create directories
    if not os.path.exists(trg):
        os.makedirs(trg)
    
    # collect files
    domains, problems = [], []
    for f in os.listdir(src):
        if "domain" in f and f.endswith(".pddl"):
            domains.append(os.path.join(src, f))
        elif "problem" in f and f.endswith(".pddl"):
            problems.append(os.path.join(src, f))
    domains.sort()
    problems.sort()
    
    # assign agents
    agents = []
    for i in range(len(domains)):
        agents.append("tcp://{}:{}".format(host, str(port+i)))
    
    # create command
    tmpl = ("{} {} {} {} --agent-url " + " --agent-url ".join(agents) +
            " --agent-id {} --output {} --json")
    cmd = ""
    for i, d in enumerate(domains):
        s = tmpl.format(python, translate, d, problems[i], i,
                os.path.join(trg,str(i)+'.json')) + ' & '
        print(s)
        cmd += s
    cmd = cmd[:-2]
    
    os.system(cmd)


def translateall(src='benchmarks/factored/', trg='benchmarks/compiled/', **kw):
    files_src = glob(src + "*/*/")
    files_trg = [os.path.join(trg, *f.split('/')[2:]) for f in files_src]

    port = 3035
    shift = 100
    errors = []
    for s, t in zip(files_src, files_trg):
        try:
            print("translating " + s + " to " + t + " port: " + str(port))
            translatefolder(s, t, port=port)
        except Exception as e:
            errors += [e]
        port += shift
    for i, error in enumerate(errors):
        print("ERR %d: %s: %s" % (i, type(error), error))


def on_translate(*args, **kw):
    if kw['all']:
        translateall(**kw)
    else:
        translatefolder(**kw)


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Run DIMAP planner')
    parser.add_argument('src', help='path to folder containing src task')
    parser.add_argument('trg', help='destination path')
    parser.add_argument(
        '--port', 
        default=3035,
        help='the port (default: 3035)'
    )
    parser.add_argument(
        '--all', 
        help='translate all domains of given folder',
        action='store_true'
    )
    parser.set_defaults(func=on_translate)

    args, rest = parser.parse_known_args()
    args.func(*rest, **vars(args))
