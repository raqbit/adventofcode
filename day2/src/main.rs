use std::convert::identity;
use std::io;

#[derive(Debug)]
struct Set {
    red: u8,
    green: u8,
    blue: u8,
}

impl Set {
    fn from_str(set: &str) -> Option<Self> {
        let cube_colors = set.split(", ");

        let mut red = 0;
        let mut green = 0;
        let mut blue = 0;

        for cubes in cube_colors {
            let (amount, color) = cubes.split_once(" ")?;
            let amount: u8 = str::parse(amount).ok()?;
            match color {
                "red" => {
                    red += amount;
                }
                "green" => {
                    green += amount;
                }
                "blue" => {
                    blue += amount;
                }
                _ => {
                    panic!("Unknown color")
                }
            }
        }
        Some(Self {
            red,
            blue,
            green,
        })
    }
}

#[derive(Debug)]
struct Game {
    id: u8,
    sets: Vec<Set>,
}

impl Game {
    fn from_string(game: &String) -> Option<Self> {
        let (game_id, sets) = game.split_once(": ")?;
        let game_id = str::parse(game_id.split_once(" ")?.1).ok()?;
        let set_list: Vec<&str> = sets.split("; ").collect();

        let mut parsed_sets = Vec::with_capacity(set_list.len());

        for entry in set_list {
            parsed_sets.push(Set::from_str(entry)?);
        }

        Some(Self {
            id: game_id,
            sets: parsed_sets,
        })
    }
}

fn main() -> io::Result<()> {
    let lines: Vec<String> = shared::get_lines_of_file("input/day2.txt")?
        .filter_map(Result::ok)
        .collect();

    let games: Vec<Game> = lines.iter().map(Game::from_string).filter_map(identity).collect();

    let result = find_possible_games(&games);

    println!("Result part 1: {:}", result);

    Ok(())
}

fn find_possible_games(games: &Vec<Game>) -> u32 {
    const MAX_RED: u8 = 12;
    const MAX_GREEN: u8 = 13;
    const MAX_BLUE: u8 = 14;

    games.iter()
        .fold(0, |mut acc, game| {
            if game.sets.iter().all(|set| set.red <= MAX_RED && set.green <= MAX_GREEN && set.blue <= MAX_BLUE) {
                acc += u32::from(game.id);
            }
            acc
        })
}
