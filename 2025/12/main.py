import re

print(sum((w//3)*(h//3) >= sum(c)
    for l in open(0).read().splitlines()[30:]
    for w,h,*c in [map(int, re.findall(r'\d+', l))]))
