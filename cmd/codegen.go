package main

import (
	"log"
	"slices"
	"sort"
	"strings"

	"github.com/ebitengine/purego"
	"github.com/kitech/gopp"
	"github.com/kitech/gopp/cgopp"
	"github.com/qtui/qtsyms"
)

func Genqtclzszcpp() {

	classes := gopp.MapKeys(qtsyms.QtSymbols)
	sort.Strings(classes)

	cp := gopp.NewCodePager()
	cp.AP("", "#include <QtCore>")
	cp.AP("", "#include <QtGui>")
	cp.AP("", "#include <QtWidgets>")
	cp.AP("", "#include <QtQml>")
	cp.AP("", "#include <QtNetwork>")
	cp.AP("", "#include <QtOpenGL>")
	cp.AP("", "#include <QtQuick>")
	cp.AP("", "#include <QtQuickControls2>")
	cp.AP("", "#include <QtQuickTemplates2>")
	cp.AP("", "#include <QtQuickWidgets>")

	cp.AP("", "int genqtclzsz(quint64 crc){")
	cp.AP("", "switch (crc){")
	gopp.Mapdo(classes, func(idx int, v string) {
		// log.Println(idx, v)
		// if strings.HasPrefix(v, "QAbstract") {
		// 	return
		// }
		// if strings.HasPrefix(v, "QAccessible") {
		// 	return
		// }
		// if strings.HasPrefix(v, "QApple") {
		// 	return
		// }
		// if strings.HasPrefix(v, "QBacking") {
		// 	return
		// }
		// if strings.HasPrefix(v, "QBenchmark") {
		// 	return
		// }

		// if slices.Contains(gopp.Sliceof("QArgumentType", "QActionAnimation", "QAdoptedThread", "QAlphaPaintEngine", "QAlphaWidget", "QAnimationGroupJob", "QAnimationTimer", "QBasicDrag"), v) {
		// 	return
		// }
		if slices.Contains(gopp.Sliceof("QNativeInterface", "QPasswordDigestor", "QQuickOpenGLUtils", "QSsl"), v) {
			return
		}
		// if slices.Contains(gopp.Sliceof("QPain", "QRa", "QLockFile", "QLoggingC", "QP"), v) {
		// 	return
		// }
		crc := gopp.Crc64Str(v)
		// cp.APf("", "#include <%s> // %d", v, idx)
		// todo todo replace #ifdef
		cp.APf("", "#ifdef %s_H", strings.ToUpper(v))
		cp.APf("", "case %vUL: return int(sizeof(%s));  // %d", crc, v, idx)
		cp.APf("", "#endif")
	})
	cp.AP("", "}")
	cp.AP("", "return -1;")
	cp.AP("", "}")
	codestr := cp.ExportAll()
	// log.Println(cp.ExportAll())

	savefile := "../qtallcc/genqtclzsz.cpp"
	gopp.SafeWriteFile(savefile, []byte(codestr), 0644)
}

func Rungenclzsz() {

	dlh, err := purego.Dlopen("../qtallcc/libQtAllInline.dylib", purego.RTLD_LAZY)
	gopp.ErrPrint(err)
	name := "_Z10genqtclzszy"
	sym, err := purego.Dlsym(dlh, name)
	gopp.ErrPrint(err)

	classes := gopp.MapKeys(qtsyms.QtSymbols)
	sort.Strings(classes)

	validcnt := 0
	cp := gopp.NewCodePager()
	cp.AP("", "package qtclzsz\n")
	cp.AP("", "// for amd64")
	cp.AP("", "func getfromgened(class string) int {")
	cp.AP("", "switch class {")
	gopp.Mapdo(classes, func(idx int, v string) {
		crc := gopp.Crc64Str(v)
		sz := cgopp.FfiCall[int](sym, crc)
		// log.Println(idx, v, sz)
		comment := ""
		if sz == -1 || sz == 4294967295 {
			comment = "// "
		} else {
			validcnt++
		}
		cp.APf("", "%scase \"%s\": return %v", comment, v, sz)
	})
	cp.AP("", "}")
	cp.AP("", "return -1")
	cp.AP("", "}")
	codestr := cp.ExportAll()
	// log.Println(codestr)

	savefile := "../../qtui/qtclzsz/sizes_gened.go"
	gopp.SafeWriteFile(savefile, []byte(codestr), 0644)
	log.Println("gen valid", validcnt, "of", len(classes))
}
