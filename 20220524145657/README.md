# Strip image metadata

Included in the `imagemagick` package is a tool for stripping metadata!
```
mogrify -strip /path/to/image.jpg
```

It works for other image formats too!

From the docs:
>Strip the image of any profiles, comments or these PNG chunks: bKGD,cHRM,EXIF,gAMA,iCCP,iTXt,sRGB,tEXt,zCCP,zTXt,date.
>To remove the orientation chunk, orNT, set the orientation to undefined, e.g., -orient Undefined.

Related:
* https://www.imagemagick.org/script/command-line-options.php#strip

    #metadata #tips #images
