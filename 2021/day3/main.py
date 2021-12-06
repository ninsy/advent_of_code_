def read_lines(path = "./input.txt"):
    with open(path) as f:
        lines = f.readlines()
        lines = [line.rstrip() for line in lines]
        return lines


class Counter:
    def __init__(self, lines):
        self.counter = []
        for i in range(len(lines[0])):
            self.counter.append([0,0])

        for _, line in enumerate(lines):
            bits = list(line)
            for idx, bit in enumerate(bits):
                self.counter[idx][int(bit)]+=1


def solve_part1(lines):
    s = Counter(lines)

    gamma_bits = []
    epsilon_bits = []
    for _, bit_row in enumerate(s.counter):
        zeros, ones = bit_row

        gamma_bits.append('0' if zeros > ones else '1')
        epsilon_bits.append('0' if zeros < ones else '1')
    
    return int(''.join(gamma_bits), 2) * int(''.join(epsilon_bits), 2)


def part2_solver(lines, inc_mode):
    for idx in range(len(lines[0])):
        s = Counter(lines)
    
        zeros, ones = s.counter[idx]
        if len(lines) == 1:
            return lines
        if ones > zeros:
            # keep ones
            lines = list(filter(lambda line: line[idx] == '1' if inc_mode else line[idx] == '0', lines))
        elif zeros > ones:
            # keep zeros
            lines = list(filter(lambda line: line[idx] == '0' if inc_mode else line[idx] == '1', lines))
        elif zeros == ones:
            # keep with one in that pos
            for line in lines:
                if inc_mode and line[idx] == '1':
                    return [line]
                elif not inc_mode and line[idx] == '0':
                    return [line]


def solve_part2(all_lines):
    oxygen_rating = part2_solver(all_lines.copy(), True)[0]
    co2_scruber_rating = part2_solver(all_lines.copy(), False)[0]

    return int(''.join(oxygen_rating), 2) * int(''.join(co2_scruber_rating), 2)


if __name__ == '__main__':
    result1 = solve_part1(read_lines())
    print(result1)

    result2 = solve_part2(read_lines())
    print(result2)
