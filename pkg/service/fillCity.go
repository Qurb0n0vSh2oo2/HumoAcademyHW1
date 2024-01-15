package service

func (s *Service) FillCities() (err error) {
	err = s.Repository.AddCity("Dushanbe", "RRP")

	if err != nil {
		return
	}

	err = s.Repository.AddCity("Khudzhand", "Sogd")

	if err != nil {
		return
	}

	err = s.Repository.AddCity("Istaravshan", "Sogd")

	if err != nil {
		return
	}

	err = s.Repository.AddCity("Kulob", "Khatlon")

	if err != nil {
		return
	}

	return
}
