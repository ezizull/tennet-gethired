package book

// Package  provides the use case for

import (
	domainBook "tennet/gethired/domain/book"
)

func (n *NewBook) toDomainMapper() *domainBook.Book {
	return &domainBook.Book{
		Title:       n.Title,
		Author:      n.Author,
		Description: n.Description,
	}
}
