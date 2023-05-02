package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"personal-web/connection"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Blog struct {
	ID       int
	Title    string
	Content  string
	Image    string
	Author   string
	PostDate time.Time
}

type Experience struct {
	ID      int
	Project string
	Year    int
}

// var dataBlog = []Blog{
// 	{
// 		Title:   "Pasar coding dinilai masih menjanjikan",
// 		Content: "Ketimpangan sumber daya manusia (SDM) di sektor digital masih menjadi isu yang belum terpecahkan. Berdasarkan penelitian ManpowerGroup, ketimpangan SDM global, termasuk Indonesia, meningkat dua kali lipat dalam satu dekade terakhir. Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quam, molestiae numquam! Deleniti maiores expedita eaque deserunt quaerat! Dicta, eligendi debitis?",
// 		// Author:  "Jaya Saleh",
// 		// PostDate: "14/04/2023",
// 	},
// 	{
// 		Title:   "Pasar coding dinilai masih sedikit",
// 		Content: "Ketimpangan sumber daya manusia (SDM) di sektor digital masih menjadi isu yang belum terpecahkan. Berdasarkan penelitian ManpowerGroup, ketimpangan SDM global, termasuk Indonesia, meningkat dua kali lipat dalam satu dekade terakhir. Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quam, molestiae numquam! Deleniti maiores expedita eaque deserunt quaerat! Dicta, eligendi debitis?",
// 		Author:  "Yoga Wicaksono",
// 		// PostDate: "15/04/2023",
// 	},
// 	{
// 		Title:   "Pasar coding dinilai masih sedikit",
// 		Content: "Ketimpangan sumber daya manusia (SDM) di sektor digital masih menjadi isu yang belum terpecahkan. Berdasarkan penelitian ManpowerGroup, ketimpangan SDM global, termasuk Indonesia, meningkat dua kali lipat dalam satu dekade terakhir. Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quam, molestiae numquam! Deleniti maiores expedita eaque deserunt quaerat! Dicta, eligendi debitis?",
// 		Author:  "Yoga Wicaksono",
// 		// PostDate: "15/04/2023",
// 	},
// }

func main() {
	connection.DatabaseConnect()

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

	data, _ := connection.Conn.Query(context.Background(), "SELECT * FROM public.tb_experience")

	var exp []Experience

	for data.Next() {
		var each = Experience{}

		err := data.Scan(&each.ID, &each.Project, &each.Year)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		exp = append(exp, each)
	}

	fmt.Println(exp)
	experiences := map[string]interface{}{
		"Experience": exp,
	}

	return tmpl.Execute(c.Response(), experiences)
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
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, title, post_date, content, image FROM tb_blog")
	fmt.Println(data)

	var result []Blog

	for data.Next() {
		var each = Blog{}

		err := data.Scan(&each.ID, &each.Title, &each.PostDate, &each.Content, &each.Image)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		each.Author = "Dandi Saputra"

		result = append(result, each)
	}

	blogs := map[string]interface{}{
		"Blog": result,
	}

	return tmpl.Execute(c.Response(), blogs)
}

func blogdetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id")) // 2 string => 2 int

	tmpl, err := template.ParseFiles("views/blog-detail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	var BlogDetail = Blog{}

	err = connection.Conn.QueryRow(context.Background(), "SELECT id, title, content, image, post_date FROM tb_blog WHERE id = $1", id).Scan(&BlogDetail.ID, &BlogDetail.Title, &BlogDetail.Content, &BlogDetail.Image, &BlogDetail.PostDate)

	BlogDetail.Author = "Jery"

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	data := map[string]interface{}{
		"Blog": BlogDetail,
	}

	return tmpl.Execute(c.Response(), data)
}

func addBlog(c echo.Context) error {
	title := c.FormValue("title")     // Batch 46
	content := c.FormValue("content") // Finishing CRUD
	image := "image.png"

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_blog(title, content, image, post_date) VALUES ($1, $2, $3, $4)", title, content, image, time.Now())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}

func deleteBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id")) // id = 1 string => 1 int

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_blog WHERE id=$1", id)

	// DELETE FROM tb_blog WHERE id=1
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}

// var id = 3
// var dataBlog = []string{"apple", "grape", "banana", "melon", "lemon", "mango"}
// dataBlog[:id] = "apple", "grape", "banana"
// dataBlog[id:] = "melon", "lemon", "mango"

// dataBlog[:id] = "apple", "grape", "banana"
// dataBlog[id+1:] = "lemon", "mango"
// ... = "apple", "grape", "banana", "lemon", "mango"
