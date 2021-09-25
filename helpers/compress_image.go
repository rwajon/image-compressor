package helpers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/h2non/bimg"
)

const DEFAULT_FORMAT string = "jpg"

func getOptions(o bimg.Options, opts map[string]interface{}) bimg.Options {
	for k, el := range opts {
		if k == "quality" && el != "" {
			if v, e := strconv.Atoi(fmt.Sprint(el)); e == nil {
				o.Quality = v
			}
		}
		if k == "width" && el != "" {
			if v, e := strconv.Atoi(fmt.Sprint(el)); e == nil {
				o.Width = v
			}
		}
		if k == "height" && el != "" {
			if v, e := strconv.Atoi(fmt.Sprint(el)); e == nil {
				o.Height = v
			}
		}
		if k == "crop" && el != "" {
			if v, e := strconv.ParseBool(fmt.Sprint(el)); e == nil {
				o.Crop = v
			}
		}
		if k == "format" && el != "" {
			imageType := strings.ToUpper(fmt.Sprint(el))
			if imageType == "WEBP" {
				o.Type = bimg.WEBP
			} else if imageType == "PNG" {
				o.Type = bimg.PNG
			} else if imageType == "JPG" || imageType == "JPEG" {
				o.Type = bimg.JPEG
			} else {
				o.Type = bimg.JPEG
			}
		}
	}
	return o
}

func CompressImage(buffer []byte, dirName string, opts map[string]interface{}) (string, error) {
	options := getOptions(bimg.Options{}, opts)
	ext := func() string {
		if opts["format"] != "" && options.Type != 0 {
			return fmt.Sprint(opts["format"])
		}
		return DEFAULT_FORMAT
	}()

	fileName := strings.Replace(uuid.New().String(), "-", "", -1) + "." + ext
	img, err := bimg.NewImage(buffer).Convert(options.Type)

	if err != nil {
		return fileName, err
	}
	if img, err = bimg.NewImage(img).Process(options); err != nil {
		return fileName, err
	}
	if err = bimg.Write(dirName+"/"+fileName, img); err != nil {
		return fileName, err
	}

	return fileName, nil
}
