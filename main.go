package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/a-h/templ"
	"github.com/gorilla/feeds"
	"kot.pink/blog/pages"
)

//go:generate templ generate
func main() {
	posts := map[string]pages.PostMeta{
		"ssg.md": {
			Title:   "Building a static site generator using Go and Templ",
			Blurb:   "'cos who needs frameworks and bundlers!",
			Slug:    "ssg",
			Created: "2024-07-20",
			Updated: "2024-07-20",
		},
	}

	html := map[string]templ.Component{
		"index.html": pages.Index(posts),
		"rss.xml": templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			feedAuthor := &feeds.Author{Name: "Kot C", Email: "kot@yukata.dev"}

			feed := &feeds.Feed{
				Title:       "Kot's Blog",
				Link:        &feeds.Link{Href: "https://blog.kot.pink"},
				Description: "Heya! This is where I document my thoughts and the results of my adventures.",
				Author:      feedAuthor,
				Updated:     time.Now(),
			}

			for _, meta := range posts {
				createdTime, err := time.Parse("2006-01-02", meta.Created)
				if err != nil {
					log.Fatal(err)
				}

				updatedTime, err := time.Parse("2006-01-02", meta.Updated)
				if err != nil {
					log.Fatal(err)
				}

				feed.Items = append(feed.Items, &feeds.Item{
					Title:       meta.Title,
					Link:        &feeds.Link{Href: "https://blog.kot.pink/" + meta.Slug},
					Description: meta.Blurb,
					Author:      feedAuthor,
					Created:     createdTime,
					Updated:     updatedTime,
				})
			}

			rss, err := feed.ToRss()
			if err != nil {
				log.Fatal(err)
			}

			_, err = w.Write([]byte(rss))
			return err
		}),
	}

	for path, meta := range posts {
		data, err := os.ReadFile(filepath.Join("posts", path))
		if err != nil {
			log.Fatal(err)
		}

		os.Mkdir(filepath.Join("www", meta.Slug), 0755)
		html[filepath.Join(meta.Slug, "index.html")] = pages.Post(meta, data)
	}

	genFile, err := os.Create("_generated.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer genFile.Close()

	writer := io.MultiWriter(os.Stdout, genFile)
	for file, comp := range html {
		path := filepath.Join("www", file)
		fmt.Fprintln(writer, path)

		file, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		err = comp.Render(context.Background(), file)
		if err != nil {
			log.Fatal(err)
		}
	}
}
