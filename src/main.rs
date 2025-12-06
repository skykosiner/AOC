use std::fs;


fn main() {
    let lines = fs::read_to_string("../2024/1/input").expect("Error reading file");
    let nums: Vec<(i32, i32)> = lines
        .split("\n")
        .filter_map(|s| {
            let mut parts = s.split("   ")
                .filter_map(|i| i.trim().parse::<i32>().ok());
            match (parts.next(), parts.next()) {
                (Some(l), Some(r)) => Some((l, r)),
                _ => None,
            }
        })
        .collect();

    let (mut left, mut right): (Vec<i32>, Vec<i32>) = nums.into_iter().unzip();
    left.sort();
    right.sort();

    let mut sum = 0;
    (0..left.len()).for_each(|i| {
        sum += (left[i]-right[i]).abs()
    });

    println!("{}", sum);
}
