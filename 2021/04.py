import os

dirname = os.path.dirname(__file__)
file = open(os.path.join(dirname, "04.txt"), "r")

class Board:
    def __init__(self):
        self.not_found = set()
        self.nums = {}
        self.rows = [0, 0, 0, 0, 0]
        self.cols = [0, 0, 0, 0, 0]

    def add(self, row_data, row_idx):
        y = row_idx
        for x, num in enumerate(row_data):
            self.nums[num] = (x, y)
            self.not_found.add(int(num))
    
    def process_input(self, input):
        self.not_found.discard(int(input))
        try:
            (x, y) = self.nums[input]
            self.cols[x] += 1
            self.rows[y] += 1
        except KeyError:
            pass

    def is_bingo(self):
        for col in self.cols:
            if (col == 5):
                return True

        for row in self.rows:
            if (row == 5):
                return True

        return False
        
def sum_set(nums):
    return sum(list(nums))

def parse_data(file):
    inputs = []
    boards = []
    y = 0

    data = file.readlines()
    for idx, line in enumerate(data):
        if idx == 0:
            inputs = line.strip().split(',')
            continue

        if line == '\n':
            y = -1
            boards.append(Board())
        else:
            y += 1
            nums = line.strip().split()
            boards[-1].add(nums, y)

    return inputs, boards

def solve_1():
    inputs, boards = parse_data(file)
    for input in inputs:
        for board in boards:
            board.process_input(input)
            if (board.is_bingo()):
                return int(input) * sum_set(board.not_found)

print('#1:', solve_1())
