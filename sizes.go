package qtclzsz

import "log"

// -1 for 404
func Get(class string) int {
	sz := getfromhandy(class)
	if sz == -1 {
		sz = getfromgened(class)
	}
	if sz == -1 {
		log.Println("WARN clzsz -1 => 123", class)
		sz = 123
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
	case "QModelIndex":
		return 24

	case "QMainWindow":
		return 40

	case "QVBoxLayout":
		return 32

	case "QHBoxLayout":
		return 32

	case "QSpacerItem":
		return 40

	case "QTableWidgetItem":
		return 64

	case "QStringList", "QObjectList", "QList":
		return 24
	}

	return -1
}
