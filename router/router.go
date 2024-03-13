package router

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nikitamirzani323/btangkas-client/controller"
	"github.com/valyala/fasthttp"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(compress.New())
	// Custom config
	app.Static("/", "frontend/dist", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})
	app.Get("api/healthz", controller.HealthCheck)
	app.Post("api/checktoken", controller.CheckToken)
	app.Post("api/listresult", controller.Listresult)
	app.Post("api/listinvoice", controller.Listinvoice)
	app.Post("api/savetransaksi", controller.SaveTransaksiDetail)

	app.Get("/sse", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/event-stream")
		c.Set("Cache-Control", "no-cache")
		c.Set("Connection", "keep-alive")
		c.Set("Transfer-Encoding", "chunked")

		c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
			fmt.Println("WRITER")

			dbHost := os.Getenv("DB_REDIS_HOST") + ":" + os.Getenv("DB_REDIS_PORT")
			dbPass := os.Getenv("DB_REDIS_PASSWORD")
			dbName, _ := strconv.Atoi(os.Getenv("DB_REDIS_NAME"))

			rdb := redis.NewClient(&redis.Options{
				Addr:     dbHost,
				Password: dbPass,
				DB:       dbName,
			})

			resultredis := rdb.Subscribe("", "payload_nuke")
			defer resultredis.Close()
			for {
				msg, err := resultredis.ReceiveMessage()
				if err != nil {
					panic(err)
				}

				// fmt.Println("Received message from " + msg.Payload + " channel.")
				// data_pubsub := strings.Split(msg.Payload, ":")

				msg_sse := msg.Payload

				fmt.Fprintf(w, "data: %s\n\n", msg_sse)
				// fmt.Println(msg_sse)
				err_sse := w.Flush()
				if err_sse != nil {
					fmt.Printf("Error while flushing: %v. Closing http connection.\n", err)

					break
				}
				// time.Sleep(1 * time.Second)

			}

		}))

		return nil
	})
	return app
}
