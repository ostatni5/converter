package project

import (
	"testing"

	"gopkg.in/mgo.v2/bson"

	"github.com/Palantir/palantir/model/test"
)

var testCases test.MarshallingCases = test.MarshallingCases{
	{
		&Version{ID: 1,
			Settings: "setId",
			SetupID:  bson.ObjectIdHex("bbbbbbbbbbbbbbbbbbbbbbbb"),
			Results:  "resultsId",
		},
		`{
			"id": 1,
			"settings": "setId",
			"setupId": "bbbbbbbbbbbbbbbbbbbbbbbb",
			"resultsId": "resultsId"
		}`,
	},

	{
		&Project{
			ID:          bson.ObjectIdHex("58cfd607dc25403a3b691781"),
			AccountID:   bson.ObjectIdHex("cccccccccccccccccccccccc"),
			Name:        "name",
			Description: "description",
			Versions:    []Version{},
		},
		`{
			"id": "58cfd607dc25403a3b691781",
			"accountId": "cccccccccccccccccccccccc",
			"name": "name",
			"description": "description",
			"versions": []
		}`,
	},

	{
		&List{[]Project{}},
		`{"projects":[]}`,
	},
}

func TestMarshal(t *testing.T) {
	test.Marshal(t, testCases)
}

func TestUnmarshal(t *testing.T) {
	test.Unmarshal(t, testCases)
}

func TestUnmarshalMarshalled(t *testing.T) {
	test.UnmarshalMarshalled(t, testCases)
}

func TestMarshalUnmarshalled(t *testing.T) {
	test.MarshalUnmarshalled(t, testCases)
}
