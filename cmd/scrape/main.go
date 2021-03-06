package main

import (
	"flag"
	"log"
	"os"

	"github.com/anacrolix/torrent/tracker"
	_ "github.com/anacrolix/torrent/tracker/udp"
	"github.com/anacrolix/torrent/util"
	metainfo "github.com/anacrolix/libtorgo/metainfo"
)

func main() {
	flag.Parse()
	mi, err := metainfo.Load(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	for _, tier := range mi.AnnounceList {
		for _, url := range tier {
			tr, err := tracker.New(url)
			if err != nil {
				log.Fatal(err)
			}
			err = tr.Connect()
			if err != nil {
				log.Fatal(err)
			}
			ar := tracker.AnnounceRequest{
				NumWant: -1,
			}
			util.CopyExact(ar.InfoHash, mi.Info.Hash)
			resp, err := tr.Announce(&ar)
			if err != nil {
				log.Fatal(err)
			}
			log.Print(resp)
		}
	}
}
