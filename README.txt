install all dependencies
```bash
run go mod tidy

cd into main.go, build and run
```bash
cd cmd/main
go build
go run main.go

This backend was created using a mySQL database. I the code I used to create the database table can be found in mydb.sql.


this repo's packages:
"github.com/guivari/sharingvision-backend/pkg/"

using gorm to interact with mySQL
"github.com/jinzhu/gorm/"

sql driver
"github.com/jinzhu/gorm/dialects/mysql"

using mux for routes
"github.com/gorilla/mux"

cmd
L main
  L main.go

pkg
L config
  L app.go
L controllers
  L post-controller.go
L models 
  L post.go
L routes 
  posts-routes.go
L utils
  Lutils.go



