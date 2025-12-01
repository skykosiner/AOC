import fs from "fs";

function splitLeftAndRight(lines: string[]): [number[], number[]] {
    const left: number[] = [];
    const right: number[] = [];

    for (let i = 0; i < lines.length; i++) {
        const nums = lines[i].split("   ");
        if (nums.length === 1) {
            break;
        }

        left.push(+nums[0]);
        right.push(+nums[1]);
    };

    return [left, right]
}

const lines = fs.readFileSync("../day1/input").toString().split("\n");
const [left, right] = splitLeftAndRight(lines);
left.sort()
right.sort()

let distance = 0;
for (let i = 0; i < left.length; i++) {
    distance += Math.abs(left[i] - right[i])
}

console.log(distance);
