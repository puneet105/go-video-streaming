package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var songDir = "content"
var port = "8080"

//ffmpeg command to make chunks of lare videos into 600 seconds segments
/*
 ffmpeg -i VIDEO_FILE.mp4 -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 600 -hls_list_size 0 -f hls index.m3u8
*/
func main(){
	router := mux.NewRouter()
	router.HandleFunc("/stream", StreamHandler)
	mediaFile := fmt.Sprintf("%s/index.m3u8",songDir)
	log.Println("media file is :",mediaFile )
	log.Println("Streaming server is listening on : ", port)
	http.ListenAndServe(":"+port, router)

}

func StreamHandler(w http.ResponseWriter, r *http.Request){
	mediaFile := fmt.Sprintf("%s/index.m3u8",songDir)
	log.Println("media file is :",mediaFile )
	http.ServeFile(w,r,mediaFile)
	w.Header().Set("Content-Type", "application/x-mpegURL")
}
