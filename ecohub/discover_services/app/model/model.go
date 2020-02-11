package model

type PostData struct {
	Type string `json:"type"`
        Connection_type string `json:"connection_type"`
	Parameters []struct {
                        Name string `json:"name"`
                        Value string `json:"value"`
                }`json:"parameters"`

}



