package handler

import (
	"context"
	"log"
	"tv360/src/db"
	"tv360/src/models"
	"tv360/src/util"
)

func JobTV(db db.Store) {
	url := "https://tv360.vn/public/v1/live/get-detail-category?id=0&offset=0&limit=1000"
	var res models.ResponseTV
	err := util.FetchData(url, &res)
	if err != nil {
		log.Printf("error: %s ", err)
	}

	for _, v1 := range res.Data {
		for _, v2 := range v1.Content {
			if !db.IsExistsContent360(context.TODO(), string(v2.ContentID), models.ContentTV{}.TableName()) {
				err := db.CreateContentTV(context.TODO(), v2)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}

}
