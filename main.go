package main

import (
	"github.com/go-zoox/cli"
)

func main() {
	app := cli.NewSingleProgram(&cli.SingleProgramConfig{
		Name:    "chatgpt-for-feishu-bot",
		Usage:   "chatgpt-for-feishu-bot is a portable chatgpt server",
		Version: Version,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Usage:   "server port",
				Aliases: []string{"p"},
				EnvVars: []string{"PORT"},
				Value:   8080,
			},
			&cli.StringFlag{
				Name:     "chatgpt-api-key",
				Usage:    "ChatGPT API Key",
				EnvVars:  []string{"CHATGPT_API_KEY"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "app-id",
				Usage:    "Feishu App ID",
				EnvVars:  []string{"APP_ID"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "app-secret",
				Usage:    "Feishu App SECRET",
				EnvVars:  []string{"APP_SECRET"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "encrypt-key",
				Usage:   "enable encryption if you need",
				EnvVars: []string{"ENCRYPT_KEY"},
			},
			&cli.StringFlag{
				Name:    "verification-token",
				Usage:   "enable token verification if you need",
				EnvVars: []string{"VERIFICATION_TOKEN"},
			},
			&cli.BoolFlag{
				Name:    "challenge",
				Usage:   "server as challenge",
				EnvVars: []string{"CHALLENGE"},
			},
		},
	})

	app.Command(func(ctx *cli.Context) (err error) {
		return ServeFeishuBot(&FeishuBotConfig{
			Port:              ctx.Int64("port"),
			ChatGPTAPIKey:     ctx.String("chatgpt-api-key"),
			AppID:             ctx.String("app-id"),
			AppSecret:         ctx.String("app-secret"),
			EncryptKey:        ctx.String("encrypt-key"),
			VerificationToken: ctx.String("verification-token"),
			AsChallenge:       ctx.Bool("challenge"),
		})
	})

	app.Run()
}
