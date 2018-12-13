package main

import "awsS3Download/s3Downloader"

func main() {
	// S3の設定
	s3Info := s3Downloader.S3DownloadInfo{
		AccessKeyId:     "アクセスキー",
		SecretAccessKey: "シークレットキー",
		Region:          "ap-northeast-1", // 東京リージョン
		BucketName:      "バケット名",
	}

	// 保存するディレクトリのパスと取得するファイルの名前
	filePath := "./images/"
	fileName := "S3上のファイル名"

	// ファイルのダウンロード
	s3Downloader.FileDownloadFromS3(s3Info, filePath, fileName)
}
