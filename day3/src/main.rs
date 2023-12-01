fn main() -> Result<(), Box<dyn std::error::Error>> {
    let lines: Vec<String> =
        shared::get_lines_of_file("input/day3.txt")?
            .filter_map(|res| res.ok())
            .collect();
    let nums: Vec<u16> = lines
        .iter()
        .filter_map(|line| u16::from_str_radix(&line, 2).ok())
        .collect();

    let mut gamma = 0;
    let mut epsilon = 0;

    for i in 0..lines[0].len() {
        let mut zero_count = 0;
        let mut one_count = 0;

        for &num in &nums {
            let mut bit = num;
            bit >>= i;
            bit &= 0x1;

            if bit & 0x1 == 0 {
                zero_count += 1;
            } else {
                one_count += 1;
            }
        }

        if zero_count > one_count {
            gamma = (gamma & !0x1) | zero_count;
            epsilon = gamma & !0x1 | one_count;
        }
    }

    println!("{} * {} = {}", gamma, epsilon, gamma*epsilon);


    Ok(())
}
