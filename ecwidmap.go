package main

import (
	"context"
	"flag"
	"fmt"
	"html/template"
	"os"
	"strconv"
	"time"

	"github.com/sevkin/ecwid"
)

const sitemap = `{{"<?"|html}}xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
{{$daily := .Daily}}{{range .Products}}<url><loc>{{.URL|urlquote}}</loc><lastmod>{{.UpdateTimestamp|lastmod}}</lastmod>{{if $daily}}<changefreq>daily</changefreq>{{end}}</url>
{{end}}{{range .Categories}}<url><loc>{{.URL|urlquote}}</loc>{{if $daily}}<changefreq>daily</changefreq>{{end}}</url>
{{end}}</urlset>`

func main() {
	const (
		storeIDEnv = "ECWID_STOREID"
		tokenEnv   = "ECWID_TOKEN"
	)

	var (
		storeID uint64
		token   string
		daily   bool
	)

	// сначала попробовать получить из переменных окружения
	storeID, _ = strconv.ParseUint(os.Getenv(storeIDEnv), 10, 64)
	token = os.Getenv(tokenEnv)

	// затем из командной строки
	flag.Uint64Var(&storeID, "storeid", storeID, fmt.Sprintf("store ID (can get from %s env)", storeIDEnv))
	flag.StringVar(&token, "token", token, fmt.Sprintf("token (can get from %s env)", tokenEnv))

	flag.BoolVar(&daily, "daily", daily, "force <changefreq>daily</changefreq>")

	flag.Parse()

	// обязательные значения storeID, token, filename
	if storeID == 0 || token == "" {
		fmt.Println("pass storeid and token via env or commandline")
		os.Exit(2)
	}

	store := ecwid.New(storeID, token)
	//store.SetDebug(true)

	template.Must(template.New("sitemap").
		Funcs(template.FuncMap{
			"html": func(s string) template.HTML { return template.HTML(s) },
			"urlquote": func(s string) string {
				return s
			},
			"lastmod": func(u uint64) string {
				return time.Unix(int64(u), 0).Format("2006-01-02")
			},
		}).
		Parse(sitemap)).
		Execute(os.Stdout, struct {
			Daily      bool
			Products   <-chan (*ecwid.Product)
			Categories <-chan (*ecwid.Category)
		}{
			Daily: daily,
			Products: store.Products(context.Background(), map[string]string{
				"enabled":   "true",
				"cleanUrls": "true",
			}),
			Categories: store.Categories(context.Background(), map[string]string{
				"cleanUrls": "true",
			}),
		})
}
