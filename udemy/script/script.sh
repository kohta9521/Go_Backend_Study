samples=("スライス" "append" "cap" "copy" "for" "可変長引数" "map" "mapfor" "channel" "channelImage" "channelGorutine" "channelClose" "channelFor" "channelSelect")

# 各サブディレクトリを作成し、その中にmain.goを作成する
for sample in "${samples[@]}"; do
    mkdir "/Users/kohtakochi/go/src/src/study/Go_Backend_Study/udemy/9.参照型/$sample" && touch "/Users/kohtakochi/go/src/src/study/Go_Backend_Study/udemy/9.参照型/$sample/main.go"
done