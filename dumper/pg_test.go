package dumper

import (
	"bytes"
	"testing"
)

var pgValueTests = []struct {
	in            []byte
	expected      []DumpValue
	expectedQuery []DumpValue
}{
	{
		[]byte{
			0x00, 0x00, 0x00, 0x52, 0x00, 0x03, 0x00, 0x00, 0x75, 0x73, 0x65, 0x72, 0x00, 0x70, 0x6f, 0x73,
			0x74, 0x67, 0x72, 0x65, 0x73, 0x00, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x00, 0x74,
			0x65, 0x73, 0x74, 0x64, 0x62, 0x00, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
			0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x00, 0x70, 0x73, 0x71, 0x6c, 0x00, 0x63, 0x6c, 0x69, 0x65,
			0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x00, 0x55, 0x54, 0x46, 0x38,
			0x00, 0x00,
		},
		[]DumpValue{
			DumpValue{
				Key:   "username",
				Value: "postgres",
			},
			DumpValue{
				Key:   "database",
				Value: "testdb",
			},
		},
		[]DumpValue{},
	},
	{
		[]byte{
			0x51, 0x00, 0x00, 0x00, 0x19, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x20, 0x2a, 0x20, 0x46, 0x52,
			0x4f, 0x4d, 0x20, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3b, 0x00,
		},
		[]DumpValue{},
		[]DumpValue{
			DumpValue{
				Key:   "query",
				Value: "SELECT * FROM users;",
			},
			DumpValue{
				Key:   "message_type",
				Value: "Q",
			},
		},
	},
}

func TestPgReadPersistentValuesStartupMessage(t *testing.T) {
	for _, tt := range pgValueTests {
		out := new(bytes.Buffer)
		dumper := &PgDumper{
			logger: NewTestLogger(out),
		}
		in := tt.in

		actual := dumper.ReadPersistentValues(in)
		expected := tt.expected

		if len(actual) != len(expected) {
			t.Errorf("actual %v\nwant %v", actual, expected)
		}
		if len(actual) == 2 {
			if actual[0] != expected[0] {
				t.Errorf("actual %v\nwant %v", actual, expected)
			}
			if actual[1] != expected[1] {
				t.Errorf("actual %v\nwant %v", actual, expected)
			}
		}
	}
}

func TestPgRead(t *testing.T) {
	for _, tt := range pgValueTests {
		out := new(bytes.Buffer)
		dumper := &PgDumper{
			logger: NewTestLogger(out),
		}
		in := tt.in

		actual := dumper.Read(in)
		expected := tt.expectedQuery

		if len(actual) != len(expected) {
			t.Errorf("actual %v\nwant %v", actual, expected)
		}
		if len(actual) == 2 {
			if actual[0] != expected[0] {
				t.Errorf("actual %#v\nwant %#v", actual[0], expected[0])
			}
			if actual[1] != expected[1] {
				t.Errorf("actual %#v\nwant %#v", actual[1], expected[1])
			}
		}
	}
}