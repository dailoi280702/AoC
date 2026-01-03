import math

boxes = []
with open("input.txt", "r") as f:
    for l in f:
        boxes.append([float(x) for x in l.strip().split(",")])

dists = []
for i1 in range(len(boxes)):
    for i2 in range(i1 + 1, len(boxes)):
        b1, b2 = boxes[i1], boxes[i2]
        d = ((b1[0] - b2[0]) ** 2 + (b1[1] - b2[1]) ** 2 + (b1[2] - b2[2]) ** 2) ** 0.5
        dists.append((d, i1, i2))

dists.sort()

# Union find
parent = list(range(len(boxes)))


def find(i):
    if parent[i] == i:
        return i
    parent[i] = find(parent[i])
    return parent[i]


num_circuits = len(boxes)
last_connection = None

for d, i1, i2 in dists:
    root1 = find(i1)
    root2 = find(i2)

    if root1 != root2:
        parent[root1] = root2
        num_circuits -= 1

        if num_circuits == 1:
            last_connection = (i1, i2)
            break

if last_connection:
    idx1, idx2 = last_connection
    x1 = int(boxes[idx1][0])
    x2 = int(boxes[idx2][0])
    print(x1 * x2)
