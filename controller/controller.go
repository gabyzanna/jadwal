package controller

import (
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/gabyzanna/jadwal/config"
	"github.com/whatsauth/whatsauth"
	"github.com/gabyzanna/bella"
	
	//kmmdl "github.com/gocroot/kampus/model"
	//kampus "github.com/gocroot/kampus/module"
)

// var Matkul1 = "matakuliah"
// var Jdwal1 = "jadwalkuliah"
// var Kls1 = "kelas"
// var Dsn1 = "dosen"
// var Mhs1 = "mahasiswa"

type HTTPRequest struct {
	Header string `json:"header"`
	Body   string `json:"body"`
}

func Sink(c *fiber.Ctx) error {
	var req HTTPRequest
	req.Header = string(c.Request().Header.Header())
	req.Body = string(c.Request().Body())
	return c.JSON(req)
}

func WsWhatsAuthQR(c *websocket.Conn) {
	whatsauth.RunSocket(c, config.PublicKey, config.Usertables[:], config.Ulbimariaconn)
}

func PostWhatsAuthRequest(c *fiber.Ctx) error {
	if string(c.Request().Host()) == config.Internalhost {
		var req whatsauth.WhatsauthRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}
		ntfbtn := whatsauth.RunModuleLegacy(req, config.PrivateKey, config.Usertables[:], config.Ulbimariaconn)
		return c.JSON(ntfbtn)
	} else {
		var ws whatsauth.WhatsauthStatus
		ws.Status = string(c.Request().Host())
		return c.JSON(ws)
	}

}

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetMatakuliah(c *fiber.Ctx) error {
	ps := bella.GetDataMatakuliah("TI41278")
	return c.JSON(ps)
}

func GetJadwalkuliah(c *fiber.Ctx) error {
	ps := bella.GetDataJadwalkuliah("Senin")
	return c.JSON(ps)
}

func GetKelas(c *fiber.Ctx) error {
	ps := bella.GetDataKelas("310")
	return c.JSON(ps)
}

func GetDosen(c *fiber.Ctx) error {
	ps := bella.GetDataDosen("NOVIANA RIZA")
	return c.JSON(ps)
}

func GetMahasiswa(c *fiber.Ctx) error {
	ps := bella.GetDataMahasiswa("ARDIVA PUTRI TAVA PRAMESWARI")
	return c.JSON(ps)
}