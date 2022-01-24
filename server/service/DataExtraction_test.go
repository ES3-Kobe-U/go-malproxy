package service

// var DataExtraction_Sample = []string{
// 	"https://www.google.com/",
// 	"https://www.amazon.co.jp/-/en/",
// 	"https://www.google.com/search?q=Golang",
// }

// func TestDataExtraction(t *testing.T) {
// 	for _, Url := range DataExtraction_Sample {
// 		doc, err := goquery.NewDocument(Url)
// 		if err != nil {
// 			fmt.Print("url scarapping failed")
// 		}
// 		u, err := url.Parse(Url)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		res, err := doc.Html()
// 		if err != nil {
// 			fmt.Print("dom get failed")
// 		}
// 		fileName := u.Hostname()
// 		ioutil.WriteFile("test/"+fileName+".html", []byte(res), os.ModePerm)
// 	}
// }
