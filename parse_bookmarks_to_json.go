package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"os"
	"flag"
	//"encoding/base64"
	"encoding/json"
	"strconv"
	"github.com/go-redis/redis"
)

type BookMark struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	AddDate     uint64 `json:"add_date ome"`
}

type BookMarkFolder struct {
	Name            string            `json:"name"`
	AddDate         uint64            `json:"add_date"`
	LastMofified    uint64            `json:"last_modified"`
	PersonalToolBar string            `json:"personal_tool_bar"`
	Level uint32			`json:"level"`
	BookMarks       []*BookMark       `json:"bookmarks"`
	ChildrenFolders []*BookMarkFolder `json:"children"`
}

func walk_folder(root *goquery.Selection, book_mark_folder *BookMarkFolder, max_depth int) {

	root.ChildrenFiltered("dt").Each(func(i int, selection *goquery.Selection) {
		h3 := selection.ChildrenFiltered("h3")
		//fmt.Println(h3.Nodes == nil)
		if h3.Nodes != nil {
			if int(book_mark_folder.Level) >= max_depth{
				selection.Find("a").Each(func(i int, selection *goquery.Selection) {

					url, _ := selection.Attr("href")
					add_date, _ := selection.Attr("add_date")
					//icon,_ := href.Attr("icon")
					title := selection.Text()
					//fmt.Println(title)
					book_mark := new(BookMark)
					book_mark.Url = url
					//book_mark.Icon = icon
					book_mark.Title = title
					book_mark.AddDate, _ = strconv.ParseUint(add_date, 10, 32)
					book_mark_folder.BookMarks = append(book_mark_folder.BookMarks, book_mark)
				})
			}else{
				new_folder := new(BookMarkFolder)
				add_date, _ := h3.Attr("add_date")
				last_modified, _ := h3.Attr("last_modified")
				personal_tool_bar, _ := h3.Attr("personal_tool_bar")
				name := h3.Text()
				//fmt.Println(name)
				new_folder.Name = name
				new_folder.AddDate, _ = strconv.ParseUint(add_date, 10, 32)
				new_folder.LastMofified, _ = strconv.ParseUint(last_modified, 10, 32)
				new_folder.PersonalToolBar = personal_tool_bar
				new_folder.Level = book_mark_folder.Level + 1
				book_mark_folder.ChildrenFolders = append(book_mark_folder.ChildrenFolders, new_folder)
				dl := selection.ChildrenFiltered("dl").First()
				walk_folder(dl, new_folder, max_depth)

			}
		} else {
			href := selection.ChildrenFiltered("a").First()
			url, _ := href.Attr("href")
			add_date, _ := href.Attr("add_date")
			//icon,_ := href.Attr("icon")
			title := href.Text()
			//fmt.Println(title)
			book_mark := new(BookMark)
			book_mark.Url = url
			//book_mark.Icon = icon
			book_mark.Title = title
			book_mark.AddDate, _ = strconv.ParseUint(add_date, 10, 32)
			book_mark_folder.BookMarks = append(book_mark_folder.BookMarks, book_mark)

		}
	})

}

func ParseBookmarks(file io.Reader, max_depth int) (*BookMarkFolder, error) {
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		fmt.Println("haha",err)
		return nil, err
	}
	book_mark_folder := new(BookMarkFolder)
	book_mark_folder.Level = 0
	root := doc.Find("DL").First()
	root_title := root.Find("h3").First().Text()
	//fmt.Println(root_title)
	book_mark_folder.Name = root_title
	dl := root.Find("dl").First()
	//fmt.Println(dl.Text())
	//doc.Children().Each(func(i int, selection *goquery.Selection){
	//	//fmt.Println(selection.Nodes)
	//	fmt.Println(i)
	//	fmt.Println(selection.Text())
	//	//fmt.Println(selection.ChildrenFiltered("a").Text())
	//})
	walk_folder(dl, book_mark_folder, max_depth)
	return book_mark_folder, nil
}

func NewClient() *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:13379",
		Password: "",
		DB: 0,
	})
	//pong, err := client.Ping().Result()
	//fmt.Println(pong, err)
	return client
}

func main() {
	var (
		path string
		redis_key string
		user_id string
		max_depth int
	)
	flag.StringVar(&path, "path", "", "a string")
	flag.StringVar(&redis_key, "redis_key", "", "a string")
	flag.StringVar(&user_id, "user_id", "", "a string")
	flag.IntVar(&max_depth, "max_depth", 1, "an int")
	flag.Parse()

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("sdf", err)
	}
	defer file.Close()
	book_mark_folder, _ := ParseBookmarks(file, max_depth)
	b, err := json.MarshalIndent(book_mark_folder, "", "    ")
	if err != nil {
		fmt.Println("error", err)
	}
	redis_client := NewClient()
	redis_client.HSet(redis_key, user_id, b)

}
