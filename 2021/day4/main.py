def read_lines(path = "./input.txt"):
    with open(path) as f:
        lines = f.readlines()
        lines = [line.rstrip() for line in lines]
        return lines

class Number:
    def __init__(self, value):
        self.value = int(value)
        self.marked = False

class Board:
    def __init__(self, board_lines):
        # 5x5 board
        self.board = []
        for line in board_lines:
            number_arr = list(filter(None, line.split(' ')))
            row = []
            for number in number_arr:
                row.append(Number(number))
            self.board.append(row)

    def transpose(self):
        transposed = []
        for i in range(len(self.board)):
            transposed_row = []
            for j in range(len(self.board[i])):
                transposed_row.append(self.board[j][i])
            transposed.append(transposed_row)
        return transposed

    def is_winning(self):
        # check columns
        for row in self.board:
            if all(number.marked for number in row):
                return True
        # check rows
        t = self.transpose()
        for row in t:
            if all(number.marked for number in row):
                return True
        return False

class Game:
    def __init__(self) -> None:
        self.idx = 0
        self.drawn_numbers = []
        self.boards = []

    def next(self):
        curr_num = int(self.drawn_numbers[self.idx])
        for board in self.boards:
            for row in board.board:
                for number in row:
                    if number.value == curr_num:
                        # TODO: no more checks needed, exit loops
                        number.marked = True

            if board.is_winning():
                unmarked_sum = 0
                for row in board.board:
                    for number in row:
                        if number.marked == False:
                            unmarked_sum += number.value
                return unmarked_sum * curr_num
        self.idx += 1
        return None
        

def parse_input(lines):
    g = Game()
    board_buffer = []
    for idx, line in enumerate(lines):
        if idx == 0:
            # Reading initial numbers
            g.drawn_numbers = line.split(',')
            continue
        if len(line) == 0:
            continue
        board_buffer.append(line)
        if len(board_buffer) == 5:
            b = Board(board_buffer.copy())
            g.boards.append(b)
            board_buffer = []
    return g


def part1(lines):
    game = parse_input(lines)

    while game.idx < len(game.drawn_numbers) - 1:
        winner = game.next()
        if winner != None:
            return winner


if __name__ == '__main__':
    result = part1(read_lines())
    print(result)

    # result = part2(read_lines())
    # print(result)