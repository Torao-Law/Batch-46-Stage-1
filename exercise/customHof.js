// custom hof return other function
function multplyBy(n) {
    return function (x) {
        return x * n
    }
}

const double = multplyBy(2)
console.log(double(5))


// custom hof with callback
function logNumber(n) {
    console.log(`The number is ${n}`)
}

function repeat(n, action) {
    for(let i = 0; i < n; i++) {
        action(i)
    }
}

repeat(3, logNumber)