sum = 0

with open("./validNums") as f:
    lines = f.readlines()

    for line in lines:
        nums = line.split(" ")
        x = int(nums[0]) * int(nums[1])

        sum += x

print(sum)
