package models

var ArticleData = Articles{
	{
		Title: "Example 1",
		Body: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum " +
			"has been the industry's standard dummy text ever since the 1500s, when an unknown " +
			"printer took a galley of type and scrambled it to make a type specimen book.",
	},
}

type Article struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Articles []Article
