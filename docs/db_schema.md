# db スキーマ

## tasks

| カラム名   | 型      | 属性 | 説明など                      |
| :--------- | :------ | :--- | :---------------------------- |
| id         | integer |      | タスクの id                   |
| name       | string  |      | タスク名                      |
| finish     | bool    |      | タスクが終了しているかどうか  |
| importance | integer |      | タスクの優先順位              |
| hoverstar  |         |      | DB 側ではいらないかもしれない |
| created_at |         |      | タスクの作成日時              |
