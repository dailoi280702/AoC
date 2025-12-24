fresh_ranges: list[list[int]] = []
ingredients: list[int] = []

with open("input.txt", "r") as file:
    for line in file:
        line = line.strip()

        if not line:
            continue

        if "-" in line:
            fresh_ranges.append([int(x) for x in line.split("-")])
            continue

        ingredients.append(int(line))


fresh_count: int = 0

for i in ingredients:
    for r in fresh_ranges:
        if i in range(r[0], r[1] + 1):
            fresh_count += 1
            break

print(fresh_count)

# Part 2
fresh_ranges.sort()
merged = [fresh_ranges[0]]

for current_start, current_end in fresh_ranges[1:]:
    last_start, last_end = merged[-1]

    if current_start <= last_end + 1:
        merged[-1][1] = max(last_end, current_end)
    else:
        merged.append([current_start, current_end])

print(sum(r[1] - r[0] + 1 for r in merged))
