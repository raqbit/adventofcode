from pathlib import Path

from shared import read_input


def _count_word_occurrences(crossword: str, word: str) -> int:
    """
    Note: This has a known bug in the edge case that any word searched ends
    right at the start of the input. The max()'s are there to prevent this from being
    an issue
    """
    row_len = crossword.find("\n") + 1

    total = 0
    for pos in range(len(crossword)):
        if crossword[pos] != word[0]:
            continue

        # Horizontal
        if crossword[pos : pos + len(word)] == word:
            total += 1

        # Horizontal reverse
        if crossword[pos : max(pos - len(word), 0) : -1] == word:
            total += 1

        # Vertical down
        if crossword[pos : pos + row_len * len(word) : row_len] == word:
            total += 1

        # Vertical up
        if crossword[pos : max(pos - (row_len * len(word)), 0) : -row_len] == word:
            total += 1

        # Diagonal down right
        if crossword[pos : pos + row_len * len(word) : row_len + 1] == word:
            total += 1

        # Diagonal down left
        if crossword[pos : pos + row_len * (len(word) - 1) : row_len - 1] == word:
            total += 1

        # Diagonal up left
        if crossword[pos : max(pos - (row_len * len(word)), 0) : -(row_len + 1)] == word:
            total += 1

        # Diagonal up right
        if crossword[pos : max(pos - (row_len * (len(word) - 1)), 0) : -(row_len - 1)] == word:
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
    # 2607: too low
    print(f"Part 1 result: {_count_word_occurrences(crossword, "XMAS")}")
    print(f"Part 2 result: {_count_cross_mas(crossword)}")


if __name__ == "__main__":
    main()
