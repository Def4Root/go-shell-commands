# mv

## v1
`The process cannot access the file because it is being used by another process.`という他のプロセスがファイルを開いてますというエラーが解決できない

## v2
悔しいけどから`os.Rename`に変更。osパッケージは偉大。