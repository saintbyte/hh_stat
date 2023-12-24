package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	structures "github.com/saintbyte/hh_stat/internal"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var NegotiationStatus = [...]string{"all", "active", "invitations", "response", "discard", "archived", "non_archived", "deleted"}

func input(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}
func GetHHAuthURL() string {
	res := fmt.Sprintf(
		"https://hh.ru/oauth/authorize?response_type=code&client_id=%s&state=%s&redirect_uri=%s",
		os.Getenv("CLIENT_ID"),
		"hh_auth",
		os.Getenv("REDIRECT_URI"),
	)
	return res
}

func getCodeFromUrl(urlStr string) string {
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	m, _ := url.ParseQuery(parsedUrl.RawQuery)
	return m["code"][0]
}
func getAccessToken(code string) string {
	log.Println(code)
	urlStr := "https://hh.ru/oauth/token"
	data := url.Values{}
	data.Set("client_id", os.Getenv("CLIENT_ID"))
	data.Add("client_secret", os.Getenv("CLIENT_SECRET"))
	data.Add("code", code)
	data.Add("grant_type", "authorization_code")
	data.Add("redirect_uri", os.Getenv("REDIRECT_URI"))
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	resp, _ := client.Do(r)
	var jsonData map[string]string
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	json.Unmarshal(body, &jsonData)
	return jsonData["access_token"]
}
func getChats(type_of_chat string, accessToken string) int {
	urlStr := fmt.Sprintf("https://api.hh.ru/negotiations?status=%s", type_of_chat)
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, urlStr, strings.NewReader(""))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf(" Bearer %s", accessToken))
	resp, _ := client.Do(r)
	defer resp.Body.Close()
	var jsonData structures.Negotiations
	body, _ := io.ReadAll(resp.Body)
	//log.Println(string(body))
	json.Unmarshal(body, &jsonData)
	//log.Println("Found:", strconv.Itoa(jsonData.Found))
	//log.Println("Per_page:", strconv.Itoa(jsonData.PerPage))
	return jsonData.Found
}
func main() {
	log.Println("start")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(GetHHAuthURL())
	codeUrl := input("Input redirected url:")
	accessToken := getAccessToken(getCodeFromUrl(codeUrl))
	var res int = 0
	for index, typeOfChat := range NegotiationStatus {
		res = getChats(typeOfChat, accessToken)
		log.Println(fmt.Sprintf("%d) %s = %d", index, typeOfChat, res))
	}
	log.Println(accessToken)
	log.Println("finish")
}
