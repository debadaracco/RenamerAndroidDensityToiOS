if [ ! -d "./drawables-ios" ]; then
	mkdir "./drawables-ios"
fi

for folder_drawable in drawable-*
do
	folder=""
	density=""

	if [ "$folder_drawable" == "drawable-mdpi" ]; then
		echo "mdpi 1x"
		folder="mdpi"
	elif [ "$folder_drawable" == "drawable-xhdpi" ]; then
		echo "xhdpi 2x"
		folder="xhdpi"
		density="@2x"
	elif [ "$folder_drawable" == "drawable-xxhdpi" ]; then
		echo "xxhdpi 3x"
		folder="xxhdpi"
		density="@3x"
	fi

	if [ ! "$folder" == "" ]; then
		cd "./drawable-$folder"

		for file in *.png
		do
			cp "$file" "../drawables-ios/$file$density.png"
		done

		cd ".."
	fi
done