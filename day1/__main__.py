import functools
from collections import Counter
from pathlib import PurePath

from shared import read_input


def main():
    lines = list(read_input(PurePath('input.txt')))

    nums_left = []
    nums_right = []

    for left, right in map(functools.partial(str.split, sep="   "), lines):
        nums_left.append(int(left))
        nums_right.append(int(right))

    nums_left.sort()
    nums_right.sort()

    print(f"Part 1 result: {sum(abs(left - right) for left, right in zip(nums_left, nums_right))}")

    counts = Counter(nums_right)

    print(F"Part 2 result: {sum(num * counts[num] for num in nums_left)}")


if __name__ == "__main__":
    main()
