
CONTENT=$1

# CONTENT 默认值为 "update"
if [ -z "$CONTENT" ]; then
    CONTENT="update"
fi

git add .
git commit -m "$CONTENT"
git push

echo -e "\033[32m代码推送成功!\033[0m"
