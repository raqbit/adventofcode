from pathlib import Path

from shared import read_input


def _count_word_occurrences(crossword: str, word: str) -> int:
    """
    Note: This has a known bug in the edge case that any word searched ends
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


def main() -> None:
    crossword = read_input(Path("input.txt"))
    # 2607: too low
    print(f"Part 1 result: {_count_word_occurrences(crossword, "XMAS")}")


if __name__ == "__main__":
    main()
