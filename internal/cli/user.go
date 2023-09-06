package cli

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

func printUserDetails() {
	cmd := flag.NewFlagSet("user-details", flag.ExitOnError)
	uuid := cmd.String("uid", "", "UID of the user to promote")
	email := cmd.String("email", "", "Email of the user to promote")
	projectId := cmd.String("project", "just-code-dev", "Project ID to promote the user to")
	cmd.Parse(os.Args[2:])

	client := newFirebaseAuthClient(*projectId)

	user := getUser(client, *uuid, *email)

	json, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		panic("Error marshalling user: " + err.Error())
	}
	fmt.Println(string(json))
}

func getUser(client *auth.Client, uid string, email string) *auth.UserRecord {
	if uid == "" && email == "" {
		panic("You must provide either a uid or email")
	}

	var user *auth.UserRecord
	var err error
	if uid != "" {
		user, err = client.GetUser(context.Background(), uid)
	} else {
		user, err = client.GetUserByEmail(context.Background(), email)
	}
	if err != nil {
		panic("Error getting user: " + err.Error())
	}

	return user
}

func setRole() {
	cmd := flag.NewFlagSet("set-role", flag.ExitOnError)
	uid := cmd.String("uid", "", "UID of the user to promote")
	email := cmd.String("email", "", "Email of the user to promote")
	projectId := cmd.String("project", "just-code-dev", "Project ID to promote the user to")
	admin := cmd.Bool("admin", false, "Whether to set admin role")
	cmd.Parse(os.Args[2:])

	client := newFirebaseAuthClient(*projectId)
	user := getUser(client, *uid, *email)
	claims := map[string]interface{}{
		"admin": admin,
	}
	err := client.SetCustomUserClaims(context.Background(), user.UID, claims)
	if err != nil {
		panic("Error setting custom claims: " + err.Error())
	}
}

func newFirebaseAuthClient(projectId string) *auth.Client {
	config := &firebase.Config{
		ProjectID: projectId,
	}
	app, err := firebase.NewApp(context.Background(), config)
	if err != nil {
		panic("Error initializing firebase app: " + err.Error())
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		panic("Error initializing firebase auth client: " + err.Error())
	}

	return client
}
