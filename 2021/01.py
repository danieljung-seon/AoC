import os

dirname = os.path.dirname(__file__)
file = open(os.path.join(dirname, "01.txt"), "r")
data = file.readlines()

def solve(steps_between):
    value_increased_cnt = 0

    for idx, val in enumerate(data):
        if idx < len(data) - steps_between:
            if int(data[idx + steps_between]) - int(val) > 0:
                value_increased_cnt += 1

    return value_increased_cnt

print('#1:', solve(1))
print('#2:', solve(3))
