from pathlib import Path

from shared import read_input_lines


def _count_positions(grid: list[list[str]], x_pos: int, y_pos: int) -> int:
    position_count = 0

    dx = 0
    dy = -1  # Start by going up

    while True:
        new_x = x_pos + dx
        new_y = y_pos + dy

        if new_y < 0 or new_y >= len(grid) or new_x < 0 or new_x >= len(grid[new_y]):
            # Outside grid, stop moving
            break

        if grid[new_y][new_x] == "#":
            if dx == 0 and dy == -1:
                # If up, go right
                dx = 1
                dy = 0
            elif dx == 1 and dy == 0:
                # If right, go down
                dx = 0
                dy = 1
            elif dx == 0 and dy == 1:
                # If down, go left
                dx = -1
                dy = 0
            elif dx == -1 and dy == 0:
                # If left, go up
                dx = 0
                dy = -1
            continue

        if grid[new_y][new_x] != "X":
            position_count += 1
            grid[new_y][new_x] = "X"

        x_pos, y_pos = new_x, new_y

    return position_count


def _find_start_pos(grid: list[list[str]]) -> tuple[int, int]:
    for y, row in enumerate(grid):
        for x, col in enumerate(row):
            if col == "^":
                return x, y

    raise ValueError("Start position not found")


def main() -> None:
    input_lines = read_input_lines(Path("input.txt"))

    grid = [list(line) for line in input_lines]

    start_pos = _find_start_pos(grid)

    print(f"Part 1 result: {_count_positions(grid, *start_pos)}")
    # print(f"Part 1 result: {part2_total}")


if __name__ == "__main__":
    main()
