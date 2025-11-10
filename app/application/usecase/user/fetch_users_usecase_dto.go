package user

type FetchUsersUsecaseOutputDTO struct {
	ID string
	Name string
}

// NOTE:
// ここは配列も含めてDTOを定義した方が良いかと思った。
// 配列全体に関する情報を持たせたいニーズもアリそう。
// 例：ページネーション
