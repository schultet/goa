#!/usr/bin/env python3
# -*- encoding: utf-8 -*-
import os
import argparse
import subprocess


def runfolder(*args, **kw):
    # collect files
    tasks = []
    for f in os.listdir(kw['task']):
        if f.endswith(".json"):
            tasks.append(os.path.join(kw['task'], f))
    tasks.sort(key=lambda x: int(x.split('/')[-1][:-5]))
    
    # assign agents
    port = int(kw['port'])
    host = "127.0.0.1"
    agents = "" 
    for i in range(len(tasks)):
        agents += "--agent \"{} {} {}\" ".format(i, host, str(port+i))
    
    problems = ""
    for i in range(len(tasks)):
        problems += "--problem {} ".format(tasks[i])
    
    tmpl = "{} {} {} {} & "
    planner = "go run src/main.go"
    if kw.get('threaded', None):
        cmd = tmpl.format(planner, problems, agents, ' '.join(args))
    else:
        cmd = ""
        for i, task in enumerate(tasks):
            s = tmpl.format(planner, "--problem " + tasks[i], agents, ' '.join(args))
            cmd += s
    cmd = cmd[:-2]
    os.system(cmd)


def on_search(*args, **kw):
    from configs import configs
    if 'config' in kw and kw['config'] in configs:
        c = configs[kw['config']]
        c[1] = '"' + kw['config'] + ' ' + c[1] + '"'
        c[3] = '"' + c[3] + '"'
        args += tuple(c)
        args = list(args)
        if '--threaded' in args:
            args.remove('--threaded')
            kw['threaded'] = True
    print('Config:', kw, args)
    print('Revision:', str(subprocess.check_output(['git', 'log', '-1', 'HEAD', 
        '--format="%ai %H"', '--date=local'])))
    print('Go version:', str(subprocess.check_output(['go', 'version'])))
    runfolder(*args, **kw)


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Run DIMAP planner')
    parser.add_argument('-t','--task', metavar='path/to/task/', help='factored planning task')
    parser.add_argument('-c','--config', metavar='CONFIG', help='use a predefined configuration')
    parser.add_argument('--threaded', action='store_true', help='run in threads')
    parser.add_argument('--port', default=3035, help='designated start port')
    parser.set_defaults(func=on_search)

    args, rest = parser.parse_known_args()
    args.func(*rest, **vars(args))
