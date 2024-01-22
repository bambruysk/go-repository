package main

import (
	"context"
	"fmt"
	"repo/entity"
	"repo/service/userservice"
	"repo/storage"
)

func main() {
	repoConfig := &storage.Config{StorageType: storage.StorageTypeInmemory}

	storage, err := storage.New(repoConfig)
	if err != nil {
		panic(err)
	}

	service := userservice.New(storage)

	ctx := context.Background()
	if err = service.Create(ctx, entity.User{
		ID:    1,
		Name:  "Aleksandr",
		Email: "Folomkin",
	}); err != nil {
		panic(err)
	}

	user, err := service.GetByID(ctx, 1)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
