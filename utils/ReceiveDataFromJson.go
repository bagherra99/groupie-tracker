package utils

// import "fmt"
import (
	"main/model"
	"main/src"
)

// type AllArtists struct {
// 	ID           int
// 	Image        string
// 	Name         string
// 	Members      []string
// 	CreationDate int
// 	FirstAlbum   string
// 	Locations    []string
// 	ConcertDates []string
// 	Relations    map[string][]string
// }

	func ReceiveDataFromJson() []model.AllArtists {
		artists := src.GetDataFromApiArtist()
		locations := src.GetDataFromApiLocations()
		dates := src.GetDataFromApiDates()
		relation := src.GetDataFromApiRelation()

		var SingleInfo model.AllArtists
		var AllInfo []model.AllArtists

		for _, artist := range artists {
			SingleInfo.ID = artist.ID
			SingleInfo.Image = artist.Image
			SingleInfo.Name = artist.Name
			SingleInfo.Members = artist.Members
			SingleInfo.CreationDate = artist.CreationDate
			SingleInfo.FirstAlbum = artist.FirstAlbum

			for _, v := range locations.Index {
				if SingleInfo.ID == v.ID {
					SingleInfo.Locations = v.Locations
				}
			}

			for _, v := range dates.Index {
				if SingleInfo.ID == v.ID {
					SingleInfo.ConcertDates = v.Dates
				}
			}
			
			for _, v := range relation.Index {
				if SingleInfo.ID == v.ID {
					SingleInfo.Relations = v.DatesandLocations
				}
			}

			AllInfo = append(AllInfo, SingleInfo)
		}

		return AllInfo
	}
