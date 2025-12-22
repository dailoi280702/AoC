banks = []

with open("input.txt", "r") as f:
    for l in f:
        line = [int(c) for c in l.strip()]
        n = len(line)

        target_count = 12
        lv = 0
        current_pos = 0

        for i in range(target_count):
            search_range = line[current_pos : n + i + 1 - target_count]
            max_val = max(search_range)
            current_pos = search_range.index(max_val) + current_pos + 1
            lv += max_val * pow(10, (target_count - 1 - i))

        banks.append(lv)

print(banks)
print(sum(banks))
