package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"personal-web/connection"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type SessionData struct {
	IsLogin bool
	Name    string
}

var userData = SessionData{}

func main() {
	connection.DatabaseConnect()

	// create new echo instance
	e := echo.New()

	// serve static files from public directory
	e.Static("/public", "public")

	// initialitation to use session
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

	// Routing = rute
	e.GET("/", home)
	e.GET("/contact", contactMe)
	e.GET("/blog", blog)
	e.GET("/blog-detail/:id", blogdetail)
	e.GET("/delete-blog/:id", deleteBlog) // /delete-blog/0
	e.GET("/form-register", formRegister)
	e.GET("/form-login", formLogin)
	e.GET("/logout", logout)
	e.POST("/add-blog", addBlog)
	e.POST("/register", register)
	e.POST("/login", login)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	sess, _ := session.Get("session", c)

	flash := map[string]interface{}{
		"FlashStatus":  sess.Values["isLogin"],
		"FlashMessage": sess.Values["message"],
		"FlashName":    sess.Values["name"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), flash)
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

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	blogs := map[string]interface{}{
		"Blog":        result,
		"DataSession": userData,
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

func formRegister(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/register.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func formLogin(c echo.Context) error {
	sess, _ := session.Get("session", c)
	flash := map[string]interface{}{
		"FlashStatus":  sess.Values["alertStatus"], // true / false
		"FlashMessage": sess.Values["message"],     // "Register success"
	}

	delete(sess.Values, "message")
	delete(sess.Values, "alertStatus")

	tmpl, err := template.ParseFiles("views/login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	return tmpl.Execute(c.Response(), flash)
}

func login(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	email := c.FormValue("email")
	password := c.FormValue("password")

	user := User{}
	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_user WHERE email=$1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		return redirectWithMessage(c, "Email Salah !", false, "/form-login")
	}

	fmt.Println(user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return redirectWithMessage(c, "Password Salah !", false, "/form-login")
	}

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800 // 3 jam
	sess.Values["message"] = "Login Success"
	sess.Values["status"] = true // show alert
	sess.Values["name"] = user.Name
	sess.Values["id"] = user.ID
	sess.Values["isLogin"] = true // access login
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func register(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	// generate password
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_user (name, email, password) VALUES ($1, $2, $3)", name, email, passwordHash)
	if err != nil {
		redirectWithMessage(c, "Register failed, please try again :)", false, "/form-register")
	}

	return redirectWithMessage(c, "Register success", true, "/form-login")
}

func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Values["isLogin"] = false
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func redirectWithMessage(c echo.Context, message string, status bool, path string) error {
	sess, _ := session.Get("session", c)
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, path)
}
