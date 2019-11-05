package login

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"
	_ "time"

	"zdxsv/pkg/assets"
	"zdxsv/pkg/config"
	"zdxsv/pkg/db"

	"github.com/axgle/mahonia"
	"github.com/golang/glog"
)

var (
	tplTop     *template.Template
	tplLogin   *template.Template
	tplMessage *template.Template
	se         mahonia.Encoder
)

func Prepare() {
	var err error
	var bin []byte

	bin, err = assets.Asset("assets/top.tpl")
	if err != nil {
		glog.Fatalln(err)
	}

	tplTop, err = template.New("top").Parse(string(bin))
	if err != nil {
		glog.Fatalln(err)
	}

	bin, err = assets.Asset("assets/login.tpl")
	if err != nil {
		glog.Fatalln(err)
	}

	tplLogin, err = template.New("login").Parse(string(bin))
	if err != nil {
		glog.Fatalln(err)
	}

	bin, err = assets.Asset("assets/message.tpl")
	if err != nil {
		glog.Fatalln(err)
	}

	tplMessage, err = template.New("message").Parse(string(bin))
	if err != nil {
		glog.Fatalln(err)
	}

	se = mahonia.NewEncoder("Shift_JIS")
}

var messageRegister = `<br><br><br><br><br>
	用戸帳號創建完成。 <br>
	<br>
	儲存 ID 至 Memory Card 後返回。<br>`
var messageLoginFail = `<br><br><br><br><br>
	登入失敗。<br>
	<br>
	如果還未建立帳號請返回上頁註冊。<br>`
var messageMainte = `<br><br><br><br><br><br>
   系統目前正在維護中。<br>
   <br>
   請稍後再嘗試登入。<br>`

type commonParam struct {
	ServerVersion string
	LoginKey      string
	SessionId     string
	ServerAddr    string
	Message       string
}

func HandleTopPage(w http.ResponseWriter, r *http.Request) {
	p := commonParam{}
	p.ServerVersion = "1.0"
	w.Header().Set("Content-Type", "text/html; charset=cp932")
	w.WriteHeader(200)
	sw := se.NewWriter(w)
	tplTop.Execute(sw, p)
}

func HandleLoginPage(w http.ResponseWriter, r *http.Request) {
	p := commonParam{}
	r.ParseForm()
	glog.Infoln(r.Form)
	loginKey := r.FormValue("login_key")

	if loginKey == "" {
		w.Header().Set("Content-Type", "text/html; charset=cp932")
		w.WriteHeader(200)
		writeMessagePage(w, r, messageLoginFail)
		return
	}

	a, err := db.DefaultDB.GetAccountByLoginKey(loginKey)
	if err == sql.ErrNoRows && len(loginKey) == 10 {
		// Since this login key seems to have been registered on another server,
		// new registration is performed.
		a, err = db.DefaultDB.RegisterAccountWithLoginKey(r.RemoteAddr, loginKey)
	}

	if err != nil {
		glog.Errorln(err)
		w.Header().Set("Content-Type", "text/html; charset=cp932")
		w.WriteHeader(200)
		writeMessagePage(w, r, messageLoginFail)
		return
	}

	err = db.DefaultDB.LoginAccount(a)

	if err != nil {
		glog.Errorln(err)
		w.Header().Set("Content-Type", "text/html; charset=cp932")
		w.WriteHeader(200)
		writeMessagePage(w, r, messageLoginFail)
		return
	}

	p.ServerVersion = "1.0"
	p.LoginKey = a.LoginKey
	p.SessionId = a.SessionId
	p.ServerAddr = config.Conf.Lobby.PublicAddr

	w.Header().Set("Content-Type", "text/html; charset=cp932")
	w.WriteHeader(200)
	sw := se.NewWriter(w)
	tplLogin.Execute(sw, p)
}

func HandleRegisterPage(w http.ResponseWriter, r *http.Request) {
	a, err := db.DefaultDB.RegisterAccount(r.RemoteAddr)
	sw := se.NewWriter(w)
	if err != nil {
		glog.Errorln(err)
		w.Header().Set("Content-Type", "text/html; charset=cp932")
		w.WriteHeader(200)
		writeMessagePage(w, r, "註冊失敗")
	}
	w.Header().Set("Content-Type", "text/html; charset=cp932")
	w.WriteHeader(200)
	fmt.Fprintf(sw, "<!--COMP-SIGNUP--><!--INPUT-IDS   %s-->\n", a.LoginKey)
	writeMessagePage(w, r, messageRegister)
}

func writeMessagePage(w http.ResponseWriter, r *http.Request, message string) {
	p := commonParam{}
	p.ServerVersion = "1.0"
	p.Message = message
	sw := se.NewWriter(w)
	tplMessage.Execute(sw, p)
}
