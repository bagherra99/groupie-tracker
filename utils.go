package main

// import "fmt"

type AllArtists struct {
	ID           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    []string
	ConcertDates []string
	Relations    map[string][]string
}

func ReceiveDataFromJson() []AllArtists {
	artists := getDataFromApiArtist()
	locations := getDataFromApiLocations()
	dates := getDataFromApiDates()
	relation:=getDataFromApiRelation()
	var SingleInfo AllArtists
	var AllInfo []AllArtists
	for _, artist := range artists {
		SingleInfo.ID = artist.ID
		SingleInfo.Image = artist.Image
		SingleInfo.Name = artist.Name
		SingleInfo.Members = artist.Members
		SingleInfo.CreationDate = artist.CreationDate
		SingleInfo.FirstAlbum = artist.FirstAlbum
		AllInfo = append(AllInfo, SingleInfo)
	}
	for _, v := range locations.Index {
		if SingleInfo.ID == v.ID {
			SingleInfo.Locations = v.Locations
		}
		AllInfo = append(AllInfo, SingleInfo)
	}
	for _, v := range dates.Index {
		if SingleInfo.ID == v.ID {
			SingleInfo.ConcertDates = v.Dates
		}
		AllInfo = append(AllInfo, SingleInfo)
	}
	for _, v:=range relation.Index{
		if SingleInfo.ID==v.ID{
			SingleInfo.Relations=v.DatesandLocations
		}
		AllInfo = append(AllInfo, SingleInfo)
	}
	// fmt.Println(AllInfo)
	return AllInfo
}
