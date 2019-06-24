package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
}

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"username"`
	HashedPass string `json:"hashed_pass"`
	Profile    string `json:"profile"`
	ImgName    string `json:"img_name"`
}

type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func postSignUpHandler(c echo.Context) error {
	req := LoginRequestBody{}
	c.Bind(&req)

	//Validation
	if req.Password == "" || req.Username == "" {
		return c.String(http.StatusBadRequest, "項目が空です")
	}
	if req.Password <= 6 {
		return c.String(http.StatusBadRequest, "パスワードが短すぎます")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("bcrypt generate error: %v", err))
	}

	var count int
	row := Db.QueryRow("SELECT id FROM users WHERE name = $1", req.Username).Scan(&count)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
	}
	if count > 0 {
		return c.String(http.StatusConflict, "ユーザーがすでに存在しています")
	}

	_, err = Db.Exec("INSERT INTO users (name, hashed_pass) values ($1, $2)", req.Username, hashedPass)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("db err: %v", err))
	}

	return c.NoContent(http.StatusCreated)
}

func postLoginHandler(c echo.Context) error {
	req := LoginRequestBody{}
	c.Bind(&req)

	user := User{}
	err := Db.QueryRow("SELECT hashed_pass FROM users WHERE name = $1", req.Username).Scan(&user.HashedPass)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPass), []byte(req.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return c.NoContent(http.StatusForbidden)
		} else {
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	sess, err := session.Get("sessions", c)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "something wrong in getting session")
	}
	sess.Values["userName"] = req.Username
	sess.Save(c.Request(), c.Response())

	return c.NoContent(http.StatusOK)

}

func checkLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("sessions", c)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "something wrong in getting session")
		}

		if sess.Values["userName"] == nil {
			return c.String(http.StatusForbidden, "please login")
		}
		c.Set("userName", sess.Values["userName"].(string))

		return next(c)
	}
}

func main() {
	e := echo.New()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte(secret))))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "front/dist/")
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})
	e.POST("/login", postLoginHandler)
	e.POST("/signup", postSignUpHandler())

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
