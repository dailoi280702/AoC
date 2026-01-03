import itertools

points = []

with open("input.txt", "r") as f:
    for l in f:
        points.append([int(x) for x in l.strip().split(",")])

areas = [
    abs(p1[0] - p2[0] + 1) * abs(p1[1] - p2[1] + 1)
    for p1, p2 in itertools.combinations(points, 2)
]

print(max(areas))
