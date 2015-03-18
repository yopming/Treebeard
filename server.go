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

    // Parse the URL
    m.Get("/files/**", func(w http.ResponseWriter, r *http.Request) {
        message := r.URL.RequestURI()
        prefix := "/files/"
        elite := strings.Split(message, prefix)[1]
        //collection := make(map[string]Json)

        if len(elite) > 0 {
            // Step in the real folder
            real_path := filepath.ToSlash(path_prefix) + elite

            itemInfo := map[string]map[string]string{}

            filepath.Walk(real_path, func(path string, fileinfo os.FileInfo, err error) error {
                f, err := os.Stat(path)
                if err != nil {
                    panic(err)
                }

                itemInfo["name"]["name"] = f.Name()
                itemInfo["name"]["time"] = f.ModTime().Format("01-02 15:04:05") // time.time to string
                itemInfo["name"]["type"] = strconv.FormatInt(f.Size(), 10) // int64 to string

                //for key, value := range itemInfo {
                    //fmt.Printf("%s: %s \n", key, value)
                //}

                return nil
            })

        }
    });

    m.Run()
}
