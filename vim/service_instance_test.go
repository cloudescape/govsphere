package vim

import (
	"testing"
)

func TestNewServiceInstance(t *testing.T) {
	service := NewServiceInstance("https://172.16.103.128", "root", "shadow04101", "5.5", false)
	content, err := service.Content()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	cap, err := service.Capability()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	t.Logf("=>%#v<=\n", cap)
	t.Logf("=>%#v<=\n", content.About)
	t.Log("Puta!")
}