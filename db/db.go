package db

import "github.com/jackc/pgx/v4/pgxpool"

// Pool is global for now
var Pool *pgxpool.Pool
