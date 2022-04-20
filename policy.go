package mdschema

type IntPolicy struct {
	Start, Stop, Step int
	Required          bool
	Whitelist         map[int]bool
	Blacklist         map[int]bool
}

func (ip IntPolicy) Validate(v interface{}) bool {
	in := v.(int)

	valid := make([]int, 0)
	for i := ip.Start; i <= ip.Stop; i += ip.Step {
		valid = append(valid, i)
	}

	if ip.Blacklist[in] {
		return false
	}

	if ip.Whitelist[in] {
		return true
	}

	for _, i := range valid {
		if i == in {
			return true
		}
	}

	return false
}

type StringPolicy struct {
	Required  bool
	Length    int
	Pattern   string
	Whitelist map[int]string
	Blacklist map[int]string
}

func (sp StringPolicy) Validate(v interface{}) bool {
	return true
}

type ArrayPolicy struct {
	Required bool
	Members  YAMLType
	Length   int
	Allowed  []interface{}
}

func (ap ArrayPolicy) Validate(v interface{}) bool {
	return true
}

type ObjectPolicy struct {
	Required bool
	// this one is the most out of control and i genuinely don't think i have the chops to make this work good enough to publish. sorry.
}

func (op ObjectPolicy) Validate(v interface{}) bool {
	return true
}

type Policy interface {
	Validate(interface{}) bool
}

// SchemaKey is the atom of the config. This represents the metadata key you want to apply a policy to.
type SchemaKey struct {
	Type     YAMLType `toml:"type"`
	Optional bool     `toml:"optional"`
	Int      IntPolicy
	String   StringPolicy
	Array    ArrayPolicy
	Object   ObjectPolicy
}

// Policies
func (sk SchemaKey) Policy() Policy {
	switch sk.Type {
	case YAMLString:
		return sk.String
	case YAMLInteger:
		return sk.Int
	case YAMLArray:
		return sk.Array
	case YAMLObject:
		return sk.Object

	}
	return nil
}
