package target

var InputParamsTestCases = []struct {
	name        string
	params      *InputParams
	field       string
	validateTag string
	wantErr     bool
	equivalency EquivalencyType
	boundary    BoundaryType
}{
	{
		name:        "Name is empty",
		params:      &InputParams{Name: ""},
		field:       "Name",
		validateTag: "required",
		wantErr:     true,
		equivalency: EquivalencyInvalid,
		boundary:    BoundaryNone,
	},
	{
		name:        "Name length lower below",
		params:      &InputParams{Name: "aa"},
		field:       "Name",
		validateTag: "min=3",
		wantErr:     true,
		equivalency: EquivalencyInvalid,
		boundary:    BoundaryLowerBelow,
	},
	{
		name:        "Name length lower",
		params:      &InputParams{Name: "aaa"},
		field:       "Name",
		validateTag: "min=3",
		wantErr:     false,
		equivalency: EquivalencyValid,
		boundary:    BoundaryLower,
	},
	{
		name:        "Name length lower above",
		params:      &InputParams{Name: "aaaa"},
		field:       "Name",
		validateTag: "min=3",
		wantErr:     false,
		equivalency: EquivalencyValid,
		boundary:    BoundaryLowerAbove,
	},
	{
		name:        "Name length upper below",
		params:      &InputParams{Name: "aaaaaaaaa"},
		field:       "Name",
		validateTag: "max=10",
		wantErr:     false,
		equivalency: EquivalencyValid,
		boundary:    BoundaryUpperBelow,
	},
	{
		name:        "Name length upper",
		params:      &InputParams{Name: "aaaaaaaaaa"},
		field:       "Name",
		validateTag: "max=10",
		wantErr:     false,
		equivalency: EquivalencyValid,
		boundary:    BoundaryUpper,
	},
	{
		name:        "Name length upper above",
		params:      &InputParams{Name: "aaaaaaaaaaa"},
		field:       "",
		validateTag: "max=10",
		wantErr:     true,
		equivalency: EquivalencyInvalid,
		boundary:    BoundaryUpperAbove,
	},
	{
		name:        "Email is empty",
		params:      &InputParams{Email: ""},
		field:       "Email",
		validateTag: "required",
		wantErr:     true,
		equivalency: EquivalencyInvalid,
		boundary:    BoundaryNone,
	},
	{
		name:        "Email invalid email",
		params:      &InputParams{Email: "notanemail"},
		field:       "Email",
		validateTag: "email",
		wantErr:     true,
		equivalency: EquivalencyInvalid,
		boundary:    BoundaryNone,
	},
	{
		name:        "Email valid email",
		params:      &InputParams{Email: "a@b.com"},
		field:       "Email",
		validateTag: "email",
		wantErr:     false,
		equivalency: EquivalencyValid,
		boundary:    BoundaryNone,
	},
}
