package behavior_pattern

import "testing"

func TestTemplatePattern(t *testing.T) {
	NewBaseTemplate(NewFast()).Do()
	NewBaseTemplate(NewSlow()).Do()
}
