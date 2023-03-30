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
              ${blogs[i].postedAt} | ${blogs[i].author}
            </div>
            <p>${blogs[i].content}</p>

            <div style="float:right; margin: 10px">
              <p style="font-size: 15px; color:grey">1 minutes ago</p>
            </div>
          </div>
        </div>`
    }
}






