package handler

import (
	"context"
	"fmt"
	"log"
	"time"
	"tv360/src/db"
	"tv360/src/models"
	"tv360/src/util"
)

// func MainJob(db db.Store) {
// 	step1(db)
// }

func MainJob(db db.Store) {
	urls := [][]string{
		{
			"https://tv360.vn/public/v1/composite/film/get-home?",
			"https://tv360.vn/public/v1/vod/get-more-content?key=film_category&",
		},
		{
			"https://tv360.vn/public/v1/composite/get-home-hbo?id=0&",
			"https://tv360.vn/public/v1/vod/get-list-item-collection?",
		},
		{
			"https://tv360.vn/public/v1/vod/get-list-category?id=0&",
			"https://tv360.vn/public/v1/vod/get-list-video-category?",
		},
	}

	for _, item := range urls {
		ids, err := getCategoryId(item[0])
		if err != nil {
			log.Printf("error: %s ", err)
		}

		for _, v := range ids {
			err = fetchContent(item[1], v, db)
			if err != nil {
				log.Printf("error: %s ", err)
			}
		}

		time.Sleep(time.Second * 20)
	}

	log.Println("--- Finish Job ---")

}

func getCategoryId(baseUrl string) ([]int64, error) {
	offset := 0
	pageSize := 30

	categoryIds := []int64{}
	for {
		var res models.Response
		url := fmt.Sprintf("%soffset=%d&limit=%d", baseUrl, offset, pageSize)
		log.Println(url)
		err := util.FetchData(url, &res)
		if err != nil {
			return nil, err
		}

		if len(res.Data) > 0 {
			for _, v := range res.Data {
				switch t := v.ID.(type) {
				case float64:
					categoryIds = append(categoryIds, int64(t))
				default:
					fmt.Printf("unexpected type %T %v \n", t, t)
				}

			}
		} else {
			break
		}
		offset += pageSize

	}

	return categoryIds, nil
}

func fetchContent(baseUrl string, idCategory int64, db db.Store) error {
	offset := 0
	pageSize := 30

	for {
		var res models.Response2
		url := fmt.Sprintf("%soffset=%d&limit=%d&id=%d", baseUrl, offset, pageSize, idCategory)
		log.Println(url)
		err := util.FetchData(url, &res)
		if err != nil {
			return err
		}

		if len(res.Data.Content) > 0 {
			for _, v := range res.Data.Content {

				if !db.IsExistsContent360(context.TODO(), v.ContentID, models.Content2{}.TableName()) {
					err := db.CreateContent(context.TODO(), v)
					if err != nil {
						log.Println(err)
					}
				}
			}

		} else {
			break
		}
		offset += pageSize
		time.Sleep(time.Millisecond * 100)
	}
	return nil
}
