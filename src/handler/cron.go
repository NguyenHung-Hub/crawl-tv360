package handler

import "log"

func (s *Server) registerJob() error {

	_, err := s.cronjob.AddFunc("10 1 * * *", func() {
		MainJob(s.store)
		log.Println("abc")
	})
	if err != nil {
		log.Println(err)
	}

	_, err = s.cronjob.AddFunc("10 2 * * *", func() {
		JobTV(s.store)
	})

	return err

}
