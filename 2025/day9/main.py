import itertools

points = []
with open("input.txt", "r") as f:
    for l in f:
        points.append([int(x) for x in l.strip().split(",")])

poly = [tuple(p) for p in points]


# Ray cast
def is_inside(x, y, polygon):
    n = len(polygon)
    inside = False

    for i in range(n):
        x1, y1 = polygon[i]
        x2, y2 = polygon[(i + 1) % n]

        # Check if point on edge
        if min(x1, x2) <= x <= max(x1, x2) and min(y1, y2) <= y <= max(y1, y2):
            if (x1 == x2 == x) or (y1 == y2 == y):
                return True

        if ((y1 > y) != (y2 > y)) and (x < (x2 - x1) * (y - y1) / (y2 - y1) + x1):
            inside = not inside

    return inside


# Points are sorted clock wise
edges = []
for i in range(len(poly)):
    p1, p2 = poly[i], poly[(i + 1) % len(poly)]
    edges.append(
        (min(p1[0], p2[0]), max(p1[0], p2[0]), min(p1[1], p2[1]), max(p1[1], p2[1]))
    )

max_area = 0
for p1, p2 in itertools.combinations(poly, 2):
    x1, y1 = p1
    x2, y2 = p2
    if x1 == x2 or y1 == y2:
        continue

    rx1, rx2 = min(x1, x2), max(x1, x2)
    ry1, ry2 = min(y1, y2), max(y1, y2)

    corners = [(rx1, ry1), (rx2, ry1), (rx1, ry2), (rx2, ry2)]
    if not all(is_inside(cx, cy, poly) for cx, cy in corners):
        continue

    cuts_edge = False
    for ex1, ex2, ey1, ey2 in edges:
        if not (ex2 <= rx1 or ex1 >= rx2 or ey2 <= ry1 or ey1 >= ry2):
            cuts_edge = True
            break

    if not cuts_edge:
        w = (rx2 - rx1) + 1
        h = (ry2 - ry1) + 1
        area = w * h

        if area > max_area:
            max_area = area

print(max_area)
