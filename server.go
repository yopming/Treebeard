package main

import (
    //"fmt"
    "time"
    "strings"
    "strconv"
    "net/http"
    "io/ioutil"
    "path/filepath"
    //"encoding/json"
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
)


func main() {
    // Static File Folder
    path_prefix := "Z:\\"

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(martini.Static(filepath.ToSlash(path_prefix)))

    // Parse the URL
    m.Get("/files/**", func(w http.ResponseWriter, r *http.Request, render render.Render, params martini.Params) {

        message := r.URL.RequestURI()
        prefix := "/files/"
        elite := strings.Split(message, prefix)[1]

        if len(elite) > 0 {
            // Step in the real folder
            real_path := filepath.ToSlash(path_prefix) + elite

            itemInfo := make(map[string]interface{})

            files, _ := ioutil.ReadDir(real_path)
            i := 0

            for _, f := range files {
                // non-hidden files(starting with .)
                if !strings.HasPrefix(f.Name(), ".") {
                    item := make(map[string]interface{})

                    if f.IsDir() {
                        item["type"] = "Directory"
                    } else {
                        item["type"] = "File"
                    }
                    if !strings.HasPrefix(f.Name(), ".") {
                        item["name"] = f.Name()
                    }
                    item["time"] = f.ModTime().Format(time.RFC3339)
                    item["size"] = strconv.FormatInt(f.Size(), 10)
                    itemInfo[strconv.Itoa(i)] = item
                }
                i++
            }

            render.JSON(200, itemInfo)
        }

    });

    m.Run()
}
