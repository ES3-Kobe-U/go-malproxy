package service

// type RemoveFileService interface {
// 	RemoveFile(fqdn string) error
// }

// //作成したHTMLファイルを削除する。
// func RemoveFile(fqdn string) error {
// 	cmdRemove := exec.Command("rm", "test/"+fqdn+".html") //指定HTMLファイルの読み込み TODO: 後でディレクトリを変更
// 	cmdRemove.Stderr = os.Stderr
// 	cmdRemove.Stdin = os.Stdin
// 	out, err := cmdRemove.Output()
// 	if err != nil {
// 		log.Println("Err", err)
// 		return err
// 	} else {
// 		log.Println("remove files --> success:", string(out))
// 	}
// 	return nil
// }