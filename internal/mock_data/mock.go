package mockdata

import "github.com/sondrefjellving/book-logger/internal/data_types"


func GetBooksMock() []data_types.Book {
	return []data_types.Book{
		{
			Title: "Lord of The Rings",
			Author: "JRR. Tolkien",
			NumPages: 550,
			Entries: []data_types.Entry{
				{
					Date: "2024-06-23",
					CurrentPage: 40,
					Summary: "Frodo got the ring from Bilbo",
				},
			},
		},
	}
}