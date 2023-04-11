// const testimonialData = [
//     {
//         author: "Surya Elidanto",
//         quote: "Keren banget jasanya!",
//         image: "https://images.unsplash.com/photo-1570295999919-56ceb5ecca61?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Mnx8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60",
//         rating: 5
//     },
//     {
//         author: "Surya Elz",
//         quote: "Keren lah pokoknya!",
//         image: "https://images.unsplash.com/photo-1568602471122-7832951cc4c5?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8M3x8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60",
//         rating: 4
//     },
//     {
//         author: "Surya Gans",
//         quote: "The best pelayanannya!",
//         image: "https://images.unsplash.com/photo-1564564321837-a57b7070ac4f?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8OHx8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60",
//         rating: 4
//     }
// ]

// function allTestimonials() {
//     let testimonialHTML = '';

//     testimonialData.forEach(function (item) {
//         testimonialHTML += `<div class="testimonial">
//             <img src="${item.image}" class="profile-testimonial" />
//             <p class="quote">"${item.quote}"</p>
//             <p class="author">- ${item.author}</p>
//             <p class="author">${item.rating} <i class="fa-solid fa-star"></i></p>
//         </div>`
//     })

//     document.getElementById('testimonials').innerHTML = testimonialHTML;
// }

// allTestimonials()


// filter testimonials
// function filterTestimonials(rating) {
//     let testimonialHTML = '';

//     const testimonialFiltered = testimonialData.filter(function (item) {
//         return item.rating === rating;
//     })

//     // console.log(testimonialFiltered);

//     if (testimonialFiltered.length === 0) {
//         testimonialHTML = `<h1> Data not found! </h1>`
//     } else {
//         testimonialFiltered.forEach(function (item) {
//             testimonialHTML += `<div class="testimonial">
//                 <img src="${item.image}" class="profile-testimonial" />
//                 <p class="quote">"${item.quote}"</p>
//                 <p class="author">- ${item.author}</p>
//                 <p class="author">${item.rating} <i class="fa-solid fa-star"></i></p>
//             </div>`
//         })
//     }

//     document.getElementById('testimonials').innerHTML = testimonialHTML;
// }