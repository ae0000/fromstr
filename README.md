# fromstr
Little helpers for converting from a string to whatever


For example
-----------

```
id := fromstr.Int64("123", -1)
articleID := fromstr.Int64(chi.URLParam(r, "articleID"), 0) 
```

or 

```
price := fromstr.Float64(r.FormPost("price"), 0.0)
```

