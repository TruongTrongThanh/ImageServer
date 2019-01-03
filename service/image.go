package service

// Image struct for image instance
type Image struct {
	ID           string
	Name         string
	Ext          string
	HasThumbnail bool
}

// GetAllImages get image list
func GetAllImages() []Image {
	return []Image{
		Image{"111", "cool", "jpg", false},
		Image{"222", "haha", "gif", false},
		Image{"333", "yes", "png", true},
	}
}

// GetImageByID get image by using ID
func GetImageByID(id string) Image {
	return Image{"12356", "img", "jpg", false}
}
