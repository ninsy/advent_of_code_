def read_lines(path = "./input.txt"):
    with open(path) as f:
        lines = f.readlines()
        lines = [line.rstrip() for line in lines]
        return lines


class Submarine:
    def __init__(self, method):
        self.x = 0
        self.y = 0
        self.aim = 0
        self.method = method

    def make_move(self, line):
        [direction, step] = line.split(" ")
        step = int(step)
        return self.method(self, direction, step)


def part_1_move(s: Submarine, direction, step):
    match direction:
        case 'up':
            s.y -= step
        case 'down':
            s.y += step
        case 'forward':
            s.x += step


def part_2_move(s: Submarine, direction, step):
    match direction:
        case 'up':
            s.aim -= step
        case 'down':
            s.aim += step
        case 'forward':
            s.x += step
            s.y += s.aim * step


def solve(lines, method):
    s = Submarine(method)
    for _, line in enumerate(lines):
        s.make_move(line)
    
    return s.x * s.y


if __name__ == '__main__':
    result1 = solve(read_lines(), part_1_move)
    print(result1)

    result2 = solve(read_lines(), part_2_move)
    print(result2)