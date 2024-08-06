go build -o goapis main.go 
git add .
date_now=$(date +"%Y-%m-%d %H")
echo "commit:$date_now"
git commit -m "$date_now"
git push origin master