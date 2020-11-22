# go-gir

The following packages do not build:

- gst-1.0

	```
	# github.com/electricface/go-gir/gst-1.0
	gst-1.0/gst_auto.go:32492:6: StaticCapsGetType redeclared in this block
		previous declaration at gst-1.0/gst_auto.go:24368:26
	gst-1.0/gst_auto.go:32508:6: StaticPadTemplateGetType redeclared in this block
		previous declaration at gst-1.0/gst_auto.go:24411:33
	gst-1.0/gst_auto.go:32875:6: TypeFindGetType redeclared in this block
		previous declaration at gst-1.0/gst_auto.go:28840:24
	```

