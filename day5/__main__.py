import dataclasses
import graphlib
from itertools import takewhile
from pathlib import Path

from shared import read_input_lines


@dataclasses.dataclass(frozen=True)
class OrderRule:
    before: int
    after: int


def _build_topo_for_page(rules: list[OrderRule], page_nums: set[int]) -> list[int]:
    # https://en.wikipedia.org/wiki/Topological_sorting
    rule_topo = graphlib.TopologicalSorter()

    for rule in rules:
        if rule.before in page_nums and rule.after in page_nums:
            rule_topo.add(rule.after, rule.before)

    return list(rule_topo.static_order())


def main() -> None:
    input_lines = read_input_lines(Path("input.txt"))

    rules = [
        OrderRule(int(left), int(right))
        for left, right in (line.split("|") for line in takewhile(lambda x: x != "", input_lines))
    ]

    pages = [[int(x) for x in line.split(",")] for line in input_lines]

    part1_total = 0
    part2_total = 0
    for page in pages:
        rule_topo = _build_topo_for_page(rules, set(page))

        if rule_topo == page:
            part1_total += page[len(page) // 2]
        else:
            part2_total += rule_topo[len(rule_topo) // 2]

    print(f"Part 1 result: {part1_total}")
    print(f"Part 1 result: {part2_total}")


if __name__ == "__main__":
    main()
