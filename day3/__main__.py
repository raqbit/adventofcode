import re
from pathlib import Path

from shared import read_input


def _calc_program(memory: str) -> int:
    matches = [match.groups() for match in re.finditer(r"mul\((\d+),(\d+)\)", memory)]
    return sum(int(left) * int(right) for left, right in matches)


def _part2(memory: str) -> int:
    total = 0
    enabled = True

    for match in re.finditer(r"don't\(\)|do\(\)|mul\((\d+),(\d+)\)", memory):
        matched_op = match.group(0)
        if matched_op == "do()":
            enabled = True
        elif matched_op == "don't()":
            enabled = False
        elif enabled:
            total += int(match.group(1)) * int(match.group(2))

    return total


def main() -> None:
    memory = read_input(Path("input.txt"))

    print(f"Part 1 result: {_calc_program(memory)}")
    print(f"Part 2 result: {_part2(memory)}")


if __name__ == "__main__":
    main()
