package main

import (
	"auth/commands"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	token := os.Getenv("DISCORD_TOKEN")

	if len(token) == 0 {
		fmt.Println("DISCORD_TOKEN is not set!")
		return
	}

	session, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Logged in as " + r.User.String())
	})

	session.AddHandler(commands.Logs)

	session.AddHandler(commands.Handler)

	err = session.Open()
	if err != nil {
		fmt.Println("Error opening Discord session,", err)
		return
	}

	fmt.Println("Bot is now running.")

	select {}
}
