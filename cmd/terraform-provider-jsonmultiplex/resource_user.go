package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	iam_v0 "github.com/json-multiplex/iam/generated/jsonmultiplex/iam/v0"
	"google.golang.org/grpc/metadata"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", client.Token),
	}))

	user, err := client.IAM.CreateUser(ctx, &iam_v0.CreateUserRequest{
		User: &iam_v0.User{
			Name:     fmt.Sprintf("users/%s", d.Get("name").(string)),
			Password: d.Get("password").(string),
		},
	})

	if err != nil {
		return err
	}

	d.SetId(user.Name)
	return nil
}

func resourceUserRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", client.Token),
	}))

	user, err := client.IAM.GetUser(ctx, &iam_v0.GetUserRequest{
		Name: d.Id(),
	})

	if err != nil {
		return err
	}

	nameSegments := strings.Split(user.Name, "/")
	d.Set("name", nameSegments[1])
	return nil
}

func resourceUserUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceUserRead(d, meta)
}

func resourceUserDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", client.Token),
	}))

	_, err := client.IAM.DeleteUser(ctx, &iam_v0.DeleteUserRequest{
		Name: fmt.Sprintf("users/%s", d.Get("name").(string)),
	})

	return err
}
