package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	coingecko "github.com/superoo7/go-gecko/v3"
	"net/http"
	"os"
	"time"
)

type Config struct {
	Frequency int `mapstructure:"frequency"`
	Timeout   int `mapstructure:"timeout"`
	Notify    struct {
		Discord struct {
			Token     string `mapstructure:"token"`
			ChannelID string `mapstructure:"channelId"`
		} `mapstructure:"discord"`
	} `mapstructure:"notify"`
}

type Bot struct {
	Config *Config
	CG     *coingecko.Client
}

func NewBot(c *Config) *Bot {
	client := &http.Client{Timeout: time.Duration(c.Timeout) * time.Millisecond}
	CG := coingecko.NewClient(client)
	return &Bot{CG: CG, Config: c}
}

func (b *Bot) Run() {
	interval := time.Duration(b.Config.Frequency) * time.Millisecond
	schedule(b.onTick, interval)
	discord, err := discordgo.New("Bot " + b.Config.Notify.Discord.Token)
	if err != nil {
		println("err", err.Error())
		os.Exit(1)
	}
	cha, err := discord.Channel(b.Config.Notify.Discord.ChannelID)
	checkErr(err)

	msg, err := discord.ChannelMessageSend(cha.ID, "Hlelo")
	checkErr(err)
	println(msg.Content)
}

func checkErr(err error) {
	if err != nil {
		println("err: ", err.Error())
		os.Exit(1)
	}
}

func (b *Bot) onTick() {
	print("tick\n")
	coin, err := b.CG.SimpleSinglePrice("ethereum", "usd")
	if nil != err {
		println("err", err.Error())
		return
	}
	fmt.Printf("price: %.2f\n", coin.MarketPrice)
}
