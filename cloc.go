package main

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"sort"
	"strings"
)

// test if the comment can count

/*
hello world
if this can cal
*/

func main() {
	args := os.Args
	localPath := "."
	languages := getDefaultLanguages()
	if len(args) == 2 {
		localPath = args[1]
	}

	analyzeFile(localPath, languages)
	var result []Language
	for _, v := range languages {
		if len(v.Files) > 0 {
			result = append(result, *v)
		}
	}
	printResult(result)
}

func analyzeFile(localPath string, languages map[string]*Language) {
	filesEntry, err := os.ReadDir(localPath)
	if err != nil {
		log.Error(err)
	}

	for _, file := range filesEntry {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		if !file.IsDir() {
			f, err := os.Open(path.Join(localPath, file.Name()))
			if err != nil {
				log.Error(err)
			}

			ext := getFileType(file.Name())
			if ext == "" || languages[ext] == nil {
				continue
			}
			languages[ext].Files = append(languages[ext].Files, path.Join(localPath, file.Name()))
			if lan, ok := languages[ext]; ok {
				sc := bufio.NewScanner(f)
				statusIsCode := true
				lastComment := 0

				for sc.Scan() {
					text := sc.Text()
					pureText := strings.Trim(text, " ")
					commented := false
					if !statusIsCode {
						languages[ext].comments += 1
						if strings.HasPrefix(pureText, lan.multiComment[lastComment][1]) {
							statusIsCode = true
							lastComment = 0
						}
						continue
					}

					if pureText == "" {
						languages[ext].blanks += 1
						continue
					}

					for i := 0; i < len(lan.comment); i++ {
						if strings.HasPrefix(pureText, lan.comment[i]) {
							languages[ext].comments += 1
							commented = true
							continue
						}
					}
					if commented {
						continue
					}
					for i := 0; i < len(lan.multiComment); i++ {
						if lan.multiComment[i][0] == "" {
							continue
						}
						if strings.HasPrefix(pureText, lan.multiComment[i][0]) {
							languages[ext].comments += 1
							statusIsCode = false
							lastComment = i
							commented = true
							continue
						}
					}
					if commented {
						continue
					}
					languages[ext].codes += 1

				}
			}
		} else {
			analyzeFile(path.Join(localPath, file.Name()), languages)
		}
	}
}

func printResult(languages []Language) {
	sort.Sort(Languages(languages))
	log.Println("------------------------------------")
	log.Printf("%10s\tfile\tblank\tcomment\tcode", "language")
	for i := 0; i < len(languages); i++ {
		log.Printf("%10s\t%d\t%d\t%d\t%d", languages[i].name, len(languages[i].Files),
			languages[i].blanks, languages[i].comments, languages[i].codes)
	}
	log.Println("------------------------------------")
}
