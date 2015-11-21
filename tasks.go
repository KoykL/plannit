package plannit
import(
  "time"
  "sort"
)
type Task struct{
  Title string
  Due time.Time

}
type Tasks []Task
func (t Tasks) Len() int {
    return len(t)
}
func (t Tasks) Less(i, j int) bool {
  if (t[i].Due.Equal(t[j].Due)){
    return t[i].Title < t[j].Title|| t[i].Title == t[j].Title
  }
    return t[i].Due.Before(t[j].Due)
}
func (t Tasks) Swap(i, j int) {
    t[i], t[j] = t[j], t[i]
}
func (t Task) EqualDue(j Task) bool {
    return t.Due.Equal(j.Due)
}
type TasksGroup struct{
  Tasks *Tasks
  Due time.Time
}
type TasksGroups []TasksGroup
func classifyTasks(ts *Tasks) TasksGroups{
  sort.Sort(ts)
  tds := make([]TasksGroup, 0)
  tasks := make([]Task, 1)
  tasks[0] = (*ts)[0]
  for i, t := range(*ts){
    if(i == 0){
      continue
    }
    if (!t.EqualDue((*ts)[i-1])){
      tasksObj := Tasks(tasks)
      td := TasksGroup{
        Tasks: &tasksObj,
        Due: tasks[0].Due,
      }
      tds = append(tds, td)
      tasks = make([]Task, 0)
    }
    tasks = append(tasks, t)
  }
  tasksObj := Tasks(tasks)
  td := TasksGroup{
    Tasks: &tasksObj,
    Due: tasks[0].Due,
  }
  tds = append(tds, td)
  return TasksGroups(tds)
}
func generateTasks(raw [][]string) (*Tasks, error){
  item_num := len(raw) - 1
  if (item_num < 0){
    item_num = 0

  }
  tasks := make([]Task, item_num)
  for i, item := range(raw){
    if(i == 0){
      continue
    }
    var dueDate time.Time
    if (len(item[7]) == 0){
      dueDate = time.Unix(0,0)
    } else {
      if dueDate_temp, err := time.Parse("2006-1-2 15:04:05 -0700", item[7]); err != nil{
        return nil, err
      }else{
        dueDate = dueDate_temp
      }
    }
    tasks[i-1] = Task{Title: item[2], Due: dueDate}
  }
  t := Tasks(tasks)
  sort.Sort(t)
  return &t, nil
}
