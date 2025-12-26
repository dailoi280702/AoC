import math

sheet = []
operators = []

with open("input.txt", "r") as file:
    for line in file:
        if "+" in line:
            operators = line.split()
            continue

        for i, c in enumerate(line):
            if len(sheet) == i:
                sheet.append(int(c if c.isdigit() else 0))
                continue

            if not c.isdigit():
                continue

            sheet[i] = sheet[i] * 10 + int(c)


operands = []
curr = []

for num in sheet:
    if num == 0:
        operands.append(curr.copy())
        curr = []
        continue

    curr.append(num)


res = 0

for i in range(len(operands)):
    if operators[i] == "+":
        res += sum(operands[i])
        continue

    res += math.prod(operands[i])

print(res)
