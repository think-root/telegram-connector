## Table of Contents

- [Table of Contents](#table-of-contents)
- [Description](#description)
- [How to run](#how-to-run)
  - [Requirements](#requirements)
  - [Clone repo](#clone-repo)
  - [Config](#config)
  - [Deploy (If deploy before deploying the chappie\_server)](#deploy-if-deploy-before-deploying-the-chappie_server)
  - [Deploy (If deploy after deploying the chappie\_server)](#deploy-if-deploy-after-deploying-the-chappie_server)
- [Telegram bot commands](#telegram-bot-commands)
- [Contribution](#contribution)
  - [run](#run)
  - [build](#build)


## Description

An application that allows you to take content generated in [chappie_server](https://github.com/Think-Root/chappie_server) and publish it in a Telegram channel (and [not only](https://github.com/Think-Root/wapp))

I once had the idea to create an app that would manage a Telegram channel by searching for and posting interesting repositories using AI. That idea eventually grew into this project. You can read more details about how the idea was born here: [Part 1](https://drukarnia.com.ua/articles/yak-chatgpt-vede-za-mene-kanal-v-telegram-i-u-nogo-ce-maizhe-vikhodit-chastina-1-VywRW) and [Part 2](https://drukarnia.com.ua/articles/yak-chatgpt-vede-za-mene-kanal-v-telegram-i-u-nogo-ce-maizhe-vikhodit-chastina-2-X9Yjz). (The articles are in Ukrainian, but I think you’ll manage to translate them into your preferred language.)

## How to run

### Requirements

- [docker](https://docs.docker.com/engine/install/) or/and [docker-compose](https://docs.docker.com/compose/install/)

### Clone repo

```shell
git clone https://github.com/Think-Root/chappie_bot.git
```

### Config

create a **.env** file in the app root directory

```properties
BOT_TOKEN=<bot token https://core.telegram.org/bots>
ADMIN_ID=<your telegram id>
CHANNEL_ID=<your channel id>
CHAPPIE_SERVER_URL=<e.g. http://localhost:8080/>
CHAPPIE_SERVER_BEARER=<chappie_server server token>
WAPP_JID=<whatsapp chat jid https://github.com/Think-Root/wapp>
WAPP_TOKEN=<whatsapp app token https://github.com/Think-Root/wapp>
```


> 🔴 IF YOU ARE NOT READY TO RISK LOSING YOUR WHATSAPP ACCOUNT, REMOVE THE SendMessageToWhatsApp FUNCTION CALL FROM THE CODE AND DO NOT SET UP THE WAPP_JID AND WAPP_TOKEN VARIABLES


### Deploy (If deploy before deploying the chappie_server)

- create network `docker network create chappie_network` (to allow the app access to the database run mariadb in this network)
- run [mariadb](https://hub.docker.com/_/mariadb) `docker run -d --name mariadb --network chappie_network -e MYSQL_ROOT_PASSWORD=your_password -p 3306:3306 mariadb:latest`
- build `docker build -t chappie_bot:latest -f Dockerfile .`
- run `docker run --name chappie_bot --restart always --env-file .env -e TZ=Europe/Kiev --network chappie_network chappie_bot:latest`
- or via docker compose `docker compose up -d`

### Deploy (If deploy after deploying the chappie_server)

- build `docker build -t chappie_bot:latest -f Dockerfile .`
- run `docker run --name chappie_bot --restart always --env-file .env -e TZ=Europe/Kiev --network chappie_network chappie_bot:latest`
- or via docker compose `docker compose up -d`

## Telegram bot commands

- **/add** - you can provide a list (or just one) of GitHub repository URLs separated by spaces, a short description will be generated for each, and they will be added to the database.
- **/next** - displays the next post scheduled for publication in the channel, if you provide a number after a space, that many posts will be shown
- **/gen** - generates a post for the channel from GitHub trends, if you specify a number after a space, that many posts will be generated and added to the database
- **/info** - shows information about the posts in the database

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