package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Blog struct {
	Title    string
	Content  string
	Author   string
	PostDate string
}

var dataBlog = []Blog{
	{
		Title:    "Pasar coding dinilai masih menjanjikan",
		Content:  "Ketimpangan sumber daya manusia (SDM) di sektor digital masih menjadi isu yang belum terpecahkan. Berdasarkan penelitian ManpowerGroup, ketimpangan SDM global, termasuk Indonesia, meningkat dua kali lipat dalam satu dekade terakhir. Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quam, molestiae numquam! Deleniti maiores expedita eaque deserunt quaerat! Dicta, eligendi debitis?",
		Author:   "Jaya Saleh",
		PostDate: "14/04/2023",
	},
	{
		Title:    "Pasar coding dinilai masih sedikit",
		Content:  "Ketimpangan sumber daya manusia (SDM) di sektor digital masih menjadi isu yang belum terpecahkan. Berdasarkan penelitian ManpowerGroup, ketimpangan SDM global, termasuk Indonesia, meningkat dua kali lipat dalam satu dekade terakhir. Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quam, molestiae numquam! Deleniti maiores expedita eaque deserunt quaerat! Dicta, eligendi debitis?",
		Author:   "Yoga Wicaksono",
		PostDate: "15/04/2023",
	},
	{
		Title:    "Pasar coding dinilai masih sedikit",
		Content:  "Ketimpangan sumber daya manusia (SDM) di sektor digital masih menjadi isu yang belum terpecahkan. Berdasarkan penelitian ManpowerGroup, ketimpangan SDM global, termasuk Indonesia, meningkat dua kali lipat dalam satu dekade terakhir. Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quam, molestiae numquam! Deleniti maiores expedita eaque deserunt quaerat! Dicta, eligendi debitis?",
		Author:   "Yoga Wicaksono",
		PostDate: "15/04/2023",
	},
}

func main() {
	// create new echo instance
	e := echo.New()

	// serve static files from public directory
	e.Static("/public", "public")

	// Routing = rute
	e.GET("/hello", helloWorld)
	e.GET("/about", about)
	e.GET("/", home)
	e.GET("/contact", contactMe)
	e.GET("/blog", blog)
	e.POST("/add-blog", addBlog)
	e.GET("/blog-detail/:id", blogdetail)
	e.GET("/delete-blog/:id", deleteBlog) // /delete-blog/0

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func about(c echo.Context) error {
	return c.String(http.StatusOK, "Ini adalah about")
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func contactMe(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/contact-me.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func blog(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	// map(tipe data) => key and value

	blogs := map[string]interface{}{
		"Blog": dataBlog,
	}

	return tmpl.Execute(c.Response(), blogs)
}

func blogdetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id")) // 2 string => 2 int

	tmpl, err := template.ParseFiles("views/blog-detail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	var BlogData = Blog{}
	for index, data := range dataBlog {
		if id == index {
			BlogData = Blog{
				Title:    data.Title,
				Content:  data.Content,
				PostDate: data.PostDate,
				Author:   data.Author,
			}
		}
	}

	data := map[string]interface{}{
		"Blog": BlogData,
	}

	return tmpl.Execute(c.Response(), data)
}

func addBlog(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("content")

	var addBlog = Blog{
		Title:    title,
		Content:  content,
		Author:   "Dandi Saputra",
		PostDate: time.Now().String(),
	}

	fmt.Println(addBlog)
	dataBlog = append(dataBlog, addBlog)

	// fmt.Println(title)
	// fmt.Println(content)
	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id")) // id = 0 string => 0 int

	dataBlog = append(dataBlog[:id], dataBlog[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}

// var id = 3
// var dataBlog = []string{"apple", "grape", "banana", "melon", "lemon", "mango"}
// dataBlog[:id] = "apple", "grape", "banana"
// dataBlog[id:] = "melon", "lemon", "mango"

// dataBlog[:id] = "apple", "grape", "banana"
// dataBlog[id+1:] = "lemon", "mango"
// ... = "apple", "grape", "banana", "lemon", "mango"
