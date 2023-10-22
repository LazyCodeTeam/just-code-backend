package cli

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"

	"github.com/LazyCodeTeam/just-code-backend/internal/api/dto"
)

type authRequestBody struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type authResponseBody struct {
	IdToken      string `json:"idToken"`
	Kind         string `json:"kind"`
	LocalId      string `json:"localId"`
	RefreshToken string `json:"refreshToken"`
	Email        string `json:"email"`
	Registered   bool   `json:"registered"`
	ExpiresIn    string `json:"expiresIn"`
}

func preludeContent() {
	cmd := flag.NewFlagSet("prelude-content", flag.ExitOnError)
	email := cmd.String("email", "", "Email of user that has permission to edit content")
	password := cmd.String("password", "", "Password of user that has permission to edit content")
	firebaseApiToken := cmd.String("firebase-api-token", "", "Firebase API token")
	userToken := cmd.String("token", "", "Token of user that has permission to edit content")
	baseUrl := cmd.String("base-url", "http://localhost:8080", "Base URL of the Just Code server")
	maxTasks := cmd.Int("max-tasks", 10, "Maximum number of tasks per section")
	minTasks := cmd.Int("min-tasks", 5, "Minimum number of tasks per section")
	maxSections := cmd.Int("max-sections", 10, "Maximum number of sections per technology")
	minSections := cmd.Int("min-sections", 5, "Minimum number of sections per technology")
	maxTechnologies := cmd.Int("max-technologies", 10, "Maximum number of technologies")
	minTechnologies := cmd.Int("min-technologies", 5, "Minimum number of technologies")

	_ = cmd.Parse(os.Args[2:])

	token := *userToken
	if token == "" && (*email == "" || *password == "" || *firebaseApiToken == "") {
		panic("You must provide either a user token or email, password, and firebase API token")
	}

	if token == "" {
		newToken, err := getToken(*email, *password, *firebaseApiToken)
		if err != nil {
			panic("Error getting token: " + err.Error())
		}
		token = newToken
	}

	body := getRandomSlice(*minTechnologies, *maxTechnologies, func() dto.ExpectedTechnology {
		return getRandomTechnology(*minSections, *maxSections, *minTasks, *maxTasks)
	})
	jsonBody, err := json.Marshal(body)
	if err != nil {
		panic("Error marshalling body: " + err.Error())
	}

	request, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("%s/admin/api/v1/content", *baseUrl),
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		panic("Error creating request: " + err.Error())
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	result, err := http.DefaultClient.Do(request)
	if err != nil {
		panic("Error sending request: " + err.Error())
	}
	if result.StatusCode >= 400 {
		panic("Error: " + result.Status)
	}
}

const loremIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas cursus dolor maximus diam finibus, sit amet congue lacus ultrices. Vestibulum nec risus at magna interdum luctus. Nunc blandit odio non nisi vestibulum, quis vehicula ante dictum. Phasellus nec erat sit amet odio sollicitudin mattis a ut velit. In sit amet maximus neque. Donec justo justo, congue et ipsum nec, consequat ullamcorper velit. Aliquam consequat dolor at ultrices vulputate. Donec in lorem sit amet lacus congue condimentum non non tortor. Duis vestibulum ultricies commodo. Sed sollicitudin ac erat vel facilisis. Quisque in lectus luctus, sagittis ipsum viverra, convallis metus. Suspendisse elit turpis, commodo in auctor non, convallis sit amet erat. Integer tellus orci, sollicitudin sed diam dignissim, sagittis ultricies est. Curabitur vel purus et elit faucibus volutpat. Duis ut efficitur sapien, eget luctus erat. Donec turpis elit, consequat id aliquet eget, mattis sit amet diam."

var loremIpsumWords = strings.Split(loremIpsum, " ")

func getWords(n int) string {
	fullCount := n / len(loremIpsumWords)
	remainder := n % len(loremIpsumWords)

	words := make([]string, 0, n)
	for i := 0; i < fullCount; i++ {
		words = append(words, loremIpsumWords...)
	}
	words = append(words, loremIpsumWords[:remainder]...)

	return strings.Join(words, " ")
}

func getRandomNumber(min int, max int) int {
	return rand.Intn(max-min) + min
}

func getRandomString(min int, max int) string {
	return getWords(getRandomNumber(min, max))
}

