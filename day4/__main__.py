from pathlib import Path

from shared import read_input


def _pos_or_none(idx: int) -> int | None:
    """
    If the index is negative, we don't want Python's negative
    slicing wraparound behavior, but instead slice until the beginning,
    which we can accomplish by using None.
    """
    return idx if idx >= 0 else None


def _count_word_occurrences(crossword: str, word: str) -> int:
    row_len = crossword.find("\n") + 1

    total = 0
    for pos in range(len(crossword)):
        if crossword[pos] != word[0]:
            continue

        # Horizontal
        if crossword[pos : pos + len(word)] == word:
            total += 1

        # Horizontal reverse
        if crossword[pos : _pos_or_none(pos - len(word)) : -1] == word:
            total += 1

        # Vertical down
        if crossword[pos : pos + row_len * len(word) : row_len] == word:
            total += 1

        # Vertical up
        if crossword[pos : _pos_or_none(pos - (row_len * len(word))) : -row_len] == word:
            total += 1

        # Diagonal down right
        if crossword[pos : pos + row_len * len(word) : row_len + 1] == word:
            total += 1

        # Diagonal down left
        if crossword[pos : pos + row_len * (len(word) - 1) : row_len - 1] == word:
            total += 1

        # Diagonal up left
        if crossword[pos : _pos_or_none(pos - (row_len * len(word))) : -(row_len + 1)] == word:
            total += 1

        # Diagonal up right
        if crossword[pos : _pos_or_none(pos - (row_len * (len(word) - 1))) : -(row_len - 1)] == word:
            total += 1

    return total


def _count_cross_mas(input_str: str) -> int:
    crossword = [list(line) for line in input_str.splitlines()]

    total = 0
    for y in range(1, len(crossword) - 1):
        for x in range(1, len(crossword[y]) - 1):
            if crossword[y][x] != "A":
                continue

            top_left = crossword[y - 1][x - 1]
            top_right = crossword[y - 1][x + 1]
            bottom_left = crossword[y + 1][x - 1]
            bottom_right = crossword[y + 1][x + 1]

            if ((top_left == "M" and bottom_right == "S") or (top_left == "S" and bottom_right == "M")) and (
                (top_right == "M" and bottom_left == "S") or (top_right == "S" and bottom_left == "M")
            ):
                total += 1

    return total


def main() -> None:
    crossword = read_input(Path("input.txt"))
    print(f"Part 1 result: {_count_word_occurrences(crossword, "XMAS")}")
    print(f"Part 2 result: {_count_cross_mas(crossword)}")


if __name__ == "__main__":
    main()
