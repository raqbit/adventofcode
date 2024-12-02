import functools
from collections import Counter
from pathlib import Path

from shared import read_input


def main() -> None:
    lines = list(read_input(Path("input.txt")))

    nums_left = []
    nums_right = []

    for left, right in map(functools.partial(str.split, sep="   "), lines):
        nums_left.append(int(left))
        nums_right.append(int(right))

    nums_left.sort()
    nums_right.sort()

    print(
        f"Part 1 result: {sum(
            abs(left - right)
            for left, right in zip(nums_left, nums_right, strict=True)
        )}"
    )

    counts = Counter(nums_right)

    print(f"Part 2 result: {sum(num * counts[num] for num in nums_left)}")


if __name__ == "__main__":
    main()
