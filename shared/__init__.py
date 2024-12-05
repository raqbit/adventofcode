from collections.abc import Iterable
from pathlib import Path


def read_input_lines(file: Path) -> Iterable[str]:
    with file.open() as f:
        for line in f:
            yield line.strip()

def read_input(file: Path) -> str:
    with file.open() as f:
        return f.read().strip()
