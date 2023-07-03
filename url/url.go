package url

import (
	"github.com/gabyzanna/jadwal/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth

	page.Get("/", controller.Homepage)
	page.Get("/matkul", controller.GetMatakuliah)
	page.Get("/jadwal", controller.GetJadwalkuliah)
	page.Get("/kelas", controller.GetKelas) 
	page.Get("/dosen", controller.GetDosen)
	page.Get("/mhs", controller.GetMahasiswa)

	// page.Post("/", controller.Sink)
	// page.Put("/", controller.Sink)
	// page.Patch("/", controller.Sink)
	// page.Delete("/", controller.Sink)
	// page.Options("/", controller.Sink)

}