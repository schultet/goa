#!/usr/bin/env python3
# -*- coding: utf-8 -*-

# TODO: rewrite this in go
# add cmd option --no-tests to not include '_test.go'-files, or list seperately
# add cmd option --sort to sort results based on #lines of code
# count TODOs

import os, os.path
import sys


def count(f):
    cnt, todos = 0, list()
    ignore = False
    with open(f, 'r') as fd:
        for i, line in enumerate(fd):
            line = line.strip()
            if "todo" in line.lower():
                todos.append((f, i, line)) 
            if line.startswith('/*'):
                ignore = True
            if ignore and '*/' in line:
                ignore = False
            elif not ignore and line and not line.startswith('//'):
                cnt += 1
    return cnt, todos


def countall():
    testing, base, maxcnt, todos = 0, 0, -1, list()
    filecounts = []
    for (path, dirs, files) in os.walk('.'):
        for f in files:
            if f.endswith('.go'):
                pf = os.path.join(path, f)
                cnt, tds = count(pf)
                todos += tds
                if cnt > maxcnt:
                    maxcnt, maxfile = cnt, pf
                if f.endswith('_test.go'):
                    testing += cnt
                elif f.endswith('.go'):
                    base += cnt
                filecounts.append((pf, cnt))
    return filecounts, base, testing, testing + base, maxcnt, maxfile, todos


if __name__ == '__main__':
    verbose = False
    listTodos = False
    sort = True
    if "-v" in sys.argv:
        verbose = True
    if "-todos" in sys.argv:
        listTodos = True
    fcnts, base, testing, total, maxcnt, maxfile, todos = countall()
    if sort:
        fcnts = sorted(fcnts, key=lambda x:x[1], reverse=True)
    if verbose:
        for fc in fcnts:
            print("{:<35}{:>10}".format(*fc))
    if listTodos:
        for td in todos:
            print("{:<40}{:<50}".format(td[0]+" ("+str(td[1])+")", td[2]))
    print("\nStats:")
    print("-"*47)
    print(("{:<10}{:>10}\n"*3)[:-1].format("base", base, "testing", testing,
        "total", total))
    print("{:<10}{:>10} {:<30}".format("max", maxcnt, "("+maxfile+")"))
    print("{:<10}{:>10}".format("TODOs", len(todos)))
    print("-"*47)

