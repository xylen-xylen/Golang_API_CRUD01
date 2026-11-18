package store

import (
	"database/sql"
	"time"
)

type Post struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Created_at time.Time `json:"created_at"`
}

type PostStore interface {
	CreatePost(*Post) (*Post, error)
	GetPostById(id int64) (*Post, error)
}
type PostgresPostStore struct {
	db *sql.DB
}

func NewPostgresPostStore(db *sql.DB) *PostgresPostStore {
	return &PostgresPostStore{db: db}
}

func (pg *PostgresPostStore) CreatePost(post *Post) (*Post, error) {
	tx, err := pg.db.Begin() // start database here ..
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	query := `INSERT INTO posts (title,content) VALUES ($1,$2) RETURNING id`
	//RETURNING id က အသစ်ထည့်လိုက်တဲ့ row ရဲ့ ID ကို ပြန်ပေးဖို့ ပြောတာပါ။

	err = tx.QueryRow(query, post.Title, post.Content).Scan(&post.ID)
	//ဒီ SQL က Database ကို Post အသစ်ထည့်ပြီး အသစ်ဖန်တီးလိုက်တဲ့ row ရဲ့ ID ကို ပြန်ပေးပါ လို့ ခိုင်းတာပါ။
	//Scan(&post.Id) = Database ကပြန်ပေးတဲ့ ID ကို Go variable ထဲ ထည့်တာ။ post struct

	//从handler来的id也放在post struct 但是新建的这个 RETURNING id。。
	// &post က pointer ဖြစ်လို့ Store ထဲမှာ မူလ post variable ကိုပဲ ပြင်ဆင်နိုင်တယ်။

	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return post, nil

}
func (pg *PostgresPostStore) GetPostById(id int64) (*Post, error) { //ဒီ function ရဲ့ ရည်ရွယ်ချက်က Database ထဲမှာ ရှိပြီးသား Post တစ်ခုကို ID နဲ့ ရှာဖို့ ပါ။
	post := &Post{}
	query := `SELECT id, title, content, created_at FROM posts WHERE id = $1`
	err := pg.db.QueryRow(query, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.Created_at,
	)

	return post, nil
}
