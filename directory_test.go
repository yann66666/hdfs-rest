// Author: yann
// Date: 2022/5/22
// Desc: hdfsrest

package hdfsrest

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestClient_Mkdir(t *testing.T) {
	err := cli.Mkdir("/test/test2", os.ModePerm)
	assert.Equal(t, err, nil, "创建目录失败")
}
