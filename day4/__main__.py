from pathlib import Path

from shared import read_input


def _count_word_occurrences(crossword: str, word: str) -> int:
    row_len = crossword.find("\n")+1

    total = 0
    for pos in range(len(crossword)):
        x_pos = pos % row_len
        y_pos = pos // row_len
        if crossword[pos] != word[0]:
            continue

        # Horizontal
        if crossword[pos:pos+4] == word:
            total += 1

        # Horizontal reverse
        if crossword[pos:pos-4:-1] == word:
            total += 1

        # Vertical down
        if crossword[pos:pos+row_len*4:row_len] == word:
            total += 1

        # Vertical up
        if crossword[pos:pos-(row_len*4):-row_len] == word:
            total += 1

        # Diagonal down right
        if crossword[pos:pos+row_len*4:row_len+1] == word:
            total += 1

        # Diagonal down left
        if crossword[pos:pos+row_len*3:row_len-1] == word:
            total += 1

        # Diagonal up left
        if crossword[pos:pos-(row_len*4):-(row_len+1)] == word:
            total += 1

        # Diagonal up right
        if crossword[pos:pos-(row_len*3):-(row_len-1)] == word:
            total += 1

    return total


def main() -> None:
    crossword = read_input(Path("input.txt"))
    print(f"Part 1 result: {_count_word_occurrences(crossword, "XMAS")}")


if __name__ == "__main__":
    main()
