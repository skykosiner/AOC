import sys
import re

d = sys.stdin.read().split("\n")
nums = [[num for num in x.split("-")] for x in d]
print(nums)
