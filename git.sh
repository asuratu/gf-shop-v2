
CONTENT=$1

# CONTENT 默认值为 "update"
if [ -z "$CONTENT" ]; then
    CONTENT="update"
fi

git add .
git commit -m "$CONTENT"
git push
