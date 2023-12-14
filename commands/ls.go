package commands

func (fs *FileSystem) Ls(path string) []string {
	var target *Directory
	if path == "" || path == "/" {
		target = fs.Root
	} else {
		target = fs.GetDir(path)
		if target == nil {
			return []string{"Directory not found: " + path}
		}
	}

	var contents []string
	for name := range target.SubDirs {
		contents = append(contents, name+"/")
	}
	for name := range target.Files {
		contents = append(contents, name)
	}
	return contents
}
