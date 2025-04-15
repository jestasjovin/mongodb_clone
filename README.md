# MongoDB Clone — Golang

A document-based database clone inspired by MongoDB, built from scratch in Go.

## 💡 Goals

- Document-based storage .
- Page-oriented I/O like real database engines.
- Build modules progressively: Pager → Collection → Query → Engine.

---

## ✅ Phase 1: Page & Pager

- `Page`: Fixed-size data block (default: 4KB).
- `Pager`: Reads/Writes `Page`s to disk and caches them.

---

## 💡 Usage

```bash

# testing storage
make  test_storage

# # testing storage
# make test_cli

```



# MongoDB Clone — Modules

| Phase | Module              | Purpose                                       |
| ----- | ------------------- | --------------------------------------------- |
| 1     | `Page` & `Pager`    | Low-level file storage, fixed-size pages.     |
| 2     | `Document`          | JSON-like `map[string]interface{}`.           |
| 3     | `Collection`        | Store and manage documents in memory + pages. |
| 4     | `Index`             | Hash-based or B-Tree indexing by field.       |
| 5     | `QueryEngine`       | Support for `$eq`, `$gt`, `$lt`, `$in` etc.   |
| 6     | `Database`          | Multi-collection manager (Stick to low).      |
| 7     | `Transaction Layer` | Commit, Rollback (If i get time).             |
| 8     | `Wire Protocol`     | CRUD TCP Protocol. (If i get time)            |
| 9     | `Shell Parser`      | Parse queries like: `db.users.find({...})`.   |

---


mongodb-clone/
├── internal/
│   └── storage/
│       ├── page.go
│       └── pager.go
│        └── tests/
│           ├── page_test.go
│           └── pager_test.go
.........
.........
├── main.go
├── go.mod
└── README.md
