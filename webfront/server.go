package main


import (
    "fmt"
    "log"
    "io/ioutil"
    "bytes"
    "net/http"
    "encoding/json"
    "html/template"
)

type Sentiment struct {
    Message  string
    Score    string
    Weight   float64
}

type SentimentPageStruct struct {
    sentiments []Sentiment
}

var responses []Sentiment

func formHandler(w http.ResponseWriter, r *http.Request) {
    

    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }

    url := "http://34.68.129.129/text-generator"
    message := r.FormValue("message")
    if message == "" {
        tmpl := template.Must(template.ParseFiles("static/form.html"))

        tmpl.Execute(w, responses)
    }else {
        values := map[string]string{"text": message}
        jsonValue, _ := json.Marshal(values)
        req, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))

        if err != nil {
            log.Fatal( err )
        }

        defer req.Body.Close()

        body, _ := ioutil.ReadAll(req.Body)

        var data []map[string]interface{}

        err = json.Unmarshal(body, &data)
        if err != nil {
            log.Fatal( err )
        }

        if len(responses) > 20 {
            responses = []Sentiment{Sentiment{Message: message, Score: data[0]["label"].(string), Weight: data[0]["score"].(float64)}}
        }else{
            responses = append(responses, Sentiment{Message: message, Score: data[0]["label"].(string), Weight: data[0]["score"].(float64)})
        }

        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }

        tmpl := template.Must(template.ParseFiles("static/form.html"))

        tmpl.Execute(w, responses)
    }
}

func main() {
    // fileServer := http.FileServer(http.Dir("./static"))
    // http.Handle("/", fileServer)
    http.HandleFunc("/", formHandler)

    fmt.Printf("Starting server at port 8080\n")
    log.Fatal(http.ListenAndServe(":8080", nil))
}