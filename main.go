package main
import(
"html/template"
"net/http"
)
type UserDetails struct {
	Name string
	Password string
	Contactno int
	City string
	Gender string
	Age int
	BloodGroup string
	Type string
}
func main() {
tmpl:=template.Must(template.ParseFiles("signup.html"))
http.HandleFunc("/",func(w http.ResponseWriter,r * http.Request)) {
	if r.Method !=http.MethodPost {
	tmpl.Execute(w,nill)
	return
	}
	detailss :=UserDetails {
	Name : r.FormValue("username")
	Password : r.FormValue("password")
	Contactno : r.FormValue("no")
	City : r.FormValue("city")
	Gender : r.FormValue("gender")
	Age : r.FormValue("age")
	BloodGroup : r.FormValue("Blood Group")
	Type : r.FormValue("type")
}
_=details
tmpl.Execute(w,struct{Success bool}{(true)}
})
http.ListenAndServe(":8080",nil)
}