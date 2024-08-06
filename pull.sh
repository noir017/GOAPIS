git pull origin master
ps -aux | grep goapis | awk '{print $2}' | xargs kill
chmod 777 ./goapis
./goapis