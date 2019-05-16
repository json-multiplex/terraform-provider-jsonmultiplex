package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	iam_v0 "github.com/json-multiplex/iam/generated/jsonmultiplex/iam/v0"
	"google.golang.org/grpc"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the JSON-Multiplex account to operate on.",
			},
			"user_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The user ID for API operations.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The password for API operations.",
			},
			"iam_uri": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The endpoint for IAM API operations.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"jsonmultiplex_user": resourceUser(),
		},
		ConfigureFunc: providerConfigure,
	}
}

type Client struct {
	IAM   iam_v0.IAMClient
	Token string
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	iamURI := d.Get("iam_uri").(string)
	accountID := d.Get("account_id").(string)
	userID := d.Get("user_id").(string)
	password := d.Get("password").(string)

	iamConn, err := grpc.Dial(iamURI, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("cannot connect to iam: %v", err)
	}

	iam := iam_v0.NewIAMClient(iamConn)
	session, err := iam.CreateSession(context.Background(), &iam_v0.CreateSessionRequest{
		Session: &iam_v0.Session{
			Account:  fmt.Sprintf("accounts/%s", accountID),
			User:     fmt.Sprintf("users/%s", userID),
			Password: password,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("cannot create session: %v", err)
	}

	return Client{
		IAM:   iam,
		Token: session.Token,
	}, nil
}
