# heroku git:remote -a "$(dirname "pwd")"
now=`date '+%Y%m%d%H%M%S'`
echo "commit $now"
git add .
git commit -am "commit $now"
echo "git push origin master "now
git push origin master
# echo "push heroku master"
# git push heroku master
echo "run"
# heroku ps:scale web=1
