# マイグレーション
## ローカル

```bash
make migrate-local
```

## それ以外

```bash
make migrate os=linux user=${DB_USER} password=${DB_PASSWORD} host=${DB_HOST} database=${DB_DATABASE}
```

# マイグレーションファイル作成

```bash
make create-migration name=create_table_name
```
