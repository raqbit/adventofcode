from collections.abc import Iterable
from dataclasses import dataclass
from pathlib import Path
from typing import Self

from shared import read_input_lines


@dataclass(kw_only=True, frozen=True)
class Equation:
    target: int
    values: list[int]

    @classmethod
    def from_str(cls, str_: str) -> Self:
        target, rest = str_.split(": ")
        return cls(
            target=int(target),
            values=[int(val) for val in rest.split(" ")],
        )


def _equation_opts(values: list[int]) -> Iterable[int]:
    if len(values) == 1:
        yield from values
        return

    # We want to resolve from left to right, so calculate the left hand first
    *left_hand, right_hand = values
    for opt in _equation_opts(left_hand):
        yield opt * right_hand
        yield opt + right_hand


def _part1(equations: list[Equation]) -> int:
    return sum(
        equation.target
        for equation in equations
        if any(answer == equation.target for answer in _equation_opts(equation.values))
    )


def main() -> None:
    input_lines = read_input_lines(Path("input.txt"))

    equations = [Equation.from_str(line) for line in input_lines]

    print(f"Part 1 result: {_part1(equations)}")
    # print(f"Part 1 result: {part2_total}")


if __name__ == "__main__":
    main()
