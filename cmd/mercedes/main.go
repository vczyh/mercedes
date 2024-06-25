package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
	"github.com/vczyh/mercedes"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "mercedes",
		Usage: "Mercedes Cli",
		Action: func(*cli.Context) error {
			fmt.Println("https://github.com/vczyh/mercedes")
			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"v"},
				Usage:   "Debug message",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "login",
				Usage: "Login to Mercedes",
				Action: func(c *cli.Context) error {
					email := c.Args().First()
					if email == "" {
						return fmt.Errorf("email is required")
					}
					debug := c.IsSet("debug")
					fmt.Println(email)
					ctx := context.Background()
					api := mercedes.NewAPI(
						mercedes.WithAPIRegion(mercedes.RegionChina),
					)
					configRes, err := api.Config(ctx)
					if err != nil {
						return err
					}
					if debug {
						fmt.Printf("Config Response: %+v\n", configRes)
					}
					nonce := uuid.New().String()
					loginRes, err := api.Login(ctx, email, nonce)
					if err != nil {
						return err
					}
					if debug {
						fmt.Printf("Login Response: %+v\n", loginRes)
					}
					fmt.Println("Email has been sent, please enter your code: ")
					var code string
					if _, err := fmt.Scanln(&code); err != nil {
						return err
					}
					if debug {
						fmt.Printf("Your code: %s\n", code)
					}
					oauth2Res, err := api.OAuth2(ctx, email, nonce, code)
					if err != nil {
						return err
					}
					if debug {
						fmt.Printf("OAuth2 Response: %+v\n", oauth2Res)
					}
					fmt.Printf("Access Token: %s\n", oauth2Res.AccessToken)
					fmt.Printf("Refresh Token: %s\n", oauth2Res.RefreshToken)
					fmt.Printf("Expires In: %d\n", oauth2Res.ExpiresIn)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
