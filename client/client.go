package client

import (
	"fmt"
	"log"
	"strings"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/account"
)

type Client struct {
	authToken string
}

type UserInfo struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewClient(token string) *Client {
	return &Client{
		authToken: token,
	}
}

func ShowError(err error) string {
	if err != nil {
		if stripeErr, ok := err.(*stripe.Error); ok {
			switch stripeErr.Code {
			case stripe.ErrorCodeAccountInvalid:
				return fmt.Sprintln("Account Invalid!")
			case stripe.ErrorCodeAPIKeyExpired:
				return fmt.Sprintln("Invalid Api Key!")
			default:
				return fmt.Sprintln("Invalid Request Url!")
			}
		}
	}
	return fmt.Sprintln("Status Ok!")
}

func (c *Client) NewItem(params *stripe.AccountParams) (*stripe.Account, error) {
	stripe.Key = c.authToken
	Id := c.GetUserId(*params.Email)
	if len(Id) == 0 {
		user, err := account.New(params)
		if err != nil {
			log.Printf("[Create Error]: %s", ShowError(err))
			return nil, err
		}
		return user, err
	}
	return nil, fmt.Errorf("user already exists")
}

func (c *Client) GetItem(Email string) (*stripe.Account, error) {
	stripe.Key = c.authToken
	Id := c.GetUserId(Email)
	if len(Id) == 0 {
		return nil, fmt.Errorf("user does not exist")
	}
	user, err := account.GetByID(
		Id,
		nil,
	)
	if err != nil {
		log.Printf("[Read Error]: %s", ShowError(err))
		return nil, err
	}
	return user, err
}

func (c *Client) UpdateItem(params *stripe.AccountParams, Email string) (*stripe.Account, error) {
	stripe.Key = c.authToken
	Id := c.GetUserId(Email)
	if len(Id) == 0 {
		return nil, fmt.Errorf("user does not exist")
	}
	user, err := account.Update(
		Id,
		params,
	)
	if err != nil {
		log.Printf("[Update Error]: %s", ShowError(err))
		return nil, err
	}
	return user, err
}

func (c *Client) DeleteItem(Email string) (*stripe.Account, error) {
	stripe.Key = c.authToken
	Id := c.GetUserId(Email)
	if len(Id) == 0 {
		return nil, fmt.Errorf("user does not exist")
	}
	user, err := account.Del(Id, nil)
	if err != nil {
		log.Printf("[Delete Error]: %s", ShowError(err))
		return nil, err
	}
	return user, err
}

func (c *Client) IsRetry(err error) bool {
	if err != nil {
		if stripeErr, ok := err.(*stripe.RateLimitError); ok {
			if strings.Contains(stripeErr.Error(), "429") {
				return true
			}
		}
	}
	return false
}

func (c *Client) GetUserId(Email string) string {
	stripe.Key = c.authToken
	params := &stripe.AccountListParams{}
	i := account.List(params)
	for i.Next() {
		a := i.Account()
		if a.Email == Email {
			return a.ID
		}
	}
	return ""
}
