lines = []
with open("./input.test") as f:
    lines = f.read().strip().split("\n")


sides = [[], []]

for line in lines:
    left, right = map(int, line.split("   "))
    sides[0].append(left)
    sides[1].append(right)

sides[0].sort()
sides[1].sort()

sum = 0

for item in sides:
    print(item)
    print(item[0])
