package C

import "encoding/json"

type Map map[string]any

func (c *Map) Fetch(url string) error {
	// TODO
	return nil
}

func (c Map) String() string { return c.JSON() }

func (c *Map) JSON() string {
	if c == nil {
		return "null"
	}
	byt, err := json.Marshal(c)
	if err != nil {
		return "null"
	}
	return string(byt)
}
