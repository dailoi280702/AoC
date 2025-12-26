import math

sheet = []
operators = []
res = 0

with open("input.txt", "r") as file:
    for line in file:
        if "+" in line:
            operators = line.split()
            continue

        sheet.append([int(x) for x in line.split()])

transposed = [[row[i] for row in sheet] for i in range(len(sheet[0]))]

for i in range(len(transposed)):
    if operators[i] == "+":
        res += sum(transposed[i])
        continue

    res += math.prod(transposed[i])


for line in sheet:
    print(line)
print(operators)
print(res)
