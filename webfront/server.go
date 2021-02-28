package main


import (
    "fmt"
    "log"
    "strconv"
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
    Sentiments []Sentiment
}

var responses SentimentPageStruct

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
        return
    }

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
    
    responses.Sentiments = append(responses.Sentiments, Sentiment{Message: message, Score: data[0]["label"].(string), Weight: data[0]["score"].(float64)})
    
    if len(responses.Sentiments) > 5 {
        responses.Sentiments = responses.Sentiments[1:]
    }

    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }

    tmpl := template.Must(template.ParseFiles("static/form.html"))

    tmpl.Execute(w, responses)
}

func deletionHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()

    i, err := strconv.Atoi(r.Form["delete"][0])
    if err != nil {
        log.Fatal( err )
    }

    if (i > len(responses.Sentiments)) || (i < 0){
        tmpl := template.Must(template.ParseFiles("static/form.html"))

        tmpl.Execute(w, responses)
        return
    }

    responses.Sentiments = append(responses.Sentiments[:i], responses.Sentiments[i+1:]...)

    tmpl := template.Must(template.ParseFiles("static/form.html"))

    tmpl.Execute(w, responses)
}

func main() {
    http.HandleFunc("/", formHandler)
    http.HandleFunc("/delete", deletionHandler)
    fmt.Printf("Starting server at port 8080\n")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
