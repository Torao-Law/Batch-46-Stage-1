// Promise (Janji)
// Object yang mepresentasikan keberhasilan dan kegagalan sebuah peristiwa pada asynchronous dimasa datang
// Janji = (terpenuhi, ingkar)
// states = (Fulfilled, Rejected, Pending)
// callback = (resolve, reject, finally)
// action = (then, catch)

// CONTOH IMPLEMENT 1
const status = false
let promise = new Promise((resolve, reject) => {
    if(status) {
        resolve("Promise is resolved")
    } else {
        reject("Promise is rejected")
    }
})

console.log(promise)

promise
    .then((value) => console.log(value))
    .catch((value) => console.log(value))
