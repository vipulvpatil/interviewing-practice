import * as fs from "fs"
import readline from "readline"

function main() {
  processLineByLine();
}

main()

async function processLineByLine() {
  const fileStream = fs.createReadStream('input.txt');

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity,
  });
  for await (const line of rl) {
    console.log(`Line from file: ${line}`);
  }
}
