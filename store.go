package envenc

import "github.com/gouniverse/dataobject"

// store
type store struct {
	dataobject.DataObject
}

func newStore() *store {
	o := &store{}
	return o
}

func newStoreFromJSON(json string) (*store, error) {
	o := &store{}

	do, err := dataobject.NewDataObjectFromJSON(json)

	if err != nil {
		return nil, err
	}

	o.Hydrate(do.Data())

	return o, nil
}

func (o *store) ToJSON() (string, error) {
	json, err := o.DataObject.ToJSON()

	if err != nil {
		return "{}", err
	}

	if json == "null" {
		return "{}", nil
	}

	return json, nil
}

func (o *store) Remove(key string) *store {
	data := o.Data()
	delete(data, key)
	o.Hydrate(data)
	return o
}
