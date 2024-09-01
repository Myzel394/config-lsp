package common

import (
	"os"
	"strings"
)

type PasswdInfo struct {
	Name     string
	UID      string
	GID      string
	HomePath string
}

var _cachedPasswdInfo []PasswdInfo

func FetchPasswdInfo() ([]PasswdInfo, error) {
	if len(_cachedPasswdInfo) > 0 {
		return _cachedPasswdInfo, nil
	}

	readBytes, err := os.ReadFile("/etc/passwd")

	if err != nil {
		return []PasswdInfo{}, err
	}

	lines := strings.Split(string(readBytes), "\n")
	infos := make([]PasswdInfo, 0)

	for _, line := range lines {
		splitted := strings.Split(line, ":")

		if len(splitted) < 6 {
			continue
		}

		info := PasswdInfo{
			Name:     splitted[0],
			UID:      splitted[2],
			GID:      splitted[3],
			HomePath: splitted[5],
		}

		infos = append(infos, info)
	}

	_cachedPasswdInfo = infos

	return infos, nil
}

type GroupInfo struct {
	Name string
	GID  string
}

var _cachedGroupInfo []GroupInfo

func FetchGroupInfo() ([]GroupInfo, error) {
	if len(_cachedGroupInfo) > 0 {
		return _cachedGroupInfo, nil
	}

	readBytes, err := os.ReadFile("/etc/group")

	if err != nil {
		return []GroupInfo{}, err
	}

	lines := strings.Split(string(readBytes), "\n")
	infos := make([]GroupInfo, 0)

	for _, line := range lines {
		splitted := strings.Split(line, ":")

		if len(splitted) < 3 {
			continue
		}

		info := GroupInfo{
			Name: splitted[0],
			GID:  splitted[2],
		}

		infos = append(infos, info)
	}

	_cachedGroupInfo = infos

	return infos, nil
}
