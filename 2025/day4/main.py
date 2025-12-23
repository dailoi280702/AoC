grid = []
d = [(0, 1), (1, 0), (-1, 0), (0, -1), (1, 1), (1, -1), (-1, 1), (-1, -1)]

with open("input.txt", "r") as file:
    for line in file:
        grid.append([1 if c == "@" else 0 for c in line.strip()])

num_toiletroll = sum(cell != 0 for row in grid for cell in row)

for i in range(len(grid)):
    for j in range(len(grid[i])):
        if grid[i][j] == 0:
            continue

        for y, x in d:
            nx, ny = x + i, y + j

            if nx < 0 or ny < 0 or ny >= len(grid) or nx >= len(grid[j]):
                continue

            if grid[nx][ny] == 0:
                continue

            grid[nx][ny] += 1


def remove(i, j):
    if grid[i][j] == 0 or grid[i][j] > 4:
        return

    grid[i][j] = 0

    for y, x in d:
        nx, ny = x + i, y + j

        if nx < 0 or ny < 0 or ny >= len(grid) or nx >= len(grid[j]):
            continue

        if grid[nx][ny] == 0:
            continue

        grid[nx][ny] -= 1

        remove(nx, ny)


for i in range(len(grid)):
    for j in range(len(grid[i])):
        remove(i, j)


num_toiletroll2 = sum(cell != 0 for row in grid for cell in row)

print(num_toiletroll - num_toiletroll2)
