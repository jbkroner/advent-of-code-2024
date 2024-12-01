# grab input
with open("./solutions/01/input.txt", mode="r") as f:
    input = f.readlines()

# split into two lists
list_l = []
list_r = []
for line in input:
    l, r = line.split("   ")
    list_l.append(l)
    list_r.append(r)

# sort the lists
list_l.sort()
list_r.sort()

# calculate the distances
distances = []
for l, r in zip(list_l, list_r):
    distances.append(abs(int(l) - int(r)))

print(f"sum of distances: {sum(distances)}")
