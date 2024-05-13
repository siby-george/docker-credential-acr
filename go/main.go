package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

type tokenResponse struct {
	Refresh_token string `json:"refresh_token"`
}

func main() {
	registryName, err := getRegistryName()
	if err != nil {
		panic(err)
	}
	aadToken, err := getAADToken()
	if err != nil {
		panic(err)
	}
	registryToken, err := getRegistryToken(registryName, aadToken)
	if err != nil {
		panic(err)
	}
	fmt.Printf("{\"Username\": \"<token>\",\"Secret\": \"%s\"}", registryToken)
}

func getRegistryName() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		return "", err
	}
	return scanner.Text(), nil
}

func getRegistryToken(registryName string, aadToken string) (string, error) {
	formData := url.Values{
		"grant_type":   {"access_token"},
		"service":      {registryName},
		"access_token": {aadToken},
	}
	resp, err := http.Post(fmt.Sprintf("https://%s/oauth2/exchange", registryName),
		"application/x-www-form-urlencoded", strings.NewReader(formData.Encode()))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	token := new(tokenResponse)
	json.NewDecoder(resp.Body).Decode(&token)
	return token.Refresh_token, nil
}

func getAADToken() (string, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return "", err
	}
	context := context.TODO()
	aadToken, err := cred.GetToken(context, policy.TokenRequestOptions{Scopes: []string{"https://containerregistry.azure.net/.default"}})
	if err != nil {
		return "", err
	}
	return aadToken.Token, nil
}
