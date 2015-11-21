package plannit
import(
  "net/http"
  "html/template"
  "encoding/csv"
)

func init(){
  http.HandleFunc("/", handler)
  http.HandleFunc("/api/schedule", scheduleHandler)
  http.Handle("/static/", http.StripPrefix("/static/" ,http.FileServer(http.Dir("static"))))
}
func handler(w http.ResponseWriter, r *http.Request){
  /*dir := http.Dir(".")
  if page, err := dir.Open("templates/main.html"); err != nil {
    http.Error(w, err.Error(), 500)
  } else {
    defer page.Close()
    if buffer, err := ioutil.ReadAll(page); err != nil {
      http.Error(w, err.Error(), 500)
    }else{
      if _, err := w.Write(buffer); err != nil {
        http.Error(w, err.Error(), 500)
      }
    }
  }*/
  if file, _, err := r.FormFile("schedule"); err != nil{
    if t, err := template.ParseFiles("templates/main.html"); err != nil {
      http.Error(w, err.Error(), 500)
    } else {
      t.Execute(w, nil)
    }
  }else {
    defer file.Close()
    csvReader := csv.NewReader(file)
    if result, err := csvReader.ReadAll(); err != nil{
      http.Error(w, err.Error(), 500)
    }else{
      if t, err := template.ParseFiles("templates/main.html"); err != nil {
        http.Error(w, err.Error(), 500)
      } else {
        if tasks, err:= generateTasks(result); err != nil{
          http.Error(w, err.Error(), 500)
        }else{
          taskGroups := classifyTasks(tasks)
          if err := t.Execute(w, taskGroups); err!= nil{
            http.Error(w, err.Error(), 500)
          }
        }

      }

    }

  }
}

func scheduleHandler(w http.ResponseWriter, r *http.Request){

}
