package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func GetSSHDir(configs bool) string {
	var result string
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Unable to fetch home dir")
	}
	if configs {
		result = home + "/.ssh/config.d"
	} else {
		result = home + "/.ssh"
	}

	return result
}

func GetSSHFile() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Unable to fetch home dir")
	}
	return home + "/.ssh/config"
}

func ReadSSHConfig(path string) string {
	CheckIfExists(path)
	file, err := os.Open(path + "/config")
	if err != nil {
		log.Fatal("Error Reading Config")
	}
	defer file.Close()
	buffer := make([]byte, 1024)
	var data string
	for {
		n, err := file.Read(buffer)
		if err != nil {
			break // End of file reached or an error occurred
		}
		data = string(buffer[:n])
	}
	return data
}

func CheckIfExists(basepath string) {
	if _, err := os.Stat(basepath + "/config"); os.IsNotExist(err) {
		log.Println("CONFIG FILE DOESN'T EXIST")
		file, err := os.Create(basepath + "/config")
		if err != nil {
			log.Fatal("Error creating file:", err)
			return
		}
		defer file.Close()
	}
	if _, err := os.Stat(basepath + "/config.d"); os.IsNotExist(err) {
		log.Println("CONFIG.D DIR DOESN'T EXIST")
		err := os.Mkdir(basepath+"/config.d", 0755)
		if err != nil {
			log.Fatal("Error creating directory:", err)
			return
		}
	}
}

func GetSSHConfigs(withBase bool) []string {
	var items []string
	configs, err := os.ReadDir(GetSSHDir(true))
	if err != nil {
		log.Fatal("Error Reading configs")
	}
	for _, config := range configs {
		if !withBase {
			items = append(items, config.Name())
		} else {
			items = append(items, GetSSHDir(true)+"/"+config.Name())
		}
	}
	return items
}

func BasicPrompt(prompt string, nullable bool) string {
	var input string
	for {
		fmt.Println(prompt)
		fmt.Print("> ")
		fmt.Scanln(&input)
		if input != "" || nullable {
			break
		}
	}
	return input
}

func GetCategories() []string {
	var categories []string
	re := regexp.MustCompile(`(?m)^#.+`)
	categories = append(categories, re.FindAllString(ReadSSHConfig(GetSSHDir(false)), -1)...)
	return categories
}

func GetSSHKeys() []string {
	var categories []string
	files, err := os.ReadDir(GetSSHDir(false))
	if err != nil {
		return categories
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".pem") {
			categories = append(categories, file.Name())
		}
	}
	return categories
}

func CreateConfig(name, config string) {
	file, err := os.Create(GetSSHDir(true) + "/" + name)
	if err != nil {
		log.Fatal("Error Saving Config")
	}
	defer file.Close()
	file.Write([]byte(config))
}

func LinkConfig(name, category string) {
	// TODO: Add backup of config file if something goes wrong
	file, err := os.OpenFile(GetSSHFile(), os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	pattern := regexp.MustCompile(`(?m)(?s)` + category + `\s.*?(?:^\B|\z)`)
	match := pattern.Find(content)
	if len(match) != 0 {
		var newBlock string
		if strings.HasSuffix(string(match), "\n") {
			newBlock = string(match) + "Include config.d/" + name + "\n"
		} else {
			newBlock = string(match) + "\nInclude config.d/" + name + "\n"
		}
		content = []byte(strings.Replace(string(content), string(match), newBlock, 1))
	} else {
		contentStr := string(content)
		lines := strings.Split(contentStr, "\n")
		lastLine := lines[len(lines)-1]
		if strings.TrimSpace(lastLine) == "" {
			content = append(content, []byte(fmt.Sprintf("\n# %v\nInclude config.d/%v\n", strings.ToUpper(category), name))...)
		} else {
			content = append(content, []byte(fmt.Sprintf("\n\n# %v\nInclude config.d/%v\n", strings.ToUpper(category), name))...)
		}
	}
	file.WriteAt(content, 0)
}
