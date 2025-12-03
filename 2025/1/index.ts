import fs from "fs";

const lines = fs.readFileSync("./input").toString().split("\n");
// let partTwo = 0;

function partOne() {
    let dial = 50;
    let sum = 0;

    for (let line of lines) {
        const regex = new RegExp("[0-9].*")
        const res = regex.exec(line)

        if (res === null) break;

        const num = +res.toString()
        if (line[0] === "L") dial += num
        if (line[0] === "R") dial -= num

        dial = (dial + 100) % 100;
        if (dial === 0) sum++
    }

    console.log(sum);
}

partOne();
