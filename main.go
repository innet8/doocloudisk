// Code generated by hertz generator.

package main

import (
	_ "github.com/cloudisk/biz/dal"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()

	register(h)
	h.Spin()
}
