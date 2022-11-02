import os

dirname = os.path.dirname(__file__)
file = open(os.path.join(dirname, "03.txt"), "r")
data = file.readlines()

def bits_to_number(bit_list):
    out = 0
    for bit in bit_list:
        out = (out << 1) | bit

    return out

def solve_1():
    nums = [0] * len(list(data[0].strip()))

    for line in data:
        line_nums = list(line.strip())

        for i, n in enumerate(line_nums):
            nums[i] += 1 if n == "1" else -1

    gamma_rate = bits_to_number(list(map(lambda x: 1 if x > 0 else 0, nums)))
    epsilon_rate = bits_to_number(list(map(lambda x: 0 if x > 0 else 1, nums)))
    return gamma_rate * epsilon_rate

print('#1:', solve_1())
