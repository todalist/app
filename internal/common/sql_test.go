package common

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Pager
	TestArg  *string
	TestArg2 int
	Arg3     *string
	Arg5     string
	Ignore   string `cond:"-"`
}

func TestToWhere(t *testing.T) {
	nonEmptyStr := "not empty"
	emptyStr := ""
	tests := []struct {
		name     string
		querier  TestStruct
		expected string
		argMap   map[string]interface{}
	}{
		{"allFieldsNil", TestStruct{}, "", map[string]interface{}{}},
		{"someFieldsNil", TestStruct{TestArg: nil, TestArg2: 0, Arg3: nil, Arg5: emptyStr}, "", map[string]interface{}{}},
		{"oneFiled", TestStruct{TestArg: &nonEmptyStr, TestArg2: 0, Arg3: nil, Arg5: emptyStr},
			"test_arg = @TestArg",
			map[string]interface{}{"TestArg": &nonEmptyStr}},
		{"twoFields", TestStruct{TestArg: &nonEmptyStr, TestArg2: 1, Arg3: nil, Arg5: emptyStr},
			"test_arg = @TestArg AND test_arg2 = @TestArg2",
			map[string]interface{}{"TestArg": &nonEmptyStr, "TestArg2": 1}},
		{"threeFields", TestStruct{TestArg: &nonEmptyStr, TestArg2: 1, Arg3: &nonEmptyStr, Arg5: emptyStr},
			"test_arg = @TestArg AND test_arg2 = @TestArg2 AND arg3 = @Arg3",
			map[string]interface{}{"TestArg": &nonEmptyStr, "TestArg2": 1, "Arg3": &nonEmptyStr}},
		{"fourFields", TestStruct{TestArg: &nonEmptyStr, TestArg2: 1, Arg3: &nonEmptyStr, Arg5: nonEmptyStr},
			"test_arg = @TestArg AND test_arg2 = @TestArg2 AND arg3 = @Arg3 AND arg5 = @Arg5",
			map[string]interface{}{"TestArg": &nonEmptyStr, "TestArg2": 1, "Arg3": &nonEmptyStr, "Arg5": nonEmptyStr}},
		{"fiveFields", TestStruct{TestArg: &nonEmptyStr, TestArg2: 1, Arg3: &nonEmptyStr, Arg5: nonEmptyStr},
			"test_arg = @TestArg AND test_arg2 = @TestArg2 AND arg3 = @Arg3 AND arg5 = @Arg5",
			map[string]interface{}{"TestArg": &nonEmptyStr, "TestArg2": 1, "Arg3": &nonEmptyStr, "Arg5": nonEmptyStr}},
		{"testIgnore", TestStruct{Ignore: nonEmptyStr}, "", map[string]interface{}{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, argMap := QuerierToSqlCondition(nil, &tt.querier)
			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.argMap, *argMap)
		})
	}
}

func TestFieldToSqlName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"emptyString", "", ""},
		{"allLowercase", "namestring", "namestring"},
		{"allUppercase", "NAMESTRING", "n_a_m_e_s_t_r_i_n_g"},
		{"mixedCase", "NameString", "name_string"},
		{"startsWithUppercase", "AnotherName", "another_name"},
		{"numberIncluded", "Name1String", "name1_string"},
		{"symbolIncluded", "Name$tring", "name$tring"},
		{"multipleUppercaseTogether", "NaMeSTrING", "na_me_s_tr_i_n_g"},
		{"noUppercase", "nouppercase", "nouppercase"},
		{"uppercaseAtEnd", "nameendUP", "nameend_u_p"},
		{"multipleUppercaseAtEnd", "nameendUPPER", "nameend_u_p_p_e_r"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := FieldToSqlName(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
