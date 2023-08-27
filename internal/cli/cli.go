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

func HandleCommand() {
	if len(os.Args) < 2 {
		println("No command provided")
		os.Exit(1)
	}
	command := os.Args[1]
	switch command {
	case "set-role":
		setRole()
	case "user-details":
		printUserDetails()
	default:
		println("Unknown command")
		os.Exit(1)
	}
}

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
		println("Error marshalling user:", err)
		os.Exit(1)
	}
	fmt.Println(string(json))
}

func getUser(client *auth.Client, uid string, email string) *auth.UserRecord {
	if uid == "" && email == "" {
		println("You must provide either uid or email")
		os.Exit(1)
	}

	var user *auth.UserRecord
	var err error
	if uid != "" {
		user, err = client.GetUser(context.Background(), uid)
	} else {
		user, err = client.GetUserByEmail(context.Background(), email)
	}
	if err != nil {
		println("Error fetching user:", err)
		os.Exit(1)
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
		println("Error setting custom claims:", err)
		os.Exit(1)
	}
}

func newFirebaseAuthClient(projectId string) *auth.Client {
	config := &firebase.Config{
		ProjectID: projectId,
	}
	app, err := firebase.NewApp(context.Background(), config)
	if err != nil {
		println("Error initializing firebase app:", err)

		os.Exit(1)
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		println("Error initializing firebase auth client:", err)
		os.Exit(1)
	}

	return client
}