func getRandomTechnology(
	minSections int,
	maxSections int,
	minTasks int,
	maxTasks int,
) dto.ExpectedTechnology {
	description := getRandomString(10, 50)
	return dto.ExpectedTechnology{
		Id:          getUuid(),
		Name:        getRandomString(2, 10),
		Description: &description,
		ExpectedSections: getRandomSlice(minSections, maxSections, func() dto.ExpectedSection {
			return getRandomSection(minTasks, maxTasks)
		}),
	}
}

func getRandomSection(minTasks int, maxTasks int) dto.ExpectedSection {
	description := getRandomString(10, 50)
	return dto.ExpectedSection{
		Id:          getUuid(),
		Name:        getRandomString(2, 10),
		Description: &description,
		Tasks:       getRandomSlice(minTasks, maxTasks, getRandomTask),
	}
}

func getRandomTask() dto.ExpectedTask {
	description := getRandomString(10, 50)
	name := getRandomString(2, 10)

	return dto.ExpectedTask{
		Id:          getUuid(),
		Name:        &name,
		Description: &description,
		Difficulty:  getRandomNumber(1, 10),
		IsPublic:    getRandomBoolean(),
		IsDynamic:   getRandomBoolean(),
		Content:     getRandomTaskContent(),
	}
}

func getRandomBoolean() bool {
	return getRandomNumber(0, 2) == 1
}

func getUuid() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic("Error generating uuid: " + err.Error())
	}
	return uuid.String()
}

func getRandomTaskContent() dto.ExpectedTaskContent {
	possibleTypes := []dto.TaskContentType{
		dto.TaskContentTypeLesson,
		dto.TaskContentTypeSingleSelection,
		dto.TaskContentTypeMultiSelection,
		dto.TaskContentTypeLinesArrangement,
	}

	randomType := getRandomItem(possibleTypes)

	switch randomType {
	case dto.TaskContentTypeLesson:
		return dto.ExpectedTaskContent{
			Kind:    randomType,
			Content: getRandomString(10, 50),
		}
	case dto.TaskContentTypeSingleSelection:
		options := getRandomSlice(2, 5, func() dto.ExpectedTaskOption {
			return dto.ExpectedTaskOption{
				Content: getRandomString(10, 50),
			}
		})
		validOption := getRandomNumber(0, len(options))
		return dto.ExpectedTaskContent{
			Kind:          randomType,
			Content:       getRandomString(10, 50),
			Options:       options,
			CorrectOption: &validOption,
		}
	case dto.TaskContentTypeMultiSelection:
		options := getRandomSlice(2, 5, func() dto.ExpectedTaskOption {
			return dto.ExpectedTaskOption{
				Content: getRandomString(10, 50),
			}
		})
		validOptions := getRandomSlice(1, len(options), func() int {
			return getRandomNumber(0, len(options))
		})

		return dto.ExpectedTaskContent{
			Kind:           randomType,
			Content:        getRandomString(10, 50),
			Options:        options,
			CorrectOptions: validOptions,
		}
	case dto.TaskContentTypeLinesArrangement:
		lines := getRandomSlice(2, 5, func() dto.ExpectedTaskOption {
			return dto.ExpectedTaskOption{
				Content: getRandomString(10, 50),
			}
		},
		)

		correctOrder := getRandomSlice(1, len(lines), func() int {
			return getRandomNumber(0, len(lines))
		})

		return dto.ExpectedTaskContent{
			Kind:       randomType,
			Content:    getRandomString(10, 50),
			Lines:      lines,
			LinesOrder: correctOrder,
		}

	default:
		panic("Unknown task content type")
	}
}

func getRandomSlice[T any](min int, max int, generator func() T) []T {
	length := getRandomNumber(min, max)
	items := make([]T, 0, length)
	for i := 0; i < length; i++ {
		items = append(items, generator())
	}
	return items
}

func getRandomItem[T any](items []T) T {
	return items[getRandomNumber(0, len(items))]
}

func getToken(email string, password string, firebaseApiToken string) (string, error) {
	body := authRequestBody{
		Email:             email,
		Password:          password,
		ReturnSecureToken: true,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf(
			"https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s",
			firebaseApiToken,
		),
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return "", err
	}
	request.Header.Add("Content-Type", "application/json")

	result, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer result.Body.Close()

	var responseBody authResponseBody
	err = json.NewDecoder(result.Body).Decode(&responseBody)
	if err != nil {
		return "", err
	}

	return responseBody.IdToken, nil
}
