## Table of Contents

- [Table of Contents](#table-of-contents)
- [Description](#description)
- [Social networks integrations](#social-networks-integrations)
  - [Telegram](#telegram)
  - [bot commands](#bot-commands)
  - [Twitter](#twitter)
  - [Whatsapp](#whatsapp)
- [How to run](#how-to-run)
  - [Requirements](#requirements)
  - [Clone repo](#clone-repo)
  - [Config](#config)
  - [OPTIONAL VARIABLES](#optional-variables)
  - [Deploy](#deploy)
- [Contribution](#contribution)
  - [run](#run)
  - [build](#build)


## Description

An application that allows you to take content generated in [chappie_server](https://github.com/Think-Root/chappie_server) and publish it in the social networks like telegram/twitter/whatsapp.

I once had the idea to create an app that would manage a Telegram channel by searching for and posting interesting repositories using AI. That idea eventually grew into this project. You can read more details about how the idea was born here: [Part 1](https://drukarnia.com.ua/articles/yak-chatgpt-vede-za-mene-kanal-v-telegram-i-u-nogo-ce-maizhe-vikhodit-chastina-1-VywRW) and [Part 2](https://drukarnia.com.ua/articles/yak-chatgpt-vede-za-mene-kanal-v-telegram-i-u-nogo-ce-maizhe-vikhodit-chastina-2-X9Yjz). (The articles are in Ukrainian, but I think you’ll manage to translate them into your preferred language.)

## Social networks integrations

Currently, the bot supports 3 integrations, two official (telegram, twitter) and one unofficial (whatsapp)

### Telegram

Publishes posts to a Telegram channel, allows managing some functions through a Telegram bot

### bot commands

- **/add** - you can provide a list (or just one) of GitHub repository URLs separated by spaces, a short description will be generated for each, and they will be added to the database.
- **/next** - displays the next post scheduled for publication in the channel, if you provide a number after a space, that many posts will be shown
- **/gen** - generates a post for the channel from GitHub trends, if you specify a number after a space, that many posts will be generated and added to the database
- **/info** - shows information about the posts in the database

### Twitter

This integration allows you to publish posts on your Twitter account. For operation, it requires deploying the [think-root/x](https://github.com/Think-Root/x) app, which implements the official Twitter API in the form of a small server.


### Whatsapp

This integration allows you to publish posts to a WhatsApp channel. The integration is not official, and its use may **lead to your WhatsApp account being banned**. To use this integration, you will need to deploy the [think-root/wapp](https://github.com/Think-Root/wapp) app

## How to run

### Requirements

- [docker](https://docs.docker.com/engine/install/) or/and [docker-compose](https://docs.docker.com/compose/install/)

### Clone repo

```shell
git clone https://github.com/Think-Root/chappie_bot.git
```

### Config

Create a **.env** file in the app root directory

```properties
BOT_TOKEN=<bot token https://core.telegram.org/bots>
ADMIN_ID=<your telegram id>
CHANNEL_ID=<your channel id>
CHAPPIE_SERVER_URL=<e.g. http://localhost:8080/>
CHAPPIE_SERVER_BEARER=<chappie_server server token>
X_API_KEY=<api key for selfhosted x api server https://github.com/Think-Root/x>
X_URL=<url for selfhosted x api server https://github.com/Think-Root/x>
```

### OPTIONAL VARIABLES

> 🔴 IF YOU ARE NOT READY TO RISK LOSING YOUR WHATSAPP ACCOUNT, DO NOT SET WAPP_ VARIABLES IN THE ENV FILE

Watch this [repository](https://github.com/Think-Root/wapp)

```properties
ENABLE_CRON=<true/false to enable CollectPostsCron, default is false>
WAPP_ENABLE=<true/false to enable/disable WhatsApp functionality, default is false>
WAPP_JID=<whatsapp chat jid>
WAPP_TOKEN=<whatsapp app token>
WAPP_SERVER_URL<server url for>
```

### Deploy

- deploy [this](https://github.com/Think-Root/chappie_server?tab=readme-ov-file#deploy) app
- build `docker build -t chappie_bot:latest -f Dockerfile .`
- run `docker run --name chappie_bot --restart always --env-file .env -e TZ=Europe/Kiev --network chappie_network chappie_bot:latest`
- or via docker compose `docker compose up -d`

## Contribution

- install [go](https://go.dev/dl/)

### run
```shell
 go run ./cmd/bot/main.go  
```

### build
```shell
go build -o chappie_bot ./cmd/bot/main.go
```