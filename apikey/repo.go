package apikey

import (
	"context"
	"tod/helper"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

type Repository interface {
	GetAllApikey(ctx context.Context) ([]ApikeyModel, error)
	CreateApikey(ctx context.Context, input ApikeyInput) (ApikeyModel, error)
	DeleteApikey(ctx context.Context, key string) error
	CheckApikey(ctx context.Context, key string) (bool, error)
}

type repository struct {
	client *qmgo.Collection
}

func NewRepo(client *qmgo.Collection) *repository {
	return &repository{client: client}
}

func (r *repository) CreateApikey(ctx context.Context, input ApikeyInput) (ApikeyModel, error) {
	model := ApikeyModel{
		Owner: input.Owner,
		Key:   input.Key,
	}
	_, err := r.client.InsertOne(ctx, model)
	if err != nil {
		go helper.SendNotify("Error Create Apikey", err.Error())
		return ApikeyModel{}, err
	}

	go func() {
		helper.SendNotify("New Apikey", "New Apikey Created\n"+model.Key)
	}()

	return model, nil
}

func (r *repository) DeleteApikey(ctx context.Context, key string) error {
	err := r.client.Remove(ctx, bson.M{"key": key})
	if err != nil {
		go helper.SendNotify("Error Delete Apikey", err.Error())
		return err
	}

	go func() {
		helper.SendNotify("Apikey Deleted", "Apikey Deleted\n"+key)
	}()
	return nil
}

func (r *repository) GetAllApikey(ctx context.Context) ([]ApikeyModel, error) {
	var models []ApikeyModel
	err := r.client.Find(ctx, bson.M{}).All(&models)
	if err != nil {
		go helper.SendNotify("Error Get All Apikey", err.Error())
		return nil, err
	}
	return models, nil
}

func (r *repository) CheckApikey(ctx context.Context, key string) (bool, error) {
	var model ApikeyModel
	err := r.client.Find(ctx, bson.M{"key": key}).One(&model)
	if err != nil {
		go helper.SendNotify("Error Check Apikey", err.Error())
		return false, err
	}
	return true, nil
}
