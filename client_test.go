// Author: yann
// Date: 2022/5/22
// Desc: hdfsrest

package hdfsrest

var cli *client

func init() {
	cli = New(&ClientOption{
		DataNodes: []string{"10.1.141.215:14000"},
		User:      "root",
	})
}
