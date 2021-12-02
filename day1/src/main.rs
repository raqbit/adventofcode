use shared;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let nums: Vec<u32> =
        shared::get_lines_of_file("input.txt")?
            .filter_map(|res| res.ok())
            .filter_map(|line| line.parse::<u32>().ok())
            .collect();

    let increases = part1(&nums);
    println!("Part 1 number of increases: {}", increases);

    let increases = part2(&nums);
    println!("Part 2 number of increases: {}", increases);

    Ok(())
}

fn part1(nums: &[u32]) -> u32 {
    let mut prev_num: u32 = 0;
    let mut increases: u32 = 0;

    for (i, &num) in nums.iter().enumerate() {
        if i != 0 && num > prev_num {
            increases += 1;
        }
        prev_num = num;
    }

    increases
}

fn part2(nums: &[u32]) -> u32 {
    let mut prev_num: u32 = 0;
    let mut increases: u32 = 0;

    for i in 0..nums.len() - 2 {
        let window = &nums[i..i + 3];
        let sum = window.iter().sum::<u32>();

        if i != 0 && sum > prev_num {
            increases += 1;
        }

        prev_num = sum;
    }

    increases
}

