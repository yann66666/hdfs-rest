// Author: yann
// Date: 2022/5/22
// Desc: hdfsrest

package hdfsrest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const (
	MkdirPath = "%s/webhdfs/v1%s?op=MKDIRS&permission=%s&user.name=%s"
)

// Mkdir 创建目录并设置目录权限
func (c *client) Mkdir(dirname string, perm os.FileMode) error {
	node, err := c.getDataNode()
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf(MkdirPath, node, dirname, strconv.FormatInt(int64(perm), 8), c.user), nil)
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return GetErrFromBody(resp)
	}
	return nil
}

// MkdirAll 创建目录并设置目录权限
func (c *client) MkdirAll(dirname string, perm os.FileMode) error {
	node, err := c.getDataNode()
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf(MkdirPath, node, dirname, strconv.FormatInt(int64(perm), 8), c.user), nil)
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return GetErrFromBody(resp)
	}
	return nil
}
