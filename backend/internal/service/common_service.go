package service

import (
	"context"
	"fmt"
	"log"

	cld "github.com/Sairam-04/blog-app/backend/cloudinary"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type UploadImage struct {
	cloudinary *cloudinary.Cloudinary
}

// NewUploadImage returns a new UploadImage service instance
func NewUploadImage() *UploadImage {
	return &UploadImage{
		cloudinary: cld.Cld,
	}
}

// UploadService uploads an image file to Cloudinary and returns the image URL
func (h *UploadImage) UploadService(ctx context.Context, image_file string) (string, error) {
	// Upload the file to Cloudinary
	result, err := h.cloudinary.Upload.Upload(ctx, image_file, uploader.UploadParams{PublicID: "image", Folder: "files"})
	if err != nil {
		log.Println("Cloudinary upload error:", err)
		return "", fmt.Errorf("failed to upload image: %v", err)
	}

	// Return the image URL
	return result.SecureURL, nil
}
