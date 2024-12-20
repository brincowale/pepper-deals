package pepper

import (
	"encoding/json"
	"log"
	"net/http"
	"pepper-deals/config"
	"regexp"
	"strings"

	"github.com/dghubble/oauth1"
)

type Deal struct {
	Description string  `json:"description"`
	Title       string  `json:"title"`
	Price       float64 `json:"price,omitempty"`
	DealURI     string  `json:"deal_uri"`
	Groups      []struct {
		Name string `json:"name"`
	} `json:"groups"`
	Merchant struct {
		URLName string `json:"url_name"`
	} `json:"merchant"`
}

type Deals struct {
	Data []Deal `json:"data"`
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
	req.Header.Set("Pepper-JSON-Format", "thread=full,group=full,formatted_text=html")
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

func Matches(deal Deal, filters []config.Filter) bool {
	deal.Merchant.URLName = strings.ReplaceAll(deal.Merchant.URLName, "-", ".")
	var categories string
	for _, category := range deal.Groups {
		categories += category.Name + " | "
	}
	for _, filter := range filters {
		dealText := strings.ToLower(deal.Title + " " + deal.Description + " " + categories)
		matchedIncludeText, _ := regexp.MatchString(strings.ToLower(filter.Include), dealText)
		var matchedExcludeText bool
		if filter.Exclude != "" {
			matchedExcludeText, _ = regexp.MatchString(strings.ToLower(filter.Exclude), dealText)
		}
		matchedIncludeWebsite := true
		if filter.IncludeWebsite != "" {
			matchedIncludeWebsite, _ = regexp.MatchString(strings.ToLower(filter.IncludeWebsite), deal.Merchant.URLName)
		}
		matchedExcludeWebsite := false
		if filter.ExcludeWebsite != "" {
			matchedExcludeWebsite, _ = regexp.MatchString(strings.ToLower(filter.ExcludeWebsite), deal.Merchant.URLName)
		}
		priceInRange := deal.Price >= filter.LowestPrice && deal.Price <= filter.MaximumPrice
		if priceInRange && matchedIncludeText && !matchedExcludeText && matchedIncludeWebsite && !matchedExcludeWebsite {
			return true
		}
	}
	return false
}
