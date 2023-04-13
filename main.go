package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

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

	e.Logger.Fatal(e.Start("localhost:8000"))
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func about(c echo.Context) error {
	return c.String(http.StatusOK, "Ini adalah about")
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.ht")

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

	return tmpl.Execute(c.Response(), nil)
}

func blogdetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id")) // 123 string => 123 int

	tmpl, err := template.ParseFiles("views/blog-detail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	data := map[string]interface{}{
		"Id":      id,
		"Title":   "Pasar Coding di Indonesia Dinilai Masih Menjanjikan",
		"Author":  "Dandi Saputra",
		"Content": "REPUBLIKA.CO.ID, JAKARTA -- Ketimpangan sumber daya manusia (SDM) di sektor digital masih menjadi isu yang belum terpecahkan. Berdasarkan penelitian ManpowerGroup, ketimpangan SDM global, termasuk Indonesia, meningkat dua kali lipat dalam satu dekade terakhir. Khusus di sektor teknologi yang berkembang pesat, menurut Kemendikbudristek, kekurangan sembilan juta pekerja teknologi hingga tahun 2030. Hal itu berarti Indonesia memerlukan sekitar 600 ribu SDM digital memasuki pasar setiap tahunnya.",
	}

	return tmpl.Execute(c.Response(), data)
}

func addBlog(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("content")

	fmt.Println(title)
	fmt.Println(content)

	return c.Redirect(http.StatusMovedPermanently, "/")
}
