package main 

import (
	"fmt"
	t "github.com/dvasilen/razee-rrs3/pkg/client/clientset/versioned/typed/razee/v1alpha2"
	v1alpha2 "github.com/dvasilen/razee-rrs3/pkg/apis/razee/v1alpha2"
)

func main() {
	fmt.Printf("%v", t.DeployV1alpha2Client{})
	fmt.Printf("%v", v1alpha2.RemoteResourceS3{})
}

