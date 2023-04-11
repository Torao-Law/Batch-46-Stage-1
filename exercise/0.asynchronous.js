// Synchronous
console.log("Pesan 1")
console.log("Pesan 2")
console.log("Pesan 3")

console.log("============================")

// Asynchronous 
console.log("Start 1")
setTimeout(() => console.log("Pesan 2"), 2000)
console.log("Start 3")
