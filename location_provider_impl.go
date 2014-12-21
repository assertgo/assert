package assert

func provideLocation() location {
	return location{
		Test:     "TestLocationProvider",
		FileName: "location_provider_test.go",
	}
}
