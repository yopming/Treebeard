package main

import (
    "os"
    "fmt"
    "strings"
    "strconv"
    "net/http"
    "path/filepath"
    //"encoding/json"
    "github.com/codegangsta/martini"
    "github.com/martini-contrib/render"
)

type Response struct {
    AbsPath     string
    FileItem    []struct {
        ItemInfo    []string
        ItemPath    string
    }
}

func main() {
    path_prefix := "Z:\\"

    m := martini.Classic()
    m.Use(render.Renderer())

    // Parse the URL
    m.Get("/files/**", func(
        w http.ResponseWriter,
        r *http.Request,
        render render.Render,
        params martini.Params
    ) {

        message := r.URL.RequestURI()
        prefix := "/files/"
        elite := strings.Split(message, prefix)[1]

        if len(elite) > 0 {
            // Step in the real folder
            real_path := filepath.ToSlash(path_prefix) + elite

            //itemInfo := map[string]map[string]string{}
            itemInfo := make(map[string]interface{})
            item     := make(map[string]interface{})

            filepath.Walk(real_path, func(path string, fileinfo os.FileInfo, err error) error {
                f, err := os.Stat(path)
                if err != nil {
                    panic(err)
                }

                if f.IsDir() {
                    item["type"] = "dir"
                    item["path"] = real_path
                } else {
                    item["type"] = "file"
                    item["path"] = real_path + "/" + f.Name()
                }
                item["time"] = f.ModTime().Format("2004-01-02 15:04:22") // time to string
                item["type"] = strconv.FormatInt(f.Size(), 10) // int64 to string

                itemInfo["name"] = item

                return nil
            })

            collection, err_collection := json.Marshal(itemInfo)
            if err_collection != nil {
                fmt.Println(err_collection)
            }
        }

        render.JSON(200, collection)
    });

    m.Run()
}
