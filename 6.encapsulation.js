class Car {
    contructor(make, model) {
        this._make = make
        this._model = model
    }

    // getter
    get  make() {
        return this._make;
    }

    get model() {
        return this._model
    }

    // setter
    set make(value) {
        if (value === "") {
            console.log("the make cannot be empty")
            return
        }
        this._make = value
    }
    
    set model(value) {
        if (value === "") {
            console.log("the model cannot be empty")
            return
        }
        this._model = value
    }
}

let myCar = new Car("Toyota", "Camry")
myCar.make = "Honda"
myCar.model = ""
console.log(myCar.make)
console.log(myCar.model)
