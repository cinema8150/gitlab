package git

import (
	"errors"
	"gitlab/tools"
	"log"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
)

func init() {

}

// CurrentRepoName 当前仓库名称
func CurrentRepoName() (string, error) {
	url, err := CurrentRepoURL()
	if err == nil {
		reg := regexp.MustCompile(`/[\w]+\.git`)
		name := reg.FindString(url)
		if len(name) > 0 {
			name = strings.ReplaceAll(name, "/", "")
			name = strings.ReplaceAll(name, ".git", "")
			return name, nil
		}
	}

	//未设置remote url时提取文件名
	rootPath, err := tools.Exec("git rev-parse --show-toplevel")
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Base(rootPath), nil
}

// CurrentRepoHost 当前仓库域名
func CurrentRepoHost() (string, error) {
	urlStr, err := CurrentRepoURL()
	if err == nil {
		URL, err := url.Parse(urlStr)
		if err != nil {
			return "", errors.New("url 格式化错误：" + err.Error())
		}
		println(URL.Host)
		return URL.Host, nil
	}
	return "", err
}

// CurrentRepoURL 当前仓库地址
func CurrentRepoURL() (string, error) {
	remtoeInfo, err := tools.Exec("git remote -v")
	if err != nil {
		log.Fatal(err)
	}

	unsetErr := errors.New("unset remote url, you can use `git remote add [<options>] <name> <url>` to set it")
	if len(remtoeInfo) == 0 {
		return "", unsetErr
	}

	remote := ""
	remotes := strings.Split(remtoeInfo, "\n")
	if len(remotes) > 0 {
		reg := regexp.MustCompile(`[\w]+`)
		remote = reg.FindString(remotes[0])
	}
	if len(remote) == 0 {
		return "", unsetErr
	}

	info, err := tools.Exec("git remote get-url " + remote)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(info), nil
}
