package cloudinary

import (
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
)

var Cld *cloudinary.Cloudinary

func Init(url string) {
	cld, err := cloudinary.NewFromURL(url)
	if err != nil {
		log.Fatal("Failed to initialize Cloudinary", err)
	}
	Cld = cld
}
