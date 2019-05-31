# RenamerAndroidDensityToiOS

This app must be execute on your Android Project RES folder and make drawables-ios folder with all content renamed. 
For this, the app copy the files from drawable-mdpi, drawable-xhdpi & drawable-xxhdpi and use it with the iOS Density 
(1x, 2x & 3x)

To use it, you can download the drawable_converter release and put it on /usr/local/bin. Then you can run the following 
command on your Android Project RES folder: 

```bash
drawable_converter
```