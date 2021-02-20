In the generated code i had to change:

```
if o.File != nil {

		// form param file
		var frFile io.ReadCloser
		if o.File != nil {
			frFile = *o.File
		}
		fFile := frFile.String()
		if fFile != "" {
			if err := r.SetFormParam("file", fFile); err != nil {
				return err
			}
		}
	}
```

To:

```
if o.File != nil {

		// form param file
		var frFile io.ReadCloser
		buf := new(bytes.Buffer)
		if o.File != nil {
			buf.ReadFrom(frFile)
		}
		newStr := buf.String()
		if newStr != "" {
			if err := r.SetFormParam("file", newStr); err != nil {
				return err
			}
		}
	}
```