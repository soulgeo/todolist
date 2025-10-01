package config

import "time"

func GetSelected() (string, error) {
	c, err := ReadConfig()
	if err != nil {
		return "", err
	}
	return c.Selected, nil
}

func SetSelected(name string) error {
	c, err := ReadConfig()
	if err != nil {
		return err
	}
	c.Selected = name
	c.LastOpened = time.Now()
	return WriteConfig(c)
}
