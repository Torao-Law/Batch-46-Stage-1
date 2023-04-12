function getData(){
    let name = document.getElementById("name").value
    let email = document.getElementById("email").value
    let phone = document.getElementById("phone").value
    let subject = document.getElementById("subject").value
    let message = document.getElementById("message").value

    // validation
    if(name == "") {
        return alert("Nama Harus Di Isi")
    } else if(email == "") {
        return alert("Email Harus Di Isi")
    } else if(phone == "") {
        return alert("No Telpon Harus Di Isi")
    } else if(subject == "") {
        return alert("Subject Harus Dipilih")
    } else if(message == ""){
        return alert("Pesan Harus Di Isi")
    }

    const destination = "hi.dandi9@gmail.com"
    let a = document.createElement("a")
    a.setAttribute('href', `mailto:${destination}?subject=${subject}&body= Hallo nama saya ${name} , saya ingin ${message}, bisakah anda menghubungi saya di ${phone}`)

    a.click()

    let data = {
        nama: name,
        email: email,
        telp: phone,
        subject: subject,
        pesan: message,
    }

    console.log(data)
}

let blogs = []

function getBlog(event){
  event.preventDefault()

  let title = document.getElementById("input-blog-title").value
  let content = document.getElementById("input-blog-content").value
  let image = document.getElementById("input-blog-image").files

  image = URL.createObjectURL(image[0])

  let blog = {
      title,
      content,
      image,
      author: "Dandi Saputra",
      postedAt: new Date()
  }

  blogs.push(blog)    
  console.log(blogs)
  renderBlog()
}


function renderBlog() {
  document.getElementById("contents").innerHTML = ""
  for(let i = 0; i < blogs.length; i++) {
    document.getElementById("contents").innerHTML += `
    <div class="blog-list-item">
      <div class="blog-image">
        <img src="${blogs[i].image}" alt="" />
      </div>
      <div class="blog-content">
        <div class="btn-group">
          <button class="btn-edit">Edit Post</button>
          <button class="btn-post">Post Blog</button>
        </div>
        <h1>
          <a href="blog-detail.html" target="_blank"
            >${blogs[i].title}</a
          >
        </h1>
        <div class="detail-blog-content">
          ${createTime(blogs[i].postedAt)} | ${blogs[i].author}
        </div>
        <p>${blogs[i].content}</p>

        <div style="float:right; margin: 10px">
          <p style="font-size: 15px; color:grey">${getDuration(blogs[i].postedAt)}</p>
        </div>
      </div>
    </div>`
  }
}

function createTime(time) {
  let year = time.getFullYear() //tahun
  let indexMonth = time.getMonth() // bulan
  let date = time.getDate() // tanggal
  let hour = time.getHours() // jam
  let minute = time.getMinutes() //menit
  const month = ["Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agt", "Sep", "Okt", "Nov", "Des"]

  return `${date} ${month[indexMonth]} ${year} ${hour} : ${minute} WIB`
}

function getDuration(startDate, endDate) {
  const distance = new Date(endDate) - new Date(startDate)
  const dayDistance = Math.floor(distance / (24 * 60 * 60 * 1000))
  if(dayDistance > 0) {
    return dayDistance + " Day Ago" // 1 day ago
  } else {
    const hourDistance = Math.floor(distance / (60 * 60 * 1000))
    if(hourDistance > 0) {
      return hourDistance + " Hour Ago" // 3 hour ago
    } else {
      const minuteDistance =  Math.floor(distance / (60 * 1000))
      if(minuteDistance > 0) {
        return minuteDistance + " Minute Ago" // 55 minute ago
      } else {
        const secondDistance = Math.floor(distance / 1000)
        if(secondDistance > 0) {
          return secondDistance + " Second Ago" // 45 second ago 
        }
      }
    }
  }
}

setInterval(renderBlog, 1000)

// 1 minggu = 7 hari
// 1 hari = 24 jam
// 1 jam = 60 menit
// 1 menit = 60 detik
// 1 detik = 1000




