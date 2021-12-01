from functools import reduce

def read_lines(path = "./input.txt"):
    with open(path) as f:
        lines = f.readlines()
        lines = [line.rstrip() for line in lines]
        return lines


def part1(lines):
    prev = None
    increase_count = 0
    for _, line in enumerate(lines):
        value = int(line)
        if prev != None and value > prev:
            increase_count+=1
        prev = value

    return increase_count


def part2(lines):
    increase_count = 0
    idx = 0
    prev_sum = None
    for idx, _ in enumerate(lines):
        window = lines[idx:idx+3]
        if len(window) < 3:
            break
        sum_curr = reduce(lambda x,y: int(x)+int(y), window)
        if prev_sum != None and sum_curr > prev_sum:
            increase_count+=1
        prev_sum = sum_curr
    return increase_count
        

if __name__ == '__main__':
    result = part1(read_lines())
    print(result)

    result = part2(read_lines())
    print(result)