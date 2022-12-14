const fs = require('fs')
const performance = require('perf_hooks').performance
const eol = require('os').EOL

let startTime = performance.now()
let maxX = 0
let maxY = -Infinity

function parseInput() {
  return fs
    .readFileSync(__dirname + '/testdata/input.txt', 'utf8')
    .split(eol)
    .map((r) =>
      r.split(' -> ').map((i) => {
        coord = i.split(',').map(Number)
        maxX = Math.max(coord[0] + 1, maxX)
        if (coord[1]) {
          maxY = Math.max(coord[1] + 1, maxY)
        }
        return coord
      }),
    )
}

class Cave {
  cave
  isFull = false
  constructor(width, height) {
    this.cave = Array(height)
      .fill()
      .map(() => Array(width).fill(' '))
  }
  draw(x, y, value) {
    this.cave[y][x] = value
  }
  getAtPos(x, y) {
    if (this.cave[y] === undefined) {
      this.isFull = true
      return ''
    }
    let state = this.cave[y][x]
    if (state === undefined) {
      this.isFull = true
    }
    return state
  }
  addSand(x, y) {
    this.draw(x, y, 'X')
    while (true) {
      if (this.getAtPos(x, y + 1) == ' ') {
        y += 1
      } else if (this.getAtPos(x - 1, y + 1) == ' ') {
        x -= 1
        y += 1
      } else if (this.getAtPos(x + 1, y + 1) == ' ') {
        x += 1
        y += 1
      } else {
        break
      }
    }
    if (!this.isFull) {
      this.draw(x, y, '0')
    }
    if (x == 500 && y == 0) {
      this.isFull = true
    }
    return this.isFull
  }
  drawWall(startX, startY, endX, endY) {
    let [sX, eX] = [startX, endX].sort((a, b) => a - b)
    let [sY, eY] = [startY, endY].sort((a, b) => a - b)
    if (sX == eX)
      for (let y = sY; y <= eY; y++) {
        this.draw(sX, y, '#')
      }
    if (sY == eY)
      for (let x = sX; x <= eX; x++) {
        this.draw(x, sY, '#')
      }
  }
}

function partOne() {
  let time = performance.now() - startTime
  let count = 0
  const input = parseInput()
  const cave = new Cave(maxX, maxY)
  input.forEach((row) => {
    for (let i = 0; i < row.length - 1; i++) {
      cave.drawWall(...row[i], ...row[i + 1])
    }
  })
  while (!cave.addSand(500, 0)) count++
  console.log(`Part 1: ${count}\nTimer: ${time} ms`)
}

function partTwo() {
  let time = performance.now() - startTime
  let count = 0
  const input = parseInput()
  let cave = new Cave(maxX * 2, maxY + 2)
  input.forEach((row) => {
    for (let i = 0; i < row.length - 1; i++) {
      cave.drawWall(...row[i], ...row[i + 1])
    }
  })
  cave.drawWall(0, maxY + 1, maxX * 2, maxY + 1)
  while (!cave.addSand(500, 0)) count++
  console.log()
  console.log(`Part 2: ${count}\nTimer: ${time} ms`)
}

partOne()
partTwo()
