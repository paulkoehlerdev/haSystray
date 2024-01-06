package icons

func GetDefaultIcon() []byte {
	file, err := file.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return file
}
