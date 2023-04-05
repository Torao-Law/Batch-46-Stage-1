class Car {
    constructor(make, model) { // make = "Toyota", model = "Camry"
        this.make = make
        this.model = model
    }

    getInfo() {
        return `The car is a ${this.make} ${this.model}`
    }
}

// OBJECT
let myCar = new Car("Toyota", "Camry")
let yourCar = new Car("Toyota", "Alpard")
console.log(myCar.getInfo()) // The car is a toyota camry
console.log(yourCar.getInfo()) // The car is a toyota Alpard