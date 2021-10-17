## api仕様書
```aidl
例
エンドポイント説明
エンドポイント
```

1. ユーザー登録
PUT: api/register
   
2. JWTトークンを取得
GET: api/auth/jwt/create
   
3. 投稿の一覧
GET: api/todos
   
4. 特定の投稿取得
GET: api/todos/{todo_id}
   
5. タスク登録
POST: api/todos/{todo_id}
   
5. タスクアップデート
PUT: api/todos/{todo_id}
   
6. タスク削除
DELETE: api/todos/{todo_id}
   

# todo-go-api
