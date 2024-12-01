import pathlib
from collections.abc import Iterable


def read_input(file: pathlib.PurePath) -> Iterable[str]:
    with open(file) as f:
        for line in f:
            yield line.strip()