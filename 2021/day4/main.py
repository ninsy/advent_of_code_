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

    def mark_numbers(self, board, curr_num):
        for row in board:
            for number in row:
                if number.value == curr_num:
                    number.marked = True
                    return

    def next(self, part):
        curr_num = int(self.drawn_numbers[self.idx])
        for board in self.boards:
            self.mark_numbers(board.board, curr_num)
        # TODO: how this can be refactored? seems higly insufficient to loop twice at boards. better way to clear?
        for board in self.boards:
            if part == 1 and self.check_part1(board) or part == 2 and self.check_part2(board):
                unmarked_sum = 0
                for row in board.board:
                    for number in row:
                        if number.marked == False:
                            unmarked_sum += number.value
                return unmarked_sum * curr_num
        self.idx += 1
        return None
        
    def check_part1(self, board):
        return board.is_winning()
                
    def check_part2(self, board):
        if len(self.boards) == 1:
            return True
        if board.is_winning():
            for idx, b in enumerate(self.boards):
                if id(b.board) == id(board.board):
                    del self.boards[idx]
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


def solve(lines, part):
    game = parse_input(lines)

    while game.idx < len(game.drawn_numbers) - 1:
        winner = game.next(part)
        if winner != None:
            return winner


if __name__ == '__main__':
    result1 = solve(read_lines(), 1)
    print(result1)

    result2 = solve(read_lines(), 2)
    print(result2)
