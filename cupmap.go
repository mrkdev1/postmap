package main

import (
    "bytes"
    "fmt"
    "net/http"
)

func main() {
    url := "https://api.github.com/gists"

    var jsonStr = []byte(`{
                "description": "A (secret) gist",
                "public": false,
                "files": {
                        "file1.text": {
                            "content": "This is where map goes."
                        }
                    }
                }`)
    
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Authorization", "token 16dd785b42ead8ea7c73db50d0c2f47487b59003") // The token
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    if resp.Status == "201 Created" {
        fmt.Println("Success")
        fmt.Println("Go to the following address to access the secret gist")
        fmt.Println(resp.Header.Get("Location"))
    } else {
        fmt.Println("Failed creating secret gist")
    }
}