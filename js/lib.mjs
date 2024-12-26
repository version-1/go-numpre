export function stringToGrid(str) {
  const grid = [];
  for (let i = 0; i < 9; i++) {
    const row = [];
    for (let j = 0; j < 9; j++) {
      row.push(parseInt(str[i * 9 + j]));
    }
    grid.push(row);
  }
  return grid;
}

export function gridToString(grid) {
  let str = "";
  for (let i = 0; i < 9; i++) {
    for (let j = 0; j < 9; j++) {
      str += grid[i][j];
    }
  }
  return str;
}

export function validate(grid) {
  for (let i = 0; i < 9; i++) {
    for (let j = 0; j < 9; j++) {
      const n = grid[i][j];
      if (n === 0 || !possible(grid, i, j, n)) {
        return { x: i, y: j, n };
      }
    }
  }
  return;
}

export function possible(grid, x, y, n) {
  for (let i = 0; i < 9; i++) {
    if ((y !== i && grid[x][i] === n) || (x !== i && grid[i][y] === n)) {
      return false;
    }
  }

  const x0 = Math.floor(x / 3) * 3;
  const y0 = Math.floor(y / 3) * 3;
  for (let i = 0; i < 3; i++) {
    for (let j = 0; j < 3; j++) {
      if (x0 + i !== x && y0 + j !== y && grid[x0 + i][y0 + j] === n) {
        return false;
      }
    }
  }

  return true;
}

export function render(grid) {
  for (let i = 0; i < 9; i++) {
    const row = grid[i];
    let str = "| ";
    for (let j = 0; j < 9; j++) {
      str += row[j] + " ";
      if (j % 3 === 2) {
        str += "| ";
      }
    }
    console.log(str);
    if (i % 3 === 2) {
      console.log("-".repeat(24));
    }
  }
}

export function solve(grid) {
  for (let i = 0; i < 9; i++) {
    for (let j = 0; j < 9; j++) {
      if (grid[i][j] === 0) {
        for (let n = 1; n <= 9; n++) {
          if (possible(grid, i, j, n)) {
            grid[i][j] = n;
            if (solve(grid)) {
              return true;
            }
            grid[i][j] = 0;
          }
        }
        return false;
      }
    }
  }
  return true;
}
