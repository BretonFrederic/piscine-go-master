rep=$(ls -R | sed '/^$/d' | sed '/^[.]/d' | wc -l)
let "rep++"
echo $rep