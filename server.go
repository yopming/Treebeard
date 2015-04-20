package main

import (
	"github.com/go-humanize"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// Static File Folder
	path_prefix := "D:\\webroot\\e\\website\\demos\\ins\\设计稿"

	m := martini.Classic()
	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://fdc.vemic.com", "http://192.168.27.159:1080"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "x-requested-with", "Content-Type", "Content-Range", "Content-Disposition", "Content-Description"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))
	m.Use(render.Renderer())
	m.Use(martini.Static(filepath.ToSlash(path_prefix)))

	// Parse the URL
	m.Get("/files", func(w http.ResponseWriter, r *http.Request, render render.Render, params martini.Params) {
		real_path := path.Join(filepath.ToSlash(path_prefix))
		var elite string
		itemInfo := ReadGraphic(real_path, elite)
		render.JSON(200, itemInfo)
	})

	m.Get("/files/**", func(w http.ResponseWriter, r *http.Request, render render.Render, params martini.Params) {
		message := r.URL.RequestURI()
		prefix := "/files/"
		splitter := strings.Split(message, prefix)
		elite := DecodeURI(splitter[1])

		if len(elite) >= 0 {
			real_path := path.Join(filepath.ToSlash(path_prefix), elite)
			ReadGraphic(real_path, elite)
			itemInfo := ReadGraphic(real_path, elite)
			render.JSON(200, itemInfo)
		}

	})

	m.RunOnAddr(":54321")
}

/**
 * Read directory and return json
 */
func ReadGraphic(real_path string, elite string) interface{} {
	// Step in the real folder
	itemInfo := make(map[string]interface{})
	files, _ := ioutil.ReadDir(real_path)
	i := 0

	for _, f := range files {
		// non-hidden files(starting with .)
		if !strings.HasPrefix(f.Name(), ".") {
			item := make(map[string]interface{})

			if f.IsDir() {
				item["type"] = "Directory"
				item["down"] = elite + "/" + f.Name()
			} else {
				item["type"] = "File"
				item["down"] = "http://demo.vemic.com/demos/ins/设计稿/" + elite + "/" + f.Name()
			}
			if !strings.HasPrefix(f.Name(), ".") {
				item["name"] = f.Name()
			}
			item["time"] = f.ModTime().Format("2006/01/02 15:04")
			item["human_time"] = humanize.Time(f.ModTime())
			item["size"] = humanize.Bytes(uint64(f.Size()))
			itemInfo[strconv.Itoa(i)] = item
		}
		i++
	}

	return itemInfo
}

/**
 * Input: filepath
 * Output: encode string
 */
func EncodeURI(filepath string) string {
	if len(filepath) < 1 {
		return ""
	}

	u := url.URL{}
	u.Path = filepath
	return u.String()
}

/**
 * Input: URI string
 * Output: filepath
 */
func DecodeURI(uri string) string {
	if len(uri) < 1 {
		return ""
	}

	u, err := url.Parse(uri)
	if err != nil {
		return ""
	}
	return u.Path
}
