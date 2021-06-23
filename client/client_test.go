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
		newItem      stripe.CustomerParams
		id           string
		expectedResp *UserInfo
		expectErr    bool
	}{
		{
			testName: "user created successfully",
			newItem: stripe.CustomerParams{
				Email: stripe.String("madhurdhodria@gmail.com"),
				Name:  stripe.String("madhur dhodria"),
			},
			expectedResp: &UserInfo{
				Email: "madhurdhodria@gmail.com",
				Name:  "madhur dhodria",
			},
			expectErr: false,
		},
		{
			testName: "user already exists",
			newItem: stripe.CustomerParams{
				Email: stripe.String("ashishdhodria1999@gmail.com"),
				Name:  stripe.String("ashish dhodria"),
			},
			expectErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("STRIPE_TOKEN"))
			user, err := client.NewItem(&tc.newItem)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			log.Println("ID: ", user.ID)
			userInfo := &UserInfo{
				Email: user.Email,
				Name:  user.Name,
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
				Email: "madhurdhodria@gmail.com",
				Name:  "madhur dhodria",
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
			client := NewClient(os.Getenv("STRIPE_TOKEN"))
			user, err := client.GetItem(tc.id)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			userInfo := &UserInfo{
				Email: user.Email,
				Name:  user.Name,
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, userInfo)
		})
	}
}

func TestClient_UpdateItem(t *testing.T) {
	testCases := []struct {
		testName     string
		updatedUser  stripe.CustomerParams
		expectedResp *UserInfo
		id           string
		expectErr    bool
	}{
		{
			testName: "user exists",
			id:       "madhurdhodria@gmail.com",
			updatedUser: stripe.CustomerParams{
				Name: stripe.String("Tinku Malav"),
			},
			expectedResp: &UserInfo{
				Email: "madhurdhodria@gmail.com",
				Name:  "Tinku Malav",
			},
			expectErr: false,
		},
		{
			testName: "user does not exist",
			id:       "ashishdhodria@gmail.com",
			updatedUser: stripe.CustomerParams{
				Name: stripe.String("ashish dhodria"),
			},
			expectErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("STRIPE_TOKEN"))
			_, err := client.UpdateItem(&tc.updatedUser, tc.id)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			user, err := client.GetItem(tc.id)
			assert.NoError(t, err)
			userInfo := &UserInfo{
				Email: user.Email,
				Name:  user.Name,
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
			client := NewClient(os.Getenv("STRIPE_TOKEN"))
			_, err := client.DeleteItem(tc.id)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
