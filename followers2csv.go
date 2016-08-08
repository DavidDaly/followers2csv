package main

import "strconv"
import "github.com/ChimeraCoder/anaconda"
import "encoding/csv"
import "os"
import "net/url"

func main() {

	anaconda.SetConsumerKey("your-consumer-key")
	anaconda.SetConsumerSecret("your-consumer-secret")
	api := anaconda.NewTwitterApi("your-access-token", "your-access-token-secret")

	v := url.Values{}
	v.Set("count", "200")
	pages := api.GetFollowersListAll(v)
	
	var records = [][]string{ 
				{"Name",
				"ScreenName", 
				"Number of Followers",
				"Number of Tweets",
				"Following",
				"Location",
				"Description"} }

	for page := range pages {
    	
		for _,user := range page.Followers {
			records = append(records, []string{
				user.Name,
				user.ScreenName,
				strconv.Itoa(user.FollowersCount),
				strconv.FormatInt(user.StatusesCount, 10),
				strconv.FormatBool(user.Following),
				user.Location,
				user.Description } )
		}

	}

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(records)

}
