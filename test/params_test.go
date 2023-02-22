package test

import (
	"fmt"
	"testing"

	"github.com/cccaiyan/upipe-go-toolkit/atom"
)

func TestGetInput(t *testing.T) {
	// a := "[{\"id\":7154712,\"x\":3,\"y\":0,\"version\":\"6.1\",\"pipeline_info_id\":738232,\"name\":\"SITE\",\"value\":\"JDD\",\"param_type\":\"select\",\"identifier\":\"JDos\"},{\"id\":7154713,\"x\":3,\"y\":0,\"version\":\"6.1\",\"pipeline_info_id\":738232,\"name\":\"SYSTEM_NAME\",\"value\":\"${globalParams.user.JDOS_SYSTEM}\",\"param_type\":\"string\",\"identifier\":\"JDos\"},{\"id\":7154714,\"x\":3,\"y\":0,\"version\":\"6.1\",\"pipeline_info_id\":738232,\"name\":\"APP_NAME\",\"value\":\"${globalParams.user.JDOS_APP}\",\"param_type\":\"string\",\"identifier\":\"JDos\"},{\"id\":7154715,\"x\":3,\"y\":0,\"version\":\"6.1\",\"pipeline_info_id\":738232,\"name\":\"GROUP_NAME\",\"value\":\"${globalParams.user.JDOS_GROUP}\",\"param_type\":\"string\",\"identifier\":\"JDos\"}]"
	a := atom.GetInput("DEPLOY_AUTO_CONFIRM")
	b := atom.GetInput("HTTP_HEADER")
	c := atom.GetInput("SITE")
	fmt.Printf("a: %T", a)
	fmt.Printf("b: %T", b)
	fmt.Printf("c: %T", c)
}
