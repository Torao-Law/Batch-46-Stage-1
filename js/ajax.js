const testimonial = new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest()

    xhr.open("GET", "https://api.npoint.io/19dc1d52f9465c39624a", true)

    xhr.onload = function (){
        if(xhr.status == 200) {
            resolve(JSON.parse(xhr.response))
        } else {
            reject("Error loading data")
        }
    }

    xhr.onerror = function() {
        reject("Network Error")
    }

    xhr.send()
})

async function getAllTestimonial() {
    const response = await testimonial
    console.log(response)

    let testimonialForHtml = ""
    response.forEach((item) => {
        testimonialForHtml += `<div class="testimonial">
            <img src="${item.image}" class="profile-testimonial" />
            <p class="quote">"${item.quote}"</p>
            <p class="author">- ${item.author}</p>
            <p class="author">${item.rating} <i class="fa-solid fa-star"></i></p>
        </div>`
    })

    document.getElementById("testimonials").innerHTML = testimonialForHtml
}
getAllTestimonial()


async function filterTestimonials(rating) {
    const response =  await testimonial
    let testimonialHTML = '';

    const testimonialFiltered = response.filter(function (item) {
        return item.rating === rating;
    })

    if (testimonialFiltered.length === 0) {
        testimonialHTML = `<h1> Data not found! </h1>`
    } else {
        testimonialFiltered.forEach(function (item) {
            testimonialHTML += `<div class="testimonial">
                <img src="${item.image}" class="profile-testimonial" />
                <p class="quote">"${item.quote}"</p>
                <p class="author">- ${item.author}</p>
                <p class="author">${item.rating} <i class="fa-solid fa-star"></i></p>
            </div>`
        })
    }

    document.getElementById('testimonials').innerHTML = testimonialHTML;
}








// [
//     {
//     "image": "https://images.unsplash.com/photo-1570295999919-56ceb5ecca61?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Mnx8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60",
//     "quote": "Keren banget jasanya!",
//     "author": "Farhan Hadyan",
//     "rating": 5
//     },
//     {
//     "image": "https://images.unsplash.com/photo-1568602471122-7832951cc4c5?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8M3x8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60",
//     "quote": "Keren lah pokoknya!",
//     "author": "Adiguna Sanjaya",
//     "rating": 4
//     }
// ]

// [
//     {
//     image: "https://images.unsplash.com/photo-1570295999919-56ceb5ecca61?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Mnx8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60",
//     quote: "Keren banget jasanya!",
//     author: "Farhan Hadyan",
//     rating: 5
//     },
//     {
//     image: "https://images.unsplash.com/photo-1568602471122-7832951cc4c5?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8M3x8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60",
//     quote: "Keren lah pokoknya!",
//     author: "Adiguna Sanjaya",
//     rating: 4
//     }
// ]