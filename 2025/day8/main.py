import math

boxes = []
with open("input.txt", "r") as f:
    for l in f:
        boxes.append(l.strip().split(","))

dists = [
    (
        (
            (float(boxes[i1][0]) - float(boxes[i2][0])) ** 2
            + (float(boxes[i1][1]) - float(boxes[i2][1])) ** 2
            + (float(boxes[i1][2]) - float(boxes[i2][2])) ** 2
        )
        ** 0.5,
        i1,
        i2,
    )
    for i1 in range(len(boxes))
    for i2 in range(i1 + 1, len(boxes))
]
dists.sort(key=lambda x: x[0])
dists = dists[:1000]

parent = list(range(len(boxes)))


def find(i):
    if parent[i] == i:
        return i
    parent[i] = find(parent[i])
    return parent[i]


def union(i, j):
    root_i = find(i)
    root_j = find(j)
    if root_i != root_j:
        parent[root_i] = root_j


for _, i1, i2 in dists:
    union(i1, i2)

circuits = {}
for i in range(len(boxes)):
    root = find(i)
    if root not in circuits:
        circuits[root] = []
    circuits[root].append(i)


for c in circuits.values():
    print(c)

circuit_lens = [len(c) for c in circuits.values()]
circuit_lens.sort(reverse=True)

print(math.prod(circuit_lens[:3]))
