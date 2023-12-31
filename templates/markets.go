package templates

import (
	"regexp"
	"strings"
)

type Market struct {
	Slug        string
	Name        string
	Address     string
	Description string
	Website     string
	Latitude    string
	Longitude   string
	Image       string
}

type Markets struct {
	Markets []Market
}

func (markets Markets) Find(name string) Market {
	for _, market := range markets.Markets {
		if market.Name == name {
			return market
		}
	}
	return Market{}
}

func (markets Markets) Search(searchTerm string) Markets {
	lowerSearch := strings.ToLower(searchTerm)
	matched := []Market{}
	for _, market := range markets.Markets {
		if strings.Contains(strings.ToLower(market.Name), lowerSearch) || strings.Contains(market.Address, searchTerm) {
			matched = append(matched, market)
		}
	}
	return Markets{Markets: matched}
}

func (markets Markets) FindSlug(slug string) Market {
	for _, market := range markets.Markets {
		if market.Slug == slug {
			return market
		}
	}
	return Market{}
}

var re = regexp.MustCompile("[^a-z0-9]+")

func (market Market) GetSlug() string {
	return strings.Trim(re.ReplaceAllString(strings.ToLower(market.Name), "-"), "-")
}
