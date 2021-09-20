package helpers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/h2non/bimg"
)

const DEFAULT_IMG_FORMAT string = "jpg"

func getOptions(o bimg.Options, opts map[string]interface{}) bimg.Options {
	for key, element := range opts {
		if key == "quality" && element != "" {
			if v, e := strconv.Atoi(fmt.Sprint(element)); e == nil {
				o.Quality = v
			}
		}
		if key == "width" && element != "" {
			if v, e := strconv.Atoi(fmt.Sprint(element)); e == nil {
				o.Width = v
			}
		}
		if key == "height" && element != "" {
			if v, e := strconv.Atoi(fmt.Sprint(element)); e == nil {
				o.Height = v
			}
		}
		if key == "crop" && element != "" {
			if v, e := strconv.ParseBool(fmt.Sprint(element)); e == nil {
				o.Crop = v
			}
		}
		if key == "format" && element != "" {
			img_type := strings.ToUpper(fmt.Sprint(element))
			if img_type == "WEBP" {
				o.Type = bimg.WEBP
			} else if img_type == "PNG" {
				o.Type = bimg.PNG
			} else if img_type == "JPG" || img_type == "JPEG" {
				o.Type = bimg.JPEG
			} else {
				o.Type = bimg.JPEG
			}
		}
	}
	return o
}

func CompressImage(buffer []byte, dirname string, opts map[string]interface{}) (string, error) {
	options := getOptions(bimg.Options{}, opts)
	img_extension := func() string {
		if opts["format"] != "" && options.Type != 0 {
			return fmt.Sprint(opts["format"])
		}
		return DEFAULT_IMG_FORMAT
	}()

	filename := strings.Replace(uuid.New().String(), "-", "", -1) + "." + img_extension
	img, err := bimg.NewImage(buffer).Convert(options.Type)

	if err != nil {
		return filename, err
	}
	if img, err = bimg.NewImage(img).Process(options); err != nil {
		return filename, err
	}
	if err = bimg.Write(dirname+"/"+filename, img); err != nil {
		return filename, err
	}

	return filename, nil
}
