package viewmodel


type Article struct{
	Title string
	Active string
	Content string
	img     string
	
}


func NewPost() Article {
	result := Article{
		Title: "post",
		Active: "post",
	}
	return result
}


func ViewPost() Article {
	result := Article{
		Title: "view",
		Active: "post",
	}
	return result
}



