cd reverse/

echo "building reverse module..."
go build
echo "ok!"

cp reverse ../server/

cd ../server

echo "building server module..."
go build
echo "ok"

./server
