use std::{fs, io};
use std::io::{BufRead, BufReader};

pub fn get_lines_of_file(path: &str) -> io::Result<io::Lines<BufReader<fs::File>>> {
    let input = fs::File::open(path)?;
    let reader = BufReader::new(input);
    Ok(reader.lines())
}