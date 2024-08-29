package entity

type DatabaseMetadata struct {
	Databases []Database `json:"databases"`
}

type Database struct {
	Name   string          `json:"name"`
	Tables []DatabaseTable `json:"tables"`
}

type DatabaseTable struct {
	Name    string           `json:"name"`
	Columns []DatabaseColumn `json:"columns"`
	Indexes []DatabaseIndex  `json:"indexes"`
}

type DatabaseColumn struct {
	Name       string           `json:"name"`
	Type       string           `json:"type"`
	PrimaryKey bool             `json:"primary_key,omitempty"`
	Null       bool             `json:"null"`
	Key        string           `json:"key,omitempty"`
	Default    string           `json:"default"`
	Extra      string           `json:"extra,omitempty"`
	Children   []DatabaseColumn `json:"children,omitempty"`
	Hash       string           `json:"hash,omitempty"`
}

type DatabaseIndex struct {
	Name    string                `json:"name"`
	Type    string                `json:"type,omitempty"`
	Columns []DatabaseIndexColumn `json:"columns"`
	Unique  bool                  `json:"unique"`
	Hash    string                `json:"hash,omitempty"`
}

type DatabaseIndexColumn struct {
	Name  string `json:"name"`
	Order int    `json:"order"`
}

func (col *DatabaseIndexColumn) OrderString() string {
	if col.Order < 0 {
		return "DESC"
	} else {
		return "ASC"
	}
}
