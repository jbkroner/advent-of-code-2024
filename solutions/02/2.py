# https://adventofcode.com/2024/day/2
# open the input and build a lists of ints to process
input = "./solutions/02/input.txt"
# input = "./solutions/02/sample.txt"

with open(input, mode="r") as f:
    input = f.readlines()
    reports = [
        [int(x) for x in line.strip().split(" ")] 
        for line in input
    ]

num_valid = 0
for report in reports:
    all_increasing = True
    all_decreasing = True
    all_diff_ok = True

    for i in range(len(report) - 1):
        prev, curr = report[i], report[i + 1]

        all_decreasing = all_decreasing and (prev > curr)

        all_increasing = all_increasing and (prev < curr)

        all_diff_ok = all_diff_ok and (1 <= abs(prev - curr) <= 3)

    safe = (all_increasing ^ all_decreasing) and all_diff_ok

    if safe:
        num_valid += 1

print(f"valid reports: {num_valid}")












