package route

import (
	"dropboxshare/middleware"
	"dropboxshare/util"
	"fmt"
	"net/http"
	"regexp"
)

// 路由定义
type routeInfo struct {
	Reg     *regexp.Regexp
	Handler func(http.ResponseWriter, *http.Request, []string)
}

// 路由添加
var RoutePath = []routeInfo{
	{regexp.MustCompile(`^/files/(.*)$`), files},
	{regexp.MustCompile(`^/imgs/(.*)$`), imgs},
	{regexp.MustCompile(`^/(large|medium|small)/([\w\-]{6,12})\.(mp4|flv|webm|3gp|json)$`), youtube_video},
	{regexp.MustCompile(`^/(large|medium|small)/([\w\-]{6,12})\.(jpg|webp)$`), youtube_image},
}

func files(w http.ResponseWriter, r *http.Request, match []string) {
	err := middleware.ServeBoxFile(w, r, match)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), 500)
	}
}

func imgs(w http.ResponseWriter, r *http.Request, match []string) {
	fmt.Println(match)
}

func youtube_video(w http.ResponseWriter, r *http.Request, match []string) {
	url, jsonData, err := middleware.GetYoutubeVideoUrl(match)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), 500)
	} else if url != "" {
		middleware.ServeYoutubeVideo(w, r, url)
	} else {
		util.JsonPut(w, jsonData, true, 600)
	}

}

func youtube_image(w http.ResponseWriter, r *http.Request, match []string) {
	var url string = middleware.GetYoutubeImageUrl(match)
	middleware.ServeYoutubeImage(w, r, url)
}
