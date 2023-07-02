package controller

import (
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/whatsauth/whatsauth"
	"github.com/gabyzanna/jadwal/config"
)

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

func GetHome(c *fiber.Ctx) error {
	getip := musik.GetIPaddress()
	return c.JSON(getip)
}

func GetTam(c *fiber.Ctx) error {
	getkot := bella.GetDataListTamu("BOGOR")
	return c.JSON(getkot)
}

func GetDataUndanganRapat(c *fiber.Ctx) error {
	getun := bella.GetDataUndanganRapat("Rapat Umum")
	return c.JSON(getun)
}

func GetDataPesertaRapat(c *fiber.Ctx) error {
	getpes := bella.GetDataPesertaRapat("ULBI")
	return c.JSON(getpes)
}

func GetDataWaktuRapat(c *fiber.Ctx) error {
	getwa := bella.GetDataWaktuRapat("Generasi Muda")
	return c.JSON(getwa)
}

func GetDataRapatMulai(c *fiber.Ctx) error {
	getra := bella.GetDataRapatMulai("Jokowi")
	return c.JSON(getra)
}