use std::io;

fn main() -> io::Result<()> {
    let lines: Vec<String> = shared::get_lines_of_file("input/day1.txt")?
        .filter_map(Result::ok)
        .collect();

    let result: u32 = lines
        .iter()
        .fold(0, |mut acc, line| {
            acc += u32::from(find_calibration_value(line));
            acc
        });

    println!("Result part 1: {}", result);

    Ok(())
}

fn find_calibration_value(line: &String) -> u8 {
    fn is_digit(chr: &char) -> bool {
        return *chr >= '0' && *chr <= '9';
    }

    let first_digit = line.chars()
        .find(is_digit)
        .map(|c| (c as u8) - ('0' as u8))
        .unwrap_or(0);

    let last_digit = line.chars().rev()
        .find(is_digit)
        .map(|c| (c as u8) - ('0' as u8))
        .unwrap_or(0);

    first_digit * 10 + last_digit
}

fn is_digit(chr: &char) -> bool {
    return *chr >= '0' && *chr <= '9';
}