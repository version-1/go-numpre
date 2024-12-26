import { gridToString, stringToGrid, validate, render, solve } from "./lib.mjs";

function main() {
  const args = process.argv.slice(2);
  if (args.length !== 1) {
    console.error("Usage: node index.js <string>");
    process.exit(1);
  }
  const grid = stringToGrid(args[0]);

  console.log("[START]");
  render(grid);
  solve(grid);
  const result = validate(grid);
  if (result) {
    console.log("[END] Invalid solution");
    console.log(result);
  } else {
    console.log("[END]");
  }
  render(grid);
  console.log(gridToString(grid));
}

main();
