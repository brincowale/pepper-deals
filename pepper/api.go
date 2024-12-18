package pepper

import (
	"encoding/json"
	"log"
	"net/http"
	"pepper-deals/config"

	"github.com/dghubble/oauth1"
)

type Data struct {
	Description string  `json:"description"`
	Title       string  `json:"title"`
	Price       float64 `json:"price,omitempty"`
	DealURI     string  `json:"deal_uri"`
}

type Deals struct {
	Data []Data `json:"data"`
}

func GetNewDeals(config config.Config) *Deals {
	oauthConfig := oauth1.NewConfig(config.ConsumerKey, config.ConsumerSecret)
	httpClient := oauthConfig.Client(oauth1.NoContext, &oauth1.Token{})

	// Create a new request
	url := "https://" + config.Host + "/rest_api/v2/thread?criteria=%7B%22tab%22%3A%22new%22%2C%22whereabouts%22%3A%22deals%22%7D&history_item_needed=false&limit=25"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return nil
	}

	// Set additional headers
	req.Header.Set("User-Agent", config.PkgName+" ANDROID [v5.26.11] [22 | SM-G930K] [@2.0x]")
	req.Header.Set("Pepper-Include-Counters", "unread_alerts")
	req.Header.Set("Pepper-Include-Prev-And-Next-Ids", "true")
	req.Header.Set("Pepper-JSON-Format", "thread=full,group=ids,type=light,event=light,user=full,badge=user,formatted_text=html,message=with_code")
	req.Header.Set("Pepper-Hardware-Id", "a3f5c7e8d9b0a1c2e3f4b5a6c7d8e9f0")
	req.Header.Set("Host", config.Host)

	// Make the request using the OAuth1 client
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println("Error making request:", err)
		return nil
	}
	defer resp.Body.Close()

	// Check for a successful response
	if resp.StatusCode != http.StatusOK {
		log.Println("Error: received non-200 response status:", resp.Status)
		return nil
	}

	// Decode the JSON response
	var deals Deals
	if err := json.NewDecoder(resp.Body).Decode(&deals); err != nil {
		log.Println("Error decoding response:", err)
		return nil
	}

	return &deals
}
