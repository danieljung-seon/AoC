import os

dirname = os.path.dirname(__file__)
file = open(os.path.join(dirname, "02.txt"), "r")
data = file.readlines()

def solve_1():
    position = 0
    depth = 0

    for line in data:
        [command, value] = line.split()
        parsed_value = int(value)
        
        if command == "forward":
            position += parsed_value
        elif command == "down":
            depth += parsed_value
        elif command == "up":
            depth -= parsed_value

    return position * depth

def solve_2():
    position = 0
    aim = 0
    depth = 0

    for line in data:
        [command, value] = line.split()
        parsed_value = int(value)
        
        if command == "forward":
            position += parsed_value
            depth += aim * parsed_value
        elif command == "down":
            aim += parsed_value
        elif command == "up":
            aim -= parsed_value

    return position * depth

print('#1:', solve_1())
print('#1:', solve_2())
