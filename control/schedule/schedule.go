package schedule

import (
	"github.com/kassybas/reeu/models/resource"
)

type Scheduler struct {
	inModifierChan resource.Modifier
}

func (s *Scheduler) AddModifier(res *resource.Resource, mod resource.Modifier) {

}
