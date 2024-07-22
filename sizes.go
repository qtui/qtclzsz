package qtclzsz

// -1 for 404
func Get(class string) int {
	sz := getfromhandy(class)
	if sz == -1 {
		sz = getfromgened(class)
	}
	return sz
}

// hotfix here
func getfromhandy(class string) int {
	switch class {
	// case "QString":
	// 	return 123
	// case "QVariant":
	// 	return 123
	// case "QObject":
	// 	return 123
	// case "QCoreApplication":
	// 	return 123
	// case "QGuiApplication":
	// 	return 123
	// case "QApplication":
	// 	return 123
	// case "QQmlApplicationEngine":
	// 	return 123
	}

	return -1
}
