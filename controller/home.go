package controller

import (
	"encoding/json"
	"fmt"
	"time"

	"bitbucket.org/isbtotogroup/wigo_client_game_12digit/config"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
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
type responsebalance struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Credit  int    `json:"credit"`
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

	fmt.Println("Hostname: ", hostname)
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
		fmt.Println(err.Error())
	}
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
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
func Balance(c *fiber.Ctx) error {
	type payload_listinvoice struct {
		Client_token string `json:"client_token" `
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

	fmt.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsebalance{}).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_token": client.Client_token,
		}).
		Post(PATH + "api/balance")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
	result := resp.Result().(*responsebalance)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"credit":  result.Credit,
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
func Listresult(c *fiber.Ctx) error {
	type payload_listinvoice struct {
		Client_token string `json:"client_token" `
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

	fmt.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_token": client.Client_token,
		}).
		Post(PATH + "api/listresult")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
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
		Client_token string `json:"client_token" `
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

	fmt.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_token": client.Client_token,
		}).
		Post(PATH + "api/listinvoice")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
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
		Client_token                string          `json:"client_token" `
		Transaksidetail_idtransaksi string          `json:"transaksidetail_idtransaksi" `
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

	fmt.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_token":                client.Client_token,
			"transaksidetail_idtransaksi": client.Transaksidetail_idtransaksi,
			"transaksidetail_totalbet":    client.Transaksidetail_totalbet,
			"transaksidetail_listdatabet": string(client.Transaksidetail_listdatabet),
		}).
		Post(PATH + "api/savetransaksidetail")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
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
