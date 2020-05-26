package model
import (
	"time"
	"log"
)

type Article struct{
	id int
	title string
	content string
	publishedTime *time.Time

}

func PostArticle(title,content string,userId int)(*Article,error)  {
	result := &Article{}
	db := dbConn()
	publishedTime := time.Now()
	insForm, err := db.Prepare("INSERT INTO Article(title, content,published_at,author) VALUES(?,?,?,?)")
	if err != nil {
			panic(err.Error())
	}
	
	insForm.Exec(title,content,publishedTime,3)
	log.Println(insForm)
	log.Println(err)

	log.Println("INSERT: article: " + title + " | time: " + publishedTime.String())

	defer db.Close()
	return result,nil

	
}