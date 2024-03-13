package controller

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/btangkas-client/config"
)

type responsechecktoken struct {
	Client_status                   int         `json:"status"`
	Client_company                  string      `json:"client_company"`
	Client_name                     string      `json:"client_name"`
	Client_username                 string      `json:"client_username"`
	Client_credit                   int         `json:"client_credit"`
	Engine_multiplier_angka         float32     `json:"engine_multiplier_angka"`
	Engine_multiplier_redblack      float32     `json:"engine_multiplier_redblack"`
	Engine_multiplier_line          float32     `json:"engine_multiplier_line"`
	Engine_status_game_redblackline string      `json:"engine_status_game_redblackline"`
	Engine_status_maintenance       string      `json:"engine_status_maintenance"`
	Client_listbet                  interface{} `json:"client_listbet"`
}
type responsesavetransaksi struct {
	Client_status      int    `json:"status"`
	Client_message     string `json:"message"`
	Client_idtransaksi string `json:"idtransaksi"`
	Client_card_game   string `json:"card_game"`
	Client_card_length int    `json:"card_length"`
}
type responsedefault struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
}
type responseerror struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var PATH string = config.GetPathAPI()

func CheckToken(c *fiber.Ctx) error {
	type payload_checktoken struct {
		Client_token string `json:"client_token" `
	}
	hostname := c.Hostname()
	client := new(payload_checktoken)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsechecktoken{}).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"token": client.Client_token,
		}).
		Post(PATH + "api/checktoken")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsechecktoken)
	if result.Client_status == 200 {
		return c.JSON(fiber.Map{
			"status":                          result.Client_status,
			"client_name":                     result.Client_name,
			"client_username":                 result.Client_username,
			"client_credit":                   result.Client_credit,
			"engine_multiplier_angka":         result.Engine_multiplier_angka,
			"engine_multiplier_redblack":      result.Engine_multiplier_redblack,
			"engine_multiplier_line":          result.Engine_multiplier_line,
			"engine_status_game_redblackline": result.Engine_status_game_redblackline,
			"engine_status_maintenance":       result.Engine_status_maintenance,
			"client_company":                  result.Client_company,
			"client_listbet":                  result.Client_listbet,
			"time":                            time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Listresult(c *fiber.Ctx) error {
	type payload_listinvoice struct {
		Invoice_company string `json:"invoice_company" `
	}
	hostname := c.Hostname()
	client := new(payload_listinvoice)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"invoice_company": client.Invoice_company,
		}).
		Post(PATH + "api/listresult")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Listinvoice(c *fiber.Ctx) error {
	type payload_listinvoice struct {
		Invoice_company  string `json:"invoice_company" `
		Invoice_username string `json:"invoice_username" `
	}
	hostname := c.Hostname()
	client := new(payload_listinvoice)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"invoice_company":  client.Invoice_company,
			"invoice_username": client.Invoice_username,
		}).
		Post(PATH + "api/listinvoice")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func SaveTransaksiDetail(c *fiber.Ctx) error {
	type payload_savetransaksidetail struct {
		Transaksidetail_company     string          `json:"transaksidetail_company" `
		Transaksidetail_idtransaksi string          `json:"transaksidetail_idtransaksi" `
		Transaksidetail_username    string          `json:"transaksidetail_username" `
		Transaksidetail_totalbet    int             `json:"transaksidetail_totalbet" `
		Transaksidetail_listdatabet json.RawMessage `json:"transaksidetail_listdatabet" `
	}
	hostname := c.Hostname()
	client := new(payload_savetransaksidetail)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"transaksidetail_company":     client.Transaksidetail_company,
			"transaksidetail_idtransaksi": client.Transaksidetail_idtransaksi,
			"transaksidetail_username":    client.Transaksidetail_username,
			"transaksidetail_totalbet":    client.Transaksidetail_totalbet,
			"transaksidetail_listdatabet": string(client.Transaksidetail_listdatabet),
		}).
		Post(PATH + "api/savetransaksidetail")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
