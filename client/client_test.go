package client

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stripe/stripe-go/v72"
)

func TestClient_NewItem(t *testing.T) {
	testCases := []struct {
		testName     string
		newItem      stripe.AccountParams
		id           string
		expectedResp *UserInfo
		expectErr    bool
	}{
		{
			testName: "user created successfully",
			newItem: stripe.AccountParams{
				Email:        stripe.String("madhurdhodria@gmail.com"),
				Country:      stripe.String("IN"),
				Type:         stripe.String("custom"),
				BusinessType: stripe.String("individual"),
				Capabilities: &stripe.AccountCapabilitiesParams{
					CardPayments: &stripe.AccountCapabilitiesCardPaymentsParams{
						Requested: stripe.Bool(true),
					},
					Transfers: &stripe.AccountCapabilitiesTransfersParams{
						Requested: stripe.Bool(true),
					},
				},
				Individual: &stripe.PersonParams{
					Email:     stripe.String("madhurdhodria@gmail.com"),
					FirstName: stripe.String("madhur"),
					LastName:  stripe.String("dhodria"),
				},
			},
			expectedResp: &UserInfo{
				Email:     "madhurdhodria@gmail.com",
				FirstName: "madhur",
				LastName:  "dhodria",
			},
			expectErr: false,
		},
		{
			testName: "user already exists",
			newItem: stripe.AccountParams{
				Email:        stripe.String("ashishdhodria1999@gmail.com"),
				Country:      stripe.String("IN"),
				Type:         stripe.String("custom"),
				BusinessType: stripe.String("individual"),
				Capabilities: &stripe.AccountCapabilitiesParams{
					CardPayments: &stripe.AccountCapabilitiesCardPaymentsParams{
						Requested: stripe.Bool(true),
					},
					Transfers: &stripe.AccountCapabilitiesTransfersParams{
						Requested: stripe.Bool(true),
					},
				},
				Individual: &stripe.PersonParams{
					Email:     stripe.String("ashishdhodria1999@gmail.com"),
					FirstName: stripe.String("ashish"),
					LastName:  stripe.String("dhodria"),
				},
			},
			expectErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("STRIPE_SECRETKEY"))
			user, err := client.NewItem(&tc.newItem)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			log.Println("ID: ", user.ID)
			userInfo := &UserInfo{
				Email:     user.Email,
				FirstName: user.Individual.FirstName,
				LastName:  user.Individual.LastName,
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, userInfo)
		})
	}
}

func TestClient_GetItem(t *testing.T) {
	testCases := []struct {
		testName     string
		id           string
		expectErr    bool
		expectedResp *UserInfo
	}{
		{
			testName:  "user exists",
			id:        "madhurdhodria@gmail.com",
			expectErr: false,
			expectedResp: &UserInfo{
				Email:     "madhurdhodria@gmail.com",
				FirstName: "madhur",
				LastName:  "dhodria",
			},
		},
		{
			testName:     "user does not exist",
			id:           "ashishdhodria@gmail.com",
			expectErr:    true,
			expectedResp: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("STRIPE_SECRETKEY"))
			user, err := client.GetItem(tc.id)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			userInfo := &UserInfo{
				Email:     user.Email,
				FirstName: user.Individual.FirstName,
				LastName:  user.Individual.LastName,
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, userInfo)
		})
	}
}

func TestClient_UpdateItem(t *testing.T) {
	testCases := []struct {
		testName     string
		updatedUser  stripe.AccountParams
		expectedResp *UserInfo
		id           string
		expectErr    bool
	}{
		{
			testName: "user exists",
			id:       "madhurdhodria@gmail.com",
			updatedUser: stripe.AccountParams{
				Individual: &stripe.PersonParams{
					FirstName: stripe.String("Tinku"),
					LastName:  stripe.String("Malav"),
				},
			},
			expectedResp: &UserInfo{
				Email:     "madhurdhodria@gmail.com",
				FirstName: "Tinku",
				LastName:  "Malav",
			},
			expectErr: false,
		},
		{
			testName: "user does not exist",
			id:       "ashishdhodria@gmail.com",
			updatedUser: stripe.AccountParams{
				Individual: &stripe.PersonParams{
					FirstName: stripe.String("ashish"),
					LastName:  stripe.String("dhodria"),
				},
			},
			expectErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("STRIPE_SECRETKEY"))
			_, err := client.UpdateItem(&tc.updatedUser, tc.id)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			user, err := client.GetItem(tc.id)
			assert.NoError(t, err)
			userInfo := &UserInfo{
				Email:     user.Email,
				FirstName: user.Individual.FirstName,
				LastName:  user.Individual.LastName,
			}
			assert.Equal(t, tc.expectedResp, userInfo)
		})
	}
}

func TestClient_DeleteItem(t *testing.T) {
	testCases := []struct {
		testName  string
		id        string
		expectErr bool
	}{
		{
			testName:  "user exists",
			id:        "madhurdhodria@gmail.com",
			expectErr: false,
		},
		{
			testName:  "user does not exist",
			id:        "ashishdhodria@gmail.com",
			expectErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("STRIPE_SECRETKEY"))
			_, err := client.DeleteItem(tc.id)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
