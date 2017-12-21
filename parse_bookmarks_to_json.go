package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"os"
	//"encoding/base64"
	"encoding/json"
	"strconv"
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
	BookMarks       []*BookMark       `json:"bookmarks"`
	ChildrenFolders []*BookMarkFolder `json:"children"`
}

func walk_folder(root *goquery.Selection, book_mark_folder *BookMarkFolder) {

	root.ChildrenFiltered("dt").Each(func(i int, selection *goquery.Selection) {
		h3 := selection.ChildrenFiltered("h3")
		//fmt.Println(h3.Nodes == nil)
		if h3.Nodes != nil {
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
			book_mark_folder.ChildrenFolders = append(book_mark_folder.ChildrenFolders, new_folder)
			dl := selection.ChildrenFiltered("dl").First()
			walk_folder(dl, new_folder)
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

func ParseBookmarks(file io.Reader) (*BookMarkFolder, error) {
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		fmt.Println("haha",err)
		return nil, err
	}
	book_mark_folder := new(BookMarkFolder)
	root_title := doc.Find("H1").Text()
	book_mark_folder.Name = root_title
	dl := doc.Find("DL").First()
	//doc.Children().Each(func(i int, selection *goquery.Selection){
	//	//fmt.Println(selection.Nodes)
	//	fmt.Println(i)
	//	fmt.Println(selection.Text())
	//	//fmt.Println(selection.ChildrenFiltered("a").Text())
	//})
	walk_folder(dl, book_mark_folder)
	return book_mark_folder, nil
}

func main() {
	file, err := os.Open("/home/shaoxing/Documents/bookmarks_12_20_17.html")
	if err != nil {
		fmt.Println("sdf", err)
	}
	defer file.Close()
	book_mark_folder, _ := ParseBookmarks(file)
	b, err := json.MarshalIndent(book_mark_folder, "", "    ")
	if err != nil {
		fmt.Println("error", err)
	}
	//fmt.Printf("%T", book_mark_folder)
	os.Stdout.Write(b)
	//fmt.Printf("%T",b)
}
