echo "📦 Running database migration..."
migrate -database "${DATABASE_URL}" -path /app/common/database/migrations up

echo "🚀 Starting server..."
/todo-app-server