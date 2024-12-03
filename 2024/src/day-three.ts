import fs from "fs";


function partOne() {
    let sum = 0;

    const lines = fs.readFileSync("./pkg/day-three/input").toString();
    const r: RegExp = /mul\([0-9]+,[0-9]+\)/g;
    const matches = lines.match(r);

    matches?.map(item => {
        const nums = item.match(/\d+/g);
        if (nums) {
            const n = nums.map(Number);
            sum += n[0] * n[1]
        }
    })

    console.log(sum);
}

function partTwo() {
    let sum = 0;
    let enabled = true;

    const lines = fs.readFileSync("./pkg/day-three/input").toString();
    const r: RegExp = /mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)/g;
    const matches = lines.match(r);

    matches?.map(item => {
        switch (item) {
            case "do()":
                enabled = true
                break;
            case "don't()":
                enabled = false
                break;
            default:
                if (enabled) {
                    const nums = item.match(/\d+/g);
                    if (nums) {
                        const n = nums.map(Number);
                        sum += n[0] * n[1]
                    }
                }

                break;
        }
    })

    console.log(sum);
}

partOne();
partTwo();
