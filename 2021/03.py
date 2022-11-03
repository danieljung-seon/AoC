import os

dirname = os.path.dirname(__file__)
file = open(os.path.join(dirname, "03.txt"), "r")
data = file.readlines()

# helpers

def bits_to_number(bit_list):
    out = 0
    for bit in bit_list:
        out = (out << 1) | bit

    return out

def partition(data, fn, idx):
    left, right = [], []

    for d in data:
        left.append(d) if fn(d, idx) else right.append(d)

    return left, right

# task 1

def solve_1():
    nums = [0] * len(list(data[0].strip()))

    for line in data:
        line_nums = list(line.strip())

        for i, n in enumerate(line_nums):
            nums[i] += 1 if n == "1" else -1

    gamma_rate = bits_to_number(list(map(lambda x: 1 if x > 0 else 0, nums)))
    epsilon_rate = bits_to_number(list(map(lambda x: 0 if x > 0 else 1, nums)))
    return gamma_rate * epsilon_rate

# task 2

def get_rating(data, get_new_data_by, idx = 0):
    if len(data) == 1:
        return bits_to_number(map(
            lambda x: 1 if x == "1" else 0,
            list(data[0].strip()),
        ))

    left, right = partition(data, lambda x, i: list(x.strip())[i] == "1", idx)
    new_data = get_new_data_by(left, right)

    return get_rating(new_data, get_new_data_by, idx + 1)

def get_oxygen_generator_rating(data):
    return get_rating(data, lambda left, right: left if len(left) >= len(right) else right)

def get_co2_scrubber_rating(data):
    return get_rating(data, lambda left, right: right if len(left) >= len(right) else left)

def solve_2():
    o2_rating = get_oxygen_generator_rating(data)
    co2_rating = get_co2_scrubber_rating(data)

    return o2_rating * co2_rating

print('#1:', solve_1())
print('#2:', solve_2())
