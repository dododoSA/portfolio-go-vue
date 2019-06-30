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
	Id         int            `json:"id"`
	Name       string         `json:"username"`
	HashedPass string         `json:"hashed_pass"`
	Profile    sql.NullString `json:"profile"`
	ImgName    sql.NullString `json:"img_name"`
}

type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Product struct {
	Id      int            `json:"id"`
	Name    string         `json:"productname"`
	Intro   string         `json:"intro"`
	ImgName sql.NullString `json:"img_name"`
	Url     sql.NullString `json:"url"`
	UserId  int            `json:"user_id"`
}

func postSignUpHandler(c echo.Context) error {
	req := LoginRequestBody{}
	c.Bind(&req)

	//Validation
	if req.Password == "" || req.Username == "" {
		return c.String(http.StatusBadRequest, "項目が空です")
	}
	if len(req.Password) <= 6 {
		return c.String(http.StatusBadRequest, "パスワードが短すぎます")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("bcrypt generate error: %v", err))
	}

	var count int
	err = Db.QueryRow("SELECT COUNT(*) FROM users WHERE name = $1", req.Username).Scan(&count)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("db errorA: %v", err))
	}
	if count > 0 {
		return c.String(http.StatusConflict, "ユーザーがすでに存在しています")
	}

	_, err = Db.Exec("INSERT INTO users (name, hashed_pass) values ($1, $2)", req.Username, hashedPass)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("db errB: %v", err))
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

func getProductsHandler(c echo.Context) error {
	userId := c.Param("id")

	products := make([]Product, 0)
	rows, err := Db.Query("SELECT * FROM products WHERE user_id = $1", userId)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("db errorA: %v", err))
	}
	defer rows.Close()
	for rows.Next() {
		product := Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.Intro, &product.ImgName, &product.Url, &product.UserId)
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("db errorA: %v", err))
		}
		products = append(products, product)
	}
	return c.JSON(http.StatusOK, products)
}

func getUserHandler(c echo.Context) error {
	userId := c.Param("id")
	user := User{}
	err := Db.QueryRow("SELECT profile, img_name, name FROM users WHERE id = $1", userId).Scan(&user.Profile, &user.ImgName, &user.Name)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("db errorA: %v", err))
	}
	return c.JSON(http.StatusOK, user)
}

type Me struct {
	UserId string `json:"username"`
}

func getWhoAmIHandler(c echo.Context) error {
	userName := c.Get("userName").(string)
	if userName == "" {
		return c.String(http.StatusForbidden, "please login")
	}
	me := Me{}
	err := Db.QueryRow("SELECT id FROM users WHERE name = $1", userName).Scan(&me.UserId)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("db errorA: %v", err))
	}
	return c.JSON(http.StatusOK, me)
}

func main() {
	e := echo.New()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "front/dist/")
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})
	e.GET("/users/:id/products", getProductsHandler)
	e.GET("/users/:id", getUserHandler)
	e.GET("/whoami", getWhoAmIHandler)
	e.POST("/login", postLoginHandler)
	e.POST("/signup", postSignUpHandler)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
