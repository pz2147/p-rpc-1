package test

import (
	"fmt"
	"regexp"
	"testing"
)

var traceCtxRegExp = regexp.MustCompile("^(?P<version>[0-9a-f]{2})-(?P<traceID>[a-f0-9]{32})-(?P<spanID>[a-f0-9]{16})-(?P<traceFlags>[a-f0-9]{2})(?:-.*)?$")


//华为登录
func TestES(t *testing.T) {
	//var ctx context.Context
	//l := logic.NewESGuideLogic(ctx, testCtx)
	//resp, err := l.ESGuide(&prpc1client.EmptyReq{})
	//if err != nil {
	//	fmt.Println("失败:", err)
	//	return
	//}
	//
	//fmt.Println(resp)

	h:= "15-577eb1e05833e7a3191af2d3053abc"
	matches := traceCtxRegExp.FindStringSubmatch(h)
	fmt.Printf("%v", matches)


}
