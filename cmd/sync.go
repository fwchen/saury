package main

import (
	"github.com/fwchen/saury/model"
	"github.com/fwchen/saury/repository"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"strconv"
)

const DefaultGalleries = "./galleries"

func main() {
	args:= os.Args
	var galleriesDir string
	if len(args) < 2 {
		galleriesDir = DefaultGalleries
	} else {
		galleriesDir = args[1]
	}
	entries, err := ioutil.ReadDir(galleriesDir)
	if err != nil {
		log.Fatal(err)
	}
	appDb, err := repository.Connect("mongodb://localhost:27017", "saury")
	if err != nil {
		log.Fatal(err)
	}
	galleryRepo := repository.NewGalleryRepository(appDb)

	//var albums []model.Album
	for _, f := range entries {
		if !f.IsDir() {
			continue
		}
		var photos []string
		album := model.Album{
			Name:   "",
			Uri:    "",
		}
		files, err := ioutil.ReadDir(path.Join(galleriesDir, f.Name()))
		if err != nil {
			log.Fatal(err)
		}
		for _, p := range files {
			if p.IsDir() {
				continue
			}
			photos = append(photos, p.Name())
		}
		sort.Strings(photos)
		album.Photos = photos
		galleryRepo.Save(&album)
		//albums = append(albums, album)
	}
}