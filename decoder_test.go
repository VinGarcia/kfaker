package kfaker_test

import (
	"testing"

	"github.com/vingarcia/kfaker"
	tt "github.com/vingarcia/kfaker/internal/testtools"
)

func TestDecoder(t *testing.T) {
	t.Run("should unmarshal url values correctly", func(t *testing.T) {

		var dto struct {
			Name string `schema:"name"`
			Type string `schema:"type"`
		}
		err := kfaker.Fake(&dto, map[string]any{
			"name": "fakeName",
			"type": "fakeType",
		})
		tt.AssertNoErr(t, err)

		tt.AssertEqual(t, dto.Name, "fakeName")
		tt.AssertEqual(t, dto.Type, "fakeType")
	})

	t.Run("should unmarshal slices of strings", func(t *testing.T) {
		var dto struct {
			Names []string `schema:"name,required"`
			Type  string   `schema:"type,required"`
		}
		err := kfaker.Fake(&dto, map[string]any{
			"name": []string{"fakeName1", "fakeName2"},
			"type": []string{"fakeType"},
		})
		tt.AssertNoErr(t, err)

		tt.AssertEqual(t, dto.Names, []string{"fakeName1", "fakeName2"})
		tt.AssertEqual(t, dto.Type, "fakeType")
	})

	t.Run("should unmarshal required values correctly", func(t *testing.T) {
		var dto struct {
			Name string `schema:"name,required"`
			Type string `schema:"type,required"`
		}
		err := kfaker.Fake(&dto, map[string]any{
			"name": []string{"fakeName"},
			"type": []string{"fakeType"},
		})
		tt.AssertNoErr(t, err)

		tt.AssertEqual(t, dto.Name, "fakeName")
		tt.AssertEqual(t, dto.Type, "fakeType")
	})

	t.Run("should return an error if a required field is missing", func(t *testing.T) {
		var dto struct {
			Name string `schema:"name,required"`
			Type string `schema:"type,required"`
		}
		err := kfaker.Fake(&dto, map[string]any{
			"type": []string{"fakeType"},
		})
		tt.AssertErrContains(t, err, "missing", "query param", "name")
	})

	t.Run("should unmarshal any types != string using yaml.Unmarshal", func(t *testing.T) {
		var dto struct {
			SomeInt int `schema:"someInt"`
		}
		err := kfaker.Fake(&dto, map[string]any{
			"someInt": []string{"42"},
		})
		tt.AssertNoErr(t, err)

		tt.AssertEqual(t, dto.SomeInt, 42)
	})
}
