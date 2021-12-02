use std::str::FromStr;

use shared;

#[derive(Debug)]
enum Direction {
    FORWARD,
    DOWN,
    UP,
}

struct Command {
    direction: Direction,
    amount: u32,
}

impl FromStr for Direction {
    type Err = ();

    fn from_str(input: &str) -> Result<Self, Self::Err> {
        match input {
            "forward" => Ok(Direction::FORWARD),
            "down" => Ok(Direction::DOWN),
            "up" => Ok(Direction::UP),
            _ => Err(())
        }
    }
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let commands: Vec<Command> =
        shared::get_lines_of_file("input.txt")?
            .filter_map(|res| res.ok())
            .filter_map(|line| {
                if let [direction, amount] = line.split_whitespace().collect::<Vec<&str>>()[..] {
                    let direction = Direction::from_str(direction).ok()?;
                    let amount = amount.parse::<u32>().ok()?;
                    return Some(Command { direction, amount });
                }
                None
            })
            .collect();

    let result = part1(&commands);

    println!("Part 1 result: {:}", result);

    let result = part2(&commands);

    println!("Part 2 result: {:}", result);

    Ok(())
}

fn part1(commands: &[Command]) -> u32 {
    let mut horizontal = 0;
    let mut depth = 0;

    for cmd in commands {
        match cmd.direction {
            Direction::FORWARD => { horizontal += cmd.amount }
            Direction::DOWN => { depth += cmd.amount }
            Direction::UP => { depth -= cmd.amount }
        }
    }

    horizontal * depth
}

fn part2(commands: &[Command]) -> u32 {
    let mut horizontal = 0;
    let mut depth = 0;
    let mut aim = 0;

    for cmd in commands {
        match cmd.direction {
            Direction::FORWARD => {
                horizontal += cmd.amount;
                depth += aim * cmd.amount;
            }
            Direction::DOWN => { aim += cmd.amount }
            Direction::UP => { aim -= cmd.amount }
        }
    }

    horizontal * depth
}