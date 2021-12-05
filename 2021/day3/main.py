def read_lines(path = "./input.txt"):
    with open(path) as f:
        lines = f.readlines()
        lines = [line.rstrip() for line in lines]
        return lines


class Solver:
    def __init__(self, method, line_length):
        self.method = method
        self.counter = []
        print("init with: {}".format(line_length + 1))
        for i in range(line_length):
            self.counter.append([0,0])

    def make_step(self, line):
        bits = list(line)
        for idx, bit in enumerate(bits):
            self.method(self, idx, bit)


def part_1_step(s: Solver, idx, bit):
    s.counter[idx][int(bit)]+=1


def part_2_step(s: Solver, idx, bita):
    pass


def solve(lines, method):
    s = Solver(method, len(lines[0]))
    for _, line in enumerate(lines):
        s.make_step(line)

    gamma_bits = []
    epsilon_bits = []
    for idx, bit_row in enumerate(s.counter):
        zeros, ones = bit_row

        gamma_bits.append('0' if zeros > ones else '1')
        epsilon_bits.append('0' if zeros < ones else '1')\
    
    return int(''.join(gamma_bits), 2) * int(''.join(epsilon_bits), 2)


if __name__ == '__main__':
    result1 = solve(read_lines(), part_1_step)
    print(result1)

    # result2 = solve(read_lines(), part_2_step)
    # print(result2)