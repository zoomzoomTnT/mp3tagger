package main

import (
	"fmt"
	"github.com/bogem/id3v2"
	"log"
)

func main() {
	tag, err := id3v2.Open("从别后 - 斗破苍穹OST - COPY.mp3", id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal("Error while opening mp3 file: ", err)
	}
	defer tag.Close()

	// Read tags
	//fmt.Println(tag.Artist())
	//fmt.Println(tag.Title())

	//remove comments
	//tag.DeleteFrames("TSSE")

	//comment := id3v2.CommentFrame{
	//	Encoding:    id3v2.EncodingUTF8,
	//	Language:    "chi",
	//	Description: "zoomzoom",
	//	Text:        "云韵主题曲",
	//}
	//tag.AddCommentFrame(comment)
	////
	//////add picture
	//coverPhoto, err := os.ReadFile("doupo3.jpeg")
	//if err != nil {
	//	log.Fatal("Error while reading coverPhoto file", err)
	//}
	//
	//pic := id3v2.PictureFrame{
	//	Encoding:    id3v2.EncodingUTF8,
	//	MimeType:    "image/jpeg",
	//	PictureType: id3v2.PTFrontCover,
	//	Description: "Front cover",
	//	Picture:     coverPhoto,
	//}
	//tag.AddAttachedPicture(pic)
	//
	//load txt file
	//txt, err := os.ReadFile("lyrics3.txt")
	//if err != nil {
	//	log.Fatal("Error while reading txt file", err)
	//}

	//sylt := id3v2.SynchronisedLyricsFrame{
	//	Encoding:          id3v2.EncodingUTF8,
	//	Language:          "chi",
	//	ContentDescriptor: "从别后",
	//	TimestampFormat:   id3v2.TimestampFormatMS,
	//	Lyrics:            string(txt),
	//}
	//tag.AddSynchronisedLyricsFrame(sylt)

	//uslt := id3v2.UnsynchronisedLyricsFrame{
	//	Encoding:          id3v2.EncodingUTF8,
	//	Language:          "chi",
	//	ContentDescriptor: "从别后",
	//	Lyrics:            string(txt),
	//}
	//tag.AddUnsynchronisedLyricsFrame(uslt)
	//// Write tag to file.mp3
	//if err = tag.Save(); err != nil {
	//	log.Fatal("Error while saving a tag: ", err)
	//}

	//print all frames
	//pictures := tag.GetFrames(tag.CommonID("Attached picture"))
	//for _, f := range pictures {
	//	pic, ok := f.(id3v2.PictureFrame)
	//	if !ok {
	//		log.Fatal("Couldn't assert picture frame")
	//	}
	//
	//	// save picture to file
	//	err = os.WriteFile("test.jpeg", pic.Picture, 0644)
	//	if err != nil {
	//		log.Fatal("Error while saving picture to file: ", err)
	//	}
	//
	//	// For example, print the description:
	//	fmt.Println(pic.Description)
	//}

	for id, frame := range tag.AllFrames() {
		if id == "APIC" {
			fmt.Println(id, fmt.Sprint(frame)[:100])
			continue
		}
		fmt.Println(id, fmt.Sprint(frame))
	}

}
