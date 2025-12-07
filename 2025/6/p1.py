lines = [line.strip().split() for line in open(0)]
cols = list(zip(*lines))
print(zip(*lines))

sum = 0
for *nums, op in cols:
    sum += eval(op.join(nums))
print(sum)
