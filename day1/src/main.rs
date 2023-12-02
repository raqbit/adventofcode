use std::io;

use daachorse::{DoubleArrayAhoCorasick, DoubleArrayAhoCorasickBuilder};
use daachorse::errors::DaachorseError;

fn main() -> io::Result<()> {
    let lines: Vec<String> = shared::get_lines_of_file("input/day1.txt")?
        .filter_map(Result::ok)
        .collect();

    let ac_digits = build_ac(SearchKind::DigitsOnly).unwrap();
    let ac = build_ac(SearchKind::Normal).unwrap();
    let reverse_ac = build_ac(SearchKind::Reverse).unwrap();

    let result: u32 = lines
        .iter()
        .fold(0, |mut acc, line| {
            acc += u32::from(find_calibration_value(&ac_digits, &ac_digits, line));
            acc
        });

    println!("Result part 1: {}", result);
    let result = lines
        .iter()
        .fold(0, |mut acc, line| {
            acc += u32::from(find_calibration_value(&ac, &reverse_ac, line));
            acc
        });

    println!("Result part 2: {}", result);

    Ok(())
}

fn find_calibration_value(ac: &DoubleArrayAhoCorasick<u8>, reverse_ac: &DoubleArrayAhoCorasick<u8>, line: &String) -> u8 {

    fn find_digit(ac: &DoubleArrayAhoCorasick<u8>, line: &str) -> Option<u8> {
        let mut iter = ac.find_iter(line);

        Some(iter.next()?.value())
    }

    let first_digit = find_digit(ac, line).unwrap_or(0);

    let last_digit = find_digit(reverse_ac, line.chars().rev().collect::<String>().as_str()).unwrap_or(0);

    first_digit * 10 + last_digit
}

enum SearchKind {
    DigitsOnly,
    Normal,
    Reverse,
}

/// Builds a data structure for the Aho-Corasick Algorithm
///
/// I went into a deep rabbit hole, reading into deterministic finite state
/// machines and the Aho-Corasick algorithm. I did an attempt at an implementation
/// in Rust, but that ended up being a fight with the borrow checker due to the
/// recursive Trie data structure. As I should really go get some sleep, I ended up
/// installing a crate to prove that the idea is correct.
fn build_ac(kind: SearchKind) -> Result<DoubleArrayAhoCorasick<u8>, DaachorseError> {
    let mut needles = vec![
        ("1".to_string(), 1),
        ("2".to_string(), 2),
        ("3".to_string(), 3),
        ("4".to_string(), 4),
        ("5".to_string(), 5),
        ("6".to_string(), 6),
        ("7".to_string(), 7),
        ("8".to_string(), 8),
        ("9".to_string(), 9),
    ];

    let mut words = vec![
        ("one".to_string(), 1),
        ("two".to_string(), 2),
        ("three".to_string(), 3),
        ("four".to_string(), 4),
        ("five".to_string(), 5),
        ("six".to_string(), 6),
        ("seven".to_string(), 7),
        ("eight".to_string(), 8),
        ("nine".to_string(), 9),
    ];

    match kind {
        SearchKind::Normal => {
            needles.append(&mut words)
        }
        SearchKind::Reverse => {
            // To search from the back to the front, we reverse the haystack
            // Therefore we'll need to reverse the needles too
            needles.extend(words.iter()
                .map(|(w, x)| (
                    w.chars().rev().collect::<String>(),
                    *x
                ))
            );
        }
        SearchKind::DigitsOnly => {}
    }

    DoubleArrayAhoCorasickBuilder::new()
        .build_with_values(needles)
}
