import math
from itertools import pairwise
from pathlib import PurePath

from shared import read_input


def _is_safe_report(report: list[int]) -> bool:
    sign = None
    for x, y in pairwise(report):
        dist_sign = math.copysign(1, x - y)

        if not sign:
            sign = dist_sign
        elif sign != dist_sign:
            return False

        if not 1 <= abs(x - y) <= 3:
            return False

    return True


def _is_safe_report_with_dampener(report: list[int]) -> bool:
    if _is_safe_report(report):
        return True

    # Naive approach: Will just try again when removing one of the levels
    for i, _ in enumerate(report):
        if _is_safe_report(report[:i] + report[i + 1:]):
            return True

    return False


def main():
    reports = [
        [int(level) for level in line.split(' ')]
        for line in read_input(PurePath('input.txt'))
    ]

    print(f"Part 1 result: {sum(_is_safe_report(report) for report in reports)}")

    print(f"Part 2 result: {sum(_is_safe_report_with_dampener(report) for report in reports)}")


if __name__ == "__main__":
    main()
