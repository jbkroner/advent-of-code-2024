# grab input
with open("./01/input.txt", mode="r") as f:
    input = f.readlines()

# split into two lists
list_l = []
list_r = []
for line in input:
    l, r = line.split("   ")
    list_l.append(l.strip())
    list_r.append(r.strip())

# build a frequency map of list_r
list_r_freq = {}
for r in list_r:
    if r not in list_r_freq:
        list_r_freq[r] = 1
    else:
        list_r_freq[r] += 1

# calculate the "similiarity score"
# add up each number in the left list after multiplying it by
# the number of times that number appears in the right list
sim_list = []
for l in list_l:
    sim_list.append(int(l) * list_r_freq.get(l, 0))

print(f"similiarity score: {sum(sim_list)}")
