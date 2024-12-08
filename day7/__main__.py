from collections.abc import Callable, Iterable
from dataclasses import dataclass
from math import log10
from operator import add, mul
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


type Opr = Callable[[int, int], int]


def _equation_opts(values: list[int], oprs: Iterable[Opr]) -> Iterable[int]:
    if len(values) == 1:
        yield from values
        return

    # We want to resolve from left to right, so calculate the left hand first
    *left_hand, right_hand = values
    for option in _equation_opts(left_hand, oprs):
        for opr in oprs:
            yield opr(option, right_hand)


def _calculate(equations: list[Equation], oprs: list[Opr]) -> int:
    return sum(
        equation.target
        for equation in equations
        if any(answer == equation.target for answer in _equation_opts(equation.values, oprs))
    )


def concat(a: int, b: int) -> int:
    return 10 ** int(log10(b) + 1) * a + b


def main() -> None:
    input_lines = read_input_lines(Path("input.txt"))

    equations = [Equation.from_str(line) for line in input_lines]

    print(f"Part 1 result: {_calculate(equations, [add, mul])}")
    print(f"Part 2 result: {_calculate(equations, [add, mul, concat])}")


if __name__ == "__main__":
    main()
