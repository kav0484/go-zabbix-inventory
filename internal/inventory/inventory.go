package inventory

type Invetory struct {
	Meta struct {
		Hostvars struct {
			hostname struct {
				ansible_host string `json:"ansible_host"`
				available    string `json:"available"`
				visible_nama string `json:visible_name`
			} `json:"hostvars"`
		} `json:"hostvars"`
	} `json:"_meta"`
	All struct {
		Hosts    []string `json:hosts`    //название хостов
		children []string `json:children` //название груп
	} `json:"all"`
	Ungrouped struct {
		Hosts []string `json:"hosts"`
	} `json:ungrouped`
}
