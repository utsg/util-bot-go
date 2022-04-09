package util

import (
		"log"
	"os"

	"github.com/hekmon/transmissionrpc"
)

func initClient() transmissionrpc.Client {
	torrentUrl := os.Getenv("TORRENT_CLIENT_URL")
	torrentUser := os.Getenv("TORRENT_CLIENT_USER")
	torrentPassword := os.Getenv("TORRENT_CLIENT_PASSWORD")

	transmissionbt, err := transmissionrpc.New(torrentUrl, torrentUser, torrentPassword, nil)
	if err != nil {
		log.Fatalln(err)
	}

	return *transmissionbt
}

func AddTorrent(filePath string) {
	client := initClient()
	torrent, err := client.TorrentAddFile(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(*torrent.ID)
    log.Println(*torrent.Name)
    log.Println(*torrent.HashString)
}
