from functools import cache

m = []

with open("input.txt", "r") as f:
    for l in f:
        m.append([c for c in l.strip()])

nc = len(m[0])
nr = len(m)


@cache
def travel(x, y):
    if x < 0 or y < 0 or x >= nc or y >= nr:
        return 1

    if m[y][x] == "|":
        return 0

    if m[y][x] == "^":
        return travel(x - 1, y) + travel(x + 1, y)

    m[y][x] = "|"
    p = travel(x, y + 1)
    m[y][x] = "."

    return p


print(travel(m[0].index("S"), 1))
