# kawethradb <br>
Open Source CSV Database Module For Golang

### WHY KawethraDB?!?!?
Because KawethraDB does not distinguish datatype. 
which means that a garbage collector programming language like go will be more integrated and work more seamlessly.
<br>
<table>
<tr>
<td><a href="#create">CreateDB</a></td>
<td><a href="#insert">Insert</a></td>
<td><a href="#update">Update</a></td>
<td><a href="#delete">Delete</a></td>
<td><a href="#find">Find</a></td>
<td><a href="#count">Count</a></td>
<td><a href="#list">List</a></td>
</tr>
</table>
<br><br><br>

## Commands

<div id="create">

### Create DB
```go
package main

import (
	"fmt"
	kawethradb "github.com/Hasan-Kilici/kawethradb"
)

type Ogrenci struct {
	ID    int
	Ad    string
	Soyad string
	Sinif int
}

func main(){
	ogrenciler := []Ogrenci{
		{ID: 1, Ad: "Ali", Soyad: "Veli", Sinif: 9},
		{ID: 2, Ad: "Ahmet", Soyad: "Mehmet", Sinif: 10},
		{ID: 3, Ad: "Ayşe", Soyad: "Fatma", Sinif: 11},
		{ID: 4, Ad: "Hasan", Soyad: "KILICI", Sinif: 12},
	}

	err := kawethradb.CreateDB("Ogrenciler", "./data/Ogrenciler.csv", ogrenciler)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}
```
</div>
<div id="insert">

### Insert Single
```go
package main

import (
	"fmt"
	kawethradb "github.com/Hasan-Kilici/kawethradb"
)

type Ogrenci struct {
	ID    int
	Ad    string
	Soyad string
	Sinif int
}

func main() {
	ogrenci := Ogrenci{ID: 1, Ad: "Ali", Soyad: "Veli", Sinif: 9}
	kawethradb.Insert("./data/Ogrenciler.csv", ogrenci)
        fmt.Println("Inserted!")
}
```
### Insert Multiple
```go
package main

import (
	"fmt"
	kawethradb "github.com/Hasan-Kilici/kawethradb"
)

type Ogrenci struct {
	ID    int
	Ad    string
	Soyad string
	Sinif int
}

func main() {
	ogrenci := []Ogrenci{
		{ID: 1, Ad: "Ali", Soyad: "Veli", Sinif: 9},
		{ID: 2, Ad: "Ahmet", Soyad: "Mehmet", Sinif: 10},
		{ID: 3, Ad: "Ayşe", Soyad: "Fatma", Sinif: 11},
		{ID: 4, Ad: "Hasan", Soyad: "KILICI", Sinif: 12},
	}

	kawethradb.Insert("./data/Ogrenciler.csv", ogrenci)
        fmt.Println("Inserted!")
}
```
</div>
<div id="find">

### Find
```go
package main

import (
	"fmt"
	kawethradb "github.com/Hasan-Kilici/kawethradb"
)

func main(){
	find, _ := kawethradb.Find("./data/Ogrenciler.csv", "ID", 3)
	fmt.Println(find)
}
```
### Find By ID
```go
package main

import (
	"fmt"
	kawethradb "github.com/Hasan-Kilici/kawethradb"
)

func main(){
	find, _ := kawethradb.FindByID("./data/Ogrenciler.csv", 3)
	fmt.Println(find)
}
```
### Find All (TODO LIST EXAMPLE <a href="https://github.com/Hasan-Kilici/kawethradb-todo-list">Go To link</a>)
```go
func main() {
	r := gin.Default()

	r.LoadHTMLGlob("src/*.tmpl")
	r.Static("/static", "./static/")

	r.GET("/", func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("ID")
		if err != nil {
			id := kawethradb.Count("./Tasks.csv")
			ctx.SetCookie("ID", strconv.Itoa(id), 36000, "/", "", false, true)
			ctx.Redirect(http.StatusFound, "/")
			return
		}

		userID, _ := strconv.Atoi(cookie)
		results, _ := kawethradb.FindAll("./Tasks.csv", "UserID", userID)
		var tasks []Task
		for _, result := range results {
			taskid, _ := strconv.Atoi(result["ID"])
			task := Task{
				ID:         taskid,
				UserID:     userID,
				Tasks:      result["Tasks"],
				Taskstatus: result["Taskstatus"],
			}
			tasks = append(tasks, task)
		}

		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Tasks":  tasks,
			"UserID": userID,
		})
	})
```
</div>
<div id="delete">

