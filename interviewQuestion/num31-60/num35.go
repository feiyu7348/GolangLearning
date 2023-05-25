// @author:zfy
// @date:2023/5/25 23:31

package main

type Fragment interface {
	Exec(transInfo *TransInfo) error
}
type GetPodAction struct {
}

func (g GetPodAction) Exec(transInfo *TransInfo) error {
	...
	return nil
}

var fragment Fragment = new(GetPodAction)

//var fragment Fragment = &GetPodAction{}
//var fragment Fragment = GetPodAction{}
