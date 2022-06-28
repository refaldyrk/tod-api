package data

import (
	"context"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"tod/helper"
)

type Repository interface {
	GetAllTruthData(ctx context.Context) ([]DataApiModel, error)
	GetAllDareData(ctx context.Context) ([]DataApiModel, error)
	CreateData(ctx context.Context, input DataInput) (DataApiModel, error)
}

type repository struct {
	client *qmgo.Collection
}

func NewRepository(client *qmgo.Collection) *repository {
	return &repository{
		client: client,
	}
}

func (r *repository) GetAllTruthData(ctx context.Context) ([]DataApiModel, error) {
	var models []DataApiModel
	err := r.client.Find(ctx, bson.M{"type": "truth"}).All(&models)
	if err != nil {
		go helper.SendNotify("Error Get Data", err.Error())
		return nil, err
	}

	return models, nil
}

func (r *repository) GetAllDareData(ctx context.Context) ([]DataApiModel, error) {
	var models []DataApiModel
	err := r.client.Find(ctx, bson.M{"type": "dare"}).All(&models)
	if err != nil {
		go helper.SendNotify("Error Get Data", err.Error())
		return nil, err
	}

	return models, nil
}

func (r *repository) CreateData(ctx context.Context, input DataInput) (DataApiModel, error) {
	data := DataApiModel{
		Type: input.Type,
		Data: input.Data,
	}
	_, err := r.client.InsertOne(ctx, data)
	if err != nil {
		go helper.SendNotify("Error Insert Data", err.Error())
		return DataApiModel{}, err
	}
	return data, nil
}
