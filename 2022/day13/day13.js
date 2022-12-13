const input = require('./testdata/input.json')

function isValid(a, b) {
  if (typeof a === 'number' && typeof b === 'number') {
    if (a < b) {
      return 1
    }
    if (a === b) {
      return 0
    }
    return -1
  }

  if (Array.isArray(a) && typeof b === 'number') {
    b = [b]
  }

  if (Array.isArray(b) && typeof a === 'number') {
    a = [a]
  }

  if (Array.isArray(a) && Array.isArray(b)) {
    let i = 0
    while (i < a.length && i < b.length) {
      const x = isValid(a[i], b[i])
      if (x === 1) {
        return 1
      }
      if (x === -1) {
        return -1
      }
      i++
    }
    if (i == a.length) {
      if (a.length == b.length) {
        return 0
      }
      return 1
    }

    return -1
  }
}

function PartOne() {
  let sum = 0
  let pairNumber = 1
  for (let i = 0; i < input.length; i += 2) {
    const left = input[i]
    const right = input[i + 1]
    const validPair = isValid(left, right)
    if (validPair == 1) {
      sum += pairNumber
    }
    pairNumber++
  }
  console.log(sum)
}

function PartTwo() {
  input.push([[2]], [[6]])
  input.sort(isValid).reverse()
  let sum = 1
  for (let i = 0; i < input.length; i++) {
    if (input[i].length === 1 && input[i][0] == 2) {
      sum = sum * (i + 1)
    }

    if (input[i].length === 1 && input[i][0] == 6) {
      sum = sum * (i + 1)
    }
  }
  console.log(sum)
}

PartOne()
PartTwo()