### Delete
```go
package main

import (
	kawethradb "github.com/Hasan-Kilici/kawethradb"
)

func main(){
  kawethradb.Delete("./data/Ogrenciler.csv", "ID", 2)
}
```
	
### Delete By ID
```go
package main

import (
	kawethradb "github.com/Hasan-Kilici/kawethradb"
)

func main(){
  kawethradb.DeleteByID("./data/Ogrenciler.csv", 2)
}
```
</div>
<div id="update">

### Update
```go
package main

import (
	"fmt"
	kawethradb "github.com/Hasan-Kilici/kawethradb"
)

func main(){
yeniVeri := []string{"2", "Hasan", "Kılıcı", "12"}
	err := kawethradb.Update("./data/Ogrenciler.csv", "ID", 2, yeniVeri)
	if err != nil {
		fmt.Println("Kayıt güncellenirken bir hata oluştu:", err)
		return
	}

	fmt.Println("Kayıt başarıyla güncellendi.")
}
```
### Update By ID
```go
package main

import (
	"fmt"
	kawethradb "github.com/Hasan-Kilici/kawethradb"
)

func main(){
yeniVeri := []string{"2", "Hasan", "Kılıcı", "12"}
	err := kawethradb.UpdateByID("./data/Ogrenciler.csv", 1, yeniVeri)
	if err != nil {
		fmt.Println("Kayıt güncellenirken bir hata oluştu:", err)
		return
	}

	fmt.Println("Kayıt başarıyla güncellendi.")
}
```
</div>
<div id="count">

### Get Count
```go
package main

import (
  "fmt"
  kawethradb "github.com/Hasan-Kilici/kawethradb"
)

type Ogrenci struct {
	ID    int
	Ad    string
	Soyad string
	Sinif int
}

func main(){
	ogrenciler := []Ogrenci{
		{ID: 1, Ad: "Ali", Soyad: "Veli", Sinif: 9},
		{ID: 2, Ad: "Ahmet", Soyad: "Mehmet", Sinif: 10},
		{ID: 3, Ad: "Ayşe", Soyad: "Fatma", Sinif: 11},
		{ID: 4, Ad: "Hasan", Soyad: "KILICI", Sinif: 12},
	}

	err := kawethradb.CreateDB("Ogrenciler", "./Ogrenciler.csv", ogrenciler)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

  count := kawethradb.Count("./Ogrenciler.csv")
  fmt.Println(count)
}

```
</div>
<div id="list">

### List

```go
func main() {
  ogrenci := []Ogrenci{
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
    {ID:1,Ad:"Hasan",Soyad:"KILICI"},
  }

  fmt.Println(kawethradb.List(ogrenci,10,30))
}
```

### List in HTTP

```go
r := gin.Default()

	r.LoadHTMLGlob("src/*.tmpl")
	r.Static("/static", "./static/")

	r.GET("/", func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("ID")
		if err != nil {
			id := kawethradb.Count("./Tasks.csv")
			ctx.SetCookie("ID", strconv.Itoa(id), 36000, "/", "", false, true)
			ctx.Redirect(http.StatusFound, "/")
			return
		}

		userID, _ := strconv.Atoi(cookie)
		results, _ := kawethradb.FindAll("./Tasks.csv", "UserID", userID)
		var tasks []Task
		for _, result := range results {
			taskid, _ := strconv.Atoi(result["ID"])
			task := Task{
				ID:         taskid,
				UserID:     userID,
				Tasks:      result["Tasks"],
				Taskstatus: result["Taskstatus"],
			}
			tasks = append(tasks, task)
            tasks = kawethradb.List(tasks,0,30)
		}

		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Tasks":  tasks,
			"UserID": userID,
		})
	})
```
</div>
