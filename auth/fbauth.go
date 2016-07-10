package auth

type FbDetails struct {
	App_id      string   `json:"app_id"`
	Application string   `json:"application"`
	Expires_at  int      `json:"expires_at"`
	Is_valid    bool     `json:"is_valid"`
	Scopes      []string `json:"scopes"`
	User_id     string   `json:"user_id"`
}

type FbExtendedDetails struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

//Authentication details for the server application to connect to the Facebook API
var fbAT string = ""

const appSecret string = "ebd0d28a2930a961f8fe8e177f10a3fb"
const appID string = "1694761154116310"

//Details on the various endpoints for the Facebook API
const fbGraphApi = "https://graph.facebook.com"
const fbTokenEndpoint string = "/oauth/access_token"
const fbUTokenEndpoint string = "/debug_token"

//Retrieves an access token to allow the server to connect to the API
func getAccessToken() {
	// params := url.Values{}
	// params.Add("client_id", appID)
	// params.Add("client_secret", appSecret)
	// params.Add("grant_type", "client_credentials")

	// var reqUrl = fmt.Sprintf("%s%s?%s", fbGraphApi, fbTokenEndpoint, params.Encode())

	// resp, err := http.Get(reqUrl)
	// if err != nil {
	// 	panic("failed to get facebook access token!")
	// }

	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic("failed to parse response from facebook!")
	// }

	// respSections := strings.Split(fmt.Sprintf("%s", body), "=")
	// if respSections[0] != "access_token" {
	// 	panic("malformed response when obtaining access key from facebook")
	// } else {
	// 	fbAT = respSections[1]
	// }
}

//Takes a User ID retrieved from the Facebook Graph API previously, and
//uses this to retrieve more data on the user
func getUserAdvancedDetails(uid string) FbExtendedDetails {
	// if fbAT == "" {
	// 	getAccessToken()
	// }
	// queryEndpoint := fmt.Sprintf("/v2.6/%s", uid)
	// params := url.Values{}
	// params.Add("access_token", fbAT)
	// params.Add("fields", "id,name,email")

	// var reqUrl = fmt.Sprintf("%s%s?%s", fbGraphApi, queryEndpoint, params.Encode())

	// resp, err := http.Get(reqUrl)
	// if err != nil {
	// 	panic("failed to locate user on facebook")
	// }

	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic("failed to parse response from facebook!")
	// }
	var fbExDets FbExtendedDetails
	// err = json.Unmarshal(body, &fbExDets)
	// if err != nil {
	// 	panic("wat")
	// }
	return fbExDets
}

//Takes a user token passed to the backend by the Facebook JS library,
//and retrieves the details on the user from the Facebook API
func GetUserDetails(inputToken string) FbDetails {
	// if fbAT == "" {
	// 	getAccessToken()
	// }
	// params := url.Values{}
	// params.Add("input_token", inputToken)
	// params.Add("access_token", fbAT)

	// var reqUrl = fmt.Sprintf("%s%s?%s", fbGraphApi, fbUTokenEndpoint, params.Encode())

	// resp, err := http.Get(reqUrl)
	// if err != nil {
	// 	panic("failed to verify user token with facebook")
	// }

	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic("failed to parse response from facebook!")
	// }

	type FbResp struct {
		Data FbDetails `json:"data"`
	}
	var fbResp FbResp
	// err = json.Unmarshal(body, &fbResp)
	// if err != nil {
	// 	panic("err")
	// }

	return fbResp.Data
}
