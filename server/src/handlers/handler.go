package handlers

import (
  "fmt"
  "net/http"
  "encoding/json"
  "sync"
  "strings"
  "io/ioutil"
  "encoding/base64"
)

func LinksWithWord(w http.ResponseWriter, r *http.Request) {
  word := r.URL.Query()["word"][0];
  links := r.URL.Query()["links"][0];

  if len(word) == 0 || len(links) == 0 {
    response, _ := json.Marshal("Not enough data provided");
    fmt.Fprintf(w, string(response));
    return;
  }
  wordDecoded, _ := base64.StdEncoding.DecodeString(word);
  word = string(wordDecoded);

  linksSlice := strings.Split(links, ",");
  linksLength := len(linksSlice);

  var wg sync.WaitGroup;
  wg.Add(linksLength);

  linksWithWord := make([]string, linksLength);
  for i := 0; i < linksLength; i++  {
    go func(wg *sync.WaitGroup, linksWithWord *[]string, i int, linksSlice []string, word string) {
       defer wg.Done();
       resp, _ := http.Get(linksSlice[i]);
       defer resp.Body.Close();
       content, _ := ioutil.ReadAll(resp.Body);
       if strings.Contains(string(content), word) {
         (*linksWithWord)[i] = linksSlice[i];
       }
    }(&wg, &linksWithWord, i, linksSlice, word);
  }
  wg.Wait();
  response, _ := json.Marshal(linksWithWord);
  fmt.Fprintf(w, string(response));
}
