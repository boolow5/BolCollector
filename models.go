package main

type (
	Settings struct {
		Delay int `json:"delay"`
	}
	Selector struct {
		Base       string `json:"article"`
		TargetBase string `json:"target_base"`
		TargetText string `json:"target_text"`
		TargetLink string `json:"target_link"`
	}
	Website struct {
		Name     string    `json:"name"`
		RootUrl  string    `json:"root_url"`
		Selector *Selector `json:"selector"`
	}
)
