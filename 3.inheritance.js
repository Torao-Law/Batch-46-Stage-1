class Car {
    constructor(make, model) {
        this.make = make
        this.model = model
    }

    getInfo() {
       return `The car is a ${this.make} ${this.model}`
    }
}

class ElectricCar extends Car {
    constructor(make, model, batteryCapacity) {
        super(make, model)
        this.batteryCapacity = batteryCapacity
    }

    getInfo() {
        return `${super.getInfo()} It has a battery capicity of ${this.batteryCapacity} kwh.`
    }
}

let myElectricCar = new ElectricCar("Tesla", "Model A", 120)
console.log(myElectricCar.getInfo())