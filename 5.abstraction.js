class Car {
    contructor(make, model) {
        this._make = make
        this._model = model
    }

    // getter
    get make() {
        return this._make;
    }

    get model() {
        return this._model
    }
}

let myCar = new Car("Toyota", "Camry")
myCar.make = "honda"
myCar.model = "honda"
console.log(myCar.make)
