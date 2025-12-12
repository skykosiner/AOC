# from functools import cache

items = {}

for line in open(0):
    src, dsts = line.strip().split(": ")
    items[src] = dsts.split()

# @cache
def count(src, dst):
    if src == dst: return 1
    return sum(count(x, dst) for x in items.get(src, []))

print(count("you", "out"))

print(
    count("svr", "dac") * count("dac", "fft") * count("fft", "out") \
  + count("svr", "fft") * count("fft", "dac") * count("dac", "out")
)
