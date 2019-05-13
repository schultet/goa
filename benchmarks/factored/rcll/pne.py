import sys
import re

with open(sys.argv[1], 'r') as f, open('out' + sys.argv[1], 'w') as o:
    for line in f:
        print(re.sub(r'\d{1,3}\.\d+', lambda m: str(int(float(m.group(0)) * 10)),
            line), file=o, end='')
