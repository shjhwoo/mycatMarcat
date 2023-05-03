package scheduler

func GetTopFivePopularProductListOnEveryFridayAfternoon() error {
	task := getTopFivePopularProductList

	err := startSchedulerFor(task)
	if err != nil {
		return err
	}

	return nil
}

func getTopFivePopularProductList() {
	//db에 접근해서 가장 많이 팔린 상품 상위 5개를 보여줘야 한다
}
