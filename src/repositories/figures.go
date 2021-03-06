package repositories

import (
	"context"

	"github.com/charoleizer/my-collection/src/models"
	"github.com/charoleizer/my-collection/src/storage"
)

type IFigures interface {
	Save() error
}

type Figures struct {
	*models.Figures
}

func (f *Figures) Save() error {
	client, err := storage.Connect()
	if err != nil {
		return err
	}

	coll := client.Database("collections").Collection("figures")
	doc := f

	_, err = coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return err
	}

	return nil

}
