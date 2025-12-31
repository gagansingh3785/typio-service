package repository

const (
	GetParagraphsCountQuery = "SELECT COUNT(*) FROM paragraphs;"
	GetParagraphByIDQuery   = "SELECT id, uuid, content, created_at, updated_at FROM " +
		"(SELECT id, uuid, content, created_at, updated_at, ROW_NUMBER() OVER () AS row_num FROM paragraphs) " +
		"AS numbered WHERE row_num = :row_num;"

	InsertParagraphsQuery = "INSERT INTO paragraphs (content, created_at, updated_at) " +
		"VALUES (:content, :created_at, :updated_at) RETURNING id, uuid, content, created_at, updated_at;"
)
