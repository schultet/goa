#!/usr/bin/env python3
#! -*- encoding: utf8 -*-
""" Aggregates plans from log files in a csv file. """
import re
import sys
import json
import csv
from glob import glob


def tq(ls):
    n = len(ls)
    for i, x in enumerate(ls):
        print('{}/{}'.format(i, n))
        yield x


def progress(ls):
    """If tqdm is installed, return wrapped iterable to display progress bar."""
    try:
        assert False
        from tqdm import tqdm
        return tqdm(ls)
    except:
        return tq(ls)


def main(folder='../restest/'):
    """ Parses all log files in a directory and writes plans in a csv file. """
    files = glob("{}/*.res".format(folder))
    print(files)
    results = []
    for filename in progress(files):
        domprob = re.search(r'__(?P<Domain>[a-zA-Z0-9-]+)__-__(?P<Problem>[a-zA-Z0-9-]+)__', filename)
        with open(filename, 'r') as f:
            try:
                data = f.read()
                config = re.search(r'Config: (?P<Config>.+)$', data, flags=re.S)
                summaries = re.finditer(r'{\n.*?}\n', data, flags=re.S)
                for summary in summaries:
                    result = json.loads(summary.group(0))
                    result.update(domprob.groupdict())
                    if config:
                        result['Config'] = config.group(1).split()[0]
                    else:
                        result['Config'] = 'NaN'
                    results.append(result)
            except Exception as e:
                print(e)
    return results


if __name__ == '__main__':
    JSON_SUMMARIES = main(folder=sys.argv[1])
    import time
    FILENAME = time.strftime("%b-%d-%Y-%H%M%S-results.csv").lower()
    with open(FILENAME, 'w') as f:
        HEADER = list(JSON_SUMMARIES[0].keys())
        # The individual steps of the plan can be looked up in the log file
        HEADER.remove('Steps')
        WRITER = csv.DictWriter(f, fieldnames=HEADER)
        WRITER.writeheader()
        for json_summary in JSON_SUMMARIES:
            # Remove entries in column "Steps"
            steps_present = json_summary.pop('Steps', None)
            if steps_present is None:
                print(f'No "Steps" defined in plan {json_summary["Plan"]}')
            WRITER.writerow(json_summary)
