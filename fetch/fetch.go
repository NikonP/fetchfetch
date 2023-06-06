package fetch

import "os/exec"

type Fetch struct {
	Name        string
	Description string
}

func (fetch *Fetch) FetchExists() bool {
	_, err := exec.LookPath(fetch.Name)
	return err == nil
}
