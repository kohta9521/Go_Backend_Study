chapters=("1.os" "2.time" "3.math" "4.rand" "5.fmt" "6.fmt2" "7.log" "8.trconv" "9.strings" "10.bufio" "11.ioutil" "12.regexp" "13.regexp2" "14.sync" "15.crypto" "16.crypt2" "17.json" "18.json2" "19.sort" "20.context" "21.neturl" "22.httpsclient" "23.httpclient2" ""24.httpservere)

for chapters in "${chapters[@]}"; do
    mkdir "/Users/kohtakochi/go/src/src/study/Go_Backend_Study/udemy/10.標準ライブラリ/$chapters" && touch "/Users/kohtakochi/go/src/src/study/Go_Backend_Study/udemy/10.標準ライブラリ/$chapters/main.go"
done