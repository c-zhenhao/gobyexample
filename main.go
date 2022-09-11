package main

import (
	"context"
	"fmt"

	"github.com/go-faker/faker/v4"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"math/rand"
	"time"
)

type Product struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Image       string             `json:"image,omitempty" bson:"image,omitempty"`
	Price       int                `json:"price,omitempty" bson:"price,omitempty"`
}

func main() {
	// to connect to database
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	db := client.Database("go_search")

	app := fiber.New()

	app.Use(cors.New())

	// test endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// endpoint to seed the database
	app.Post("/api/products/populate", func(c *fiber.Ctx) error {

		collection := db.Collection("products")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

		for i := 0; i < 50; i++ {
			collection.InsertOne(ctx, Product{
				Title:       faker.Word(),
				Description: faker.Paragraph(),
				Image:       fmt.Sprintf("http://lorempixel.com/200/200?%s", faker.UUIDDigit()),
				Price:       rand.Intn(90) + 10,
			})
		}

		return c.JSON(fiber.Map{
			"message": "success",
		})
	})

	app.Get("/api/products/frontend", func(c *fiber.Ctx) error {
		collection := db.Collection("products")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

		// create a product slice
		var products []Product

		cursor, _ := collection.Find(ctx, bson.M{})
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var product Product
			cursor.Decode(&product)
			products = append(products, product)
		}

		return c.JSON(products)
	})

	app.Get("/api/products/backend", func(c *fiber.Ctx) error {
		collection := db.Collection("products")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

		var products []Product

		filter := bson.M{}

		if s := c.Query("search"); s != "" {
			filter = bson.M{
				"title": bson.M{
					"$regex": primitive.Regex{
						Pattern: s,
						Options: "i",
					},
				},
			}
		}

		cursor, _ := collection.Find(ctx, filter)
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var product Product
			cursor.Decode(&product)
			products = append(products, product)
		}

		return c.JSON(products)
	})

	app.Listen(":8000")
}
