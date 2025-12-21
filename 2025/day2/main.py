invalid_numbers = []


def is_repeating(val):
    s = str(val)
    return s in (s + s)[1:-1]


with open("input.txt", "r") as f:
    content = f.read().replace("\n", "").replace(" ", "")
    ranges = content.split(",")

    for item in ranges:

        start, end = map(int, item.split("-"))
        for num in range(start, end + 1):
            if is_repeating(num):
                invalid_numbers.append(num)

print(sum(invalid_numbers))
