// Show some info to the user. Fall back to viewing on any input.

package internal

import (
	"github.com/j-ole/twin"
	log "github.com/sirupsen/logrus"
)

type PagerModeInfo struct {
	Pager  *Pager
	Text   string
	logged bool
}

func (m *PagerModeInfo) drawFooter(_ string, _ string, _ string) {
	if !m.logged {
		log.Infof("Displaying info message to user: %q", m.Text)
		m.logged = true
	}

	m.Pager.setFooter(m.Text, "", "", "")
}

func (m *PagerModeInfo) onKey(key twin.KeyCode) {
	m.Pager.mode = PagerModeViewing{m.Pager}
	m.Pager.mode.onKey(key)
}

func (m *PagerModeInfo) onRune(char rune) {
	m.Pager.mode = PagerModeViewing{m.Pager}
	m.Pager.mode.onRune(char)
}
