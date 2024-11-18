VERSION=${VERSION:-"v8.50.10-helios-1"}

# version must be start by v8

echo "Deploy ibc-apps/modules/ibc-hooks/v8"
git add .
git commit -m "Publish $VERSION"
git push
git tag modules/ibc-hooks/$VERSION
git push origin modules/ibc-hooks/$VERSION
cd ./modules/ibc-hooks
sleep 5
GOPROXY=proxy.golang.org go list -m github.com/Helios-Chain-Labs/ibc-apps/modules/ibc-hooks/v8@$VERSION

echo "Publish done"