function hello() {
    return "Hello world"
}

function print(name, callback) {
    return name + callback()
}

const x = print("Dandi", hello)
console.log(x)

