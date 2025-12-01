use std::fs;

fn main() {
    println!("{:?}", fs::read_to_string("./2024/day1/input.test").map);
}
