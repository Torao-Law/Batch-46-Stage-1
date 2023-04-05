class Vehicle {
    drive() {
        return "The vehicle is driving"
    }
}

class Car extends Vehicle {
    drive() {
        return "The car is driving"
    }
}

class ElectricCar extends Vehicle {
    drive() {
        return "The Electric Car is Driving silently"
    }
}

let myVehicle = new Vehicle()
let myCar = new Car()
let myElectricCar = new ElectricCar()

console.log(myVehicle.drive())
console.log(myCar.drive())
console.log(myElectricCar.drive())