package transaction

func (h *Handler) insert(dto InsertDto, userId int) (*TransactionModel, error) {
	stmt := `INSERT INTO transaction (amount, description, category, type, userId) VALUE(?, ?, ?, ?, ?)`
	result, err := h.Db.Exec(stmt, dto.Amount, dto.Description, dto.Category, dto.Type, userId)

	if err != nil {
		return nil, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	transaction, err := h.get(lastId, userId)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (h *Handler) get(transactionId int64, userId int) (*TransactionModel, error) {
	stmt := `SELECT * FROM transaction WHERE id = ? AND userId = ?`
	row := h.Db.QueryRow(stmt, transactionId, userId)
	t := &TransactionModel{}

	err := row.Scan(
		&t.Id,
		&t.Amount,
		&t.Description,
		&t.Category,
		&t.Type,
		&t.UserId,
		&t.CreatedAt,
		&t.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (h *Handler) getTransaction(userId *int) ([]*TransactionModel, error) {
	stmt := `SELECT * FROM transaction WHERE userId = ? AND createdAt BETWEEN date_sub(now(), INTERVAL 1 WEEK) and now() ORDER BY createdAt DESC LIMIT 10`
	rows, err := h.Db.Query(stmt, userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	transactions := []*TransactionModel{}

	for rows.Next() {
		t := &TransactionModel{}
		err = rows.Scan(&t.Id, &t.Amount, &t.Description, &t.Category, &t.Type, &t.UserId, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
