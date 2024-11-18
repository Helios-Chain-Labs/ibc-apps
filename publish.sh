VERSION=${VERSION:-"v0.50.1-helios-1"}

echo "Deploy ibc-apps/modules/ibc-hooks/v8"
git add .
git commit -m "Publish $VERSION"
git push
git tag ibc-apps/modules/ibc-hooks/v8/$VERSION
git push origin ibc-apps/modules/ibc-hooks/v8/$VERSION
cd ./modules/ibc-hooks
GOPROXY=proxy.golang.org go list -m github.com/Helios-Chain-Labs/ibc-apps/modules/ibc-hooks/v8@$VERSION

echo "Publish done"