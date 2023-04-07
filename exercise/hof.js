// forEach
const myNumber = [1, 2, 3, 4, 5]

myNumber.forEach(function (item) {
    console.log(item)
})

// map
const myNumbers = [1, 2, 3, 4, 5]
const double = myNumbers.map((item) => {
    return item * 2
})
console.log(double)


// filter
const myAge = [20, 24, 16, 29, 30]
const canJoin = myAge.filter((data) => {
    return data % 2 === 1;
})
console.log(canJoin)

// // 20 / 2 = 10 => 0
// // 29 / 2 = 9 => 1


// Reduce
const myNumberr = [1, 2, 3, 4, 5]
const reduce = myNumberr.reduce((prev, current) => {
    return prev + current
}, 0)

console.log(reduce)