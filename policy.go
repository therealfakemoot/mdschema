package mdschema

type IntPolicy struct {
	Start, End, Step int
}

type StringPolicy struct {
	Length    int
	Pattern   string
	Whitelist []string
	Blacklist []string
}

type ArrayPolicy struct {
	Members YAMLType
	Length  int
	Allowed []interface{}
}

type ObjectPolicy struct {
	// this one is the most out of control and i genuinely don't think i have the chops to make this work good enough to publish. sorry.
}

type Policy struct {
	Type     YAMLType `toml:"type"`
	Optional bool     `toml:"optional"`
	Int      IntPolicy
	String   StringPolicy
	Array    ArrayPolicy
	Ojbect   ObjectPolicy
}
