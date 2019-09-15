API=$HOME/Developments/vuetify/packages/api-generator/dist/api.js

LIST=$(cat $API | vuetifyapi2go -list=$1)
IFS=$'\n'
for item in $LIST
do
	FN=$(echo $item|cut -c 3-)
	echo "$item => $FN"
	cat $API | vuetifyapi2go -comp=$item > ./$FN.go
done
